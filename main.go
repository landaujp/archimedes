package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
	"time"
)

type Config struct {
	DB struct {
		User     string
		Password string
		Port     int
	}
}

type Jsondata struct {
	Last, Bid, Ask, High, Low, Volume float32
	Timestamp                         int
}

func main() {
	var config Config
	_, err := toml.DecodeFile("config.toml", &config)
	if err != nil {
		// Error Handling
	}
	db, err := sql.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp(127.0.0.1:"+strconv.Itoa(config.DB.Port)+")/market")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	cc := "https://coincheck.com/api/ticker"

	for {
		resp, err := http.Get(cc)

		if err != nil {
			fmt.Println(err)
			return
		}
		if resp.StatusCode != 200 {
			continue
		}

		data := Jsondata{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			panic(err)
		}
		_, err = db.Exec("INSERT INTO coincheck (last,created_at) VALUES (?,?)", data.Last, time.Now())
		if err != nil {
			panic(err.Error())
		}
		resp.Body.Close()

		time.Sleep(1 * time.Second) // 1秒待つ
	}
}
