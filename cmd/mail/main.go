package main

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"strconv"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
)

//go:generate go-bindata config/config.toml

type Config struct {
	DB struct {
		Host     string
		Database string
		User     string
		Password string
		Port     int
	}
}

func diffRate(a float64, b float64) float64 {
	if a < b {
		return (b / a) - 1
	} else {
		return (a / b) - 1
	}
}

func encodeRFC2047(String string) string {
	// use mail's rfc2047 to encode any string
	addr := mail.Address{String, ""}
	return strings.Trim(addr.String(), " <>")
}

func main() {

	var config Config

	data, _ := Asset("config/config.toml")
	_, err := toml.Decode(string(data), &config)
	if err != nil {
		panic(err)
	}

	from := mail.Address{"アービトラージ", "admin@tk2-249-34013.vs.sakura.ne.jp"}

	db, err := sql.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp("+config.DB.Host+":"+strconv.Itoa(config.DB.Port)+")/"+config.DB.Database+"?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	dboption := redis.DialDatabase(0)
	con, err := redis.Dial("tcp", "127.0.0.1:6379", dboption)
	if err != nil {
		// handle error
	}
	defer con.Close()

	for {
		time.Sleep(5 * time.Second) // 2秒待つ

		rows, err := db.Query("SELECT exchange1, exchange2, diff FROM alert")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		ex := map[string]float64{}
		for rows.Next() {
			var exchange1 string
			var exchange2 string
			var diff float64
			if err := rows.Scan(&exchange1, &exchange2, &diff); err != nil {
				log.Fatal(err)
			}
			ex[exchange1+"_"+exchange2] = diff
		}

		rows, err = db.Query("SELECT id,border1,email FROM users")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		users := make(map[int][]interface{})
		for rows.Next() {
			var id int
			var border1 float64
			var email string
			if err := rows.Scan(&id, &border1, &email); err != nil {
				log.Fatal(err)
			}
			users[id] = []interface{}{border1, email}
		}

		// each user
		for user_id, val := range users {
			notices := map[string]float64{}
			var border1 = val[0].(float64)

			// var email = val[1].(string)

			// each exchange
			for pair, diff := range ex {

				// Hit border1 !!
				if border1 < diff {
					// fmt.Println(pair, diff, border1, email)

					// Check Redis
					key := strconv.Itoa(user_id) + ":" + strconv.FormatFloat(border1, 'f', 6, 64) + ":" + pair
					exists, _ := redis.Bool(con.Do("EXISTS", key))
					if exists {
						continue
					}

					// insert Redis
					con.Do("SET", key, 1)
					con.Do("EXPIRE", key, 60)

					// set notices(map)
					notices[pair] = diff
				}
			}
			// Send Mail using notices(map)
			var body string
			for pair, diff := range notices {
				body = body + pair + "で" + fmt.Sprint(diff*100) + "%の差があります！\n"
			}
			fmt.Println(body)

			to := mail.Address{"あなた", val[1].(string)}
			title := "差が発生しました！"

			// body := "报表内容一切正常"

			header := make(map[string]string)
			header["From"] = from.String()
			header["To"] = to.String()
			header["Subject"] = encodeRFC2047(title)
			header["MIME-Version"] = "1.0"
			header["Content-Type"] = "text/plain; charset=\"utf-8\""
			header["Content-Transfer-Encoding"] = "base64"

			message := ""
			for k, v := range header {
				message += fmt.Sprintf("%s: %s\r\n", k, v)
			}
			message += "\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

			// Connect to the server, authenticate, set the sender and recipient,
			// and send the email all in one step.
			err := smtp.SendMail(
				"127.0.0.1:25",
				nil,
				from.Address,
				[]string{to.Address},
				[]byte(message),
			)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(val[1].(string) + "にメール送りました")
		}
	}
}
