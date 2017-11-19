package main

import (
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
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

func main() {

	var config Config

	data, _ := Asset("config/config.toml")
	_, err := toml.Decode(string(data), &config)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp("+config.DB.Host+":"+strconv.Itoa(config.DB.Port)+")/"+config.DB.Database+"?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	for {
		time.Sleep(10 * time.Second) // 2秒待つ

		// あとでtimestampを使って、一定時間変更されていない場合のフィルタリングを入れる
		rows, err := db.Query("SELECT exchange, last, timestamp FROM market")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		ex := map[string]float64{}
		for rows.Next() {
			var exchange string
			var last float64
			var timestamp int
			if err := rows.Scan(&exchange, &last, &timestamp); err != nil {
				log.Fatal(err)
			}
			ex[exchange] = last
		}

		for k1, l1 := range ex {
			for k2, l2 := range ex {
				if k1 == k2 {
					continue
				}
				rate := diffRate(l1, l2)
				_, err = db.Exec("UPDATE alert SET diff = ?, created_at = ? WHERE exchange1 = ? AND exchange2 = ?", rate, time.Now(), k1, k2)
				if err != nil {
					panic(err.Error())
				}
			}
			delete(ex, k1)
		}
	}
}
