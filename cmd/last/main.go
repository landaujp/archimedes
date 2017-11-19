package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/BurntSushi/toml"
	simplejson "github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	"github.com/landaujp/archimedes/last"
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

type Exchange interface {
	GetLast() int64
	GetTimestamp() int64
	SetJson(*simplejson.Json)
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

	flag.Parse()
	argument := flag.Args()[0]

	var ex Exchange
	var url string
	table := argument
	switch argument {
	case "bitflyer":
		ex = &last.Bitflyer{}
		url = "https://api.bitflyer.jp/v1/ticker?product_code=BTC_JPY"
	case "coincheck":
		ex = &last.Coincheck{}
		url = "https://coincheck.com/api/ticker"
	case "zaif":
		ex = &last.Zaif{}
		url = "https://api.zaif.jp/api/1/ticker/btc_jpy"
	case "bitbank":
		ex = &last.Bitbank{}
		url = "https://public.bitbank.cc/btc_jpy/ticker"
	case "kraken":
		ex = &last.Kraken{}
		url = "https://api.kraken.com/0/public/Ticker?pair=XBTJPY"
	case "quoine":
		ex = &last.Quoine{}
		url = "https://api.quoine.com/products/5"
	case "btcbox":
		ex = &last.Btcbox{}
		url = "https://www.btcbox.co.jp/api/v1/ticker/"
	case "fisco":
		ex = &last.Fisco{}
		url = "https://api.fcce.jp/api/1/ticker/btc_jpy"
	default:
		fmt.Println("There is no exchanges...")
		return
	}

	var Etag string
	for {
		time.Sleep(5 * time.Second) // 2秒待つ

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

		if val, ok := resp.Header["Etag"]; ok {
			Etag = val[0]
		}

		body, _ := ioutil.ReadAll(resp.Body)
		json, _ := simplejson.NewJson(body)
		ex.SetJson(json)
		_, err = db.Exec("INSERT INTO "+table+" (last,timestamp,created_at) VALUES (?,?,?)", ex.GetLast(), ex.GetTimestamp(), time.Now())
		if err != nil {
			panic(err.Error())
		}
		_, err = db.Exec("INSERT INTO market (exchange,last,timestamp,created_at) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE last = ?,timestamp = ?, created_at = ?", table, ex.GetLast(), ex.GetTimestamp(), time.Now(), ex.GetLast(), ex.GetTimestamp(), time.Now())
		if err != nil {
			panic(err.Error())
		}

		resp.Body.Close()
	}
}
