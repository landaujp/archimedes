package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
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
	cur, err := os.Executable()
	fmt.Println(filepath.Join(filepath.Dir(cur), "config.toml"))
	_, err = toml.DecodeFile(filepath.Join(filepath.Dir(cur), "config.toml"), &config)
	// _, err = toml.DecodeFile("config.toml", &config)
	if err != nil {
		// Error Handling
	}
	db, err := sql.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp(127.0.0.1:"+strconv.Itoa(config.DB.Port)+")/market")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	cc := "https://coincheck.com/api/ticker"

	var Etag string
	for {
		req, _ := http.NewRequest("GET", cc, nil)
		req.Header.Set("if-none-match", Etag)
		client := new(http.Client)
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println(err)
			return
		}
		if resp.StatusCode != 200 {
			time.Sleep(2 * time.Second) // 2秒待つ
			continue
		}
		Etag = resp.Header["Etag"][0]

		data := Jsondata{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			panic(err)
		}
		_, err = db.Exec("INSERT INTO coincheck (last,timestamp,created_at) VALUES (?,?,?)", data.Last, data.Timestamp, time.Now())
		if err != nil {
			panic(err.Error())
		}
		resp.Body.Close()

		time.Sleep(2 * time.Second) // 2秒待つ
	}
}
