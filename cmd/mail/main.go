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
	Gmail struct {
		Username string
		Password string
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

	auth := smtp.PlainAuth(
		"",
		config.Gmail.Username,
		config.Gmail.Password,
		"smtp.gmail.com",
	)
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

		time.Sleep(60 * time.Second)

		rows, err := db.Query("SELECT id,border1,email FROM users")
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

		ex, _ := redis.StringMap(con.Do("hGetAll", "alert"))

		// each user
		for user_id, val := range users {

			notices := map[string]float64{}
			border1 := val[0].(float64)

			for pair, rate := range ex {
				diff, _ := strconv.ParseFloat(rate, 5)

				if border1 < diff {

					// Check Redis
					key := strconv.Itoa(user_id) + ":" + strconv.FormatFloat(border1, 'f', 6, 64) + ":" + pair
					exists, _ := redis.Bool(con.Do("EXISTS", key))
					if exists {
						continue
					}

					// insert Redis
					con.Do("SET", key, 1)
					con.Do("EXPIRE", key, 600)

					// set notices(map)
					notices[pair] = diff
				}
			}

			if len(notices) == 0 {
				continue
			}

			// Send Mail using notices(map)
			var body string
			for pair, diff := range notices {
				buy := strings.Split(pair, "_")[0]
				sell := strings.Split(pair, "_")[1]
				body = body + buy + "の売り板と" + sell + "の買い板で " + strconv.FormatFloat(diff, 'f', 2, 64) + "% の差が発生しています\n"
			}

			body += "\n\n全取引所のリアルタイム板情報 https://kepler.landau.jp/"

			to := mail.Address{"あなた", val[1].(string)}
			title := "差が発生しました"

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

			err := smtp.SendMail(
				"smtp.gmail.com:587",
				auth,
				from.Address,
				[]string{to.Address},
				[]byte(message),
			)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
