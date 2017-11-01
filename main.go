package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	simplejson "github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	"github.com/landaujp/archimedes/exchanges"
)

//go:generate go-bindata config/config.toml

type Config struct {
	DB struct {
		User     string
		Password string
		Port     int
	}
}

type Exchange interface {
	GetLast() float64
	GetTimestamp() int64
	SetJson(*simplejson.Json)
}

func main() {
	var config Config

	data, err := Asset("config/config.toml")
	_, err = toml.Decode(string(data), &config)
	if err != nil {
		panic(err)
	}

	db, err := sql.Open("mysql", config.DB.User+":"+config.DB.Password+"@tcp(127.0.0.1:"+strconv.Itoa(config.DB.Port)+")/market?parseTime=true&loc=Asia%2FTokyo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	flag.Parse()
	argument := flag.Args()[0]
	// argument := "zaif"

	var url string
	var ex Exchange
	switch argument {
	case "coincheck":
		fmt.Println(argument)
		url = "https://coincheck.com/api/ticker"
		ex = &exchanges.Coincheck{}
	case "zaif":
		fmt.Println(argument)
		url = "https://api.zaif.jp/api/1/ticker/btc_jpy"
		ex = &exchanges.Zaif{}
	default:
		fmt.Println("There is no exchanges...")
		return
	}

	var Etag string
	for {
		time.Sleep(2 * time.Second) // 2秒待つ

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Set("if-none-match", Etag)
		client := new(http.Client)
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println(err)
			return
		}
		if resp.StatusCode != 200 {
			continue
		}
		Etag = resp.Header["Etag"][0]
		body, err := ioutil.ReadAll(resp.Body)

		json, err := simplejson.NewJson(body)
		ex.SetJson(json)
		fmt.Println(ex.GetLast())
		fmt.Println(ex.GetTimestamp())
		// fmt.Println(reflect.TypeOf(json.Get("last").MustFloat64()))
		// data := Jsondata{}
		// if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		// 	panic(err)
		// }
		// _, err = db.Exec("INSERT INTO coincheck (last,timestamp,created_at) VALUES (?,?,?)", data.Last, data.Timestamp, time.Now())
		// if err != nil {
		// 	panic(err.Error())
		// }
		resp.Body.Close()
		os.Exit(0)
	}
}
