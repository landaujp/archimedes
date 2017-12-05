package main

import (
	"database/sql"
	"strconv"
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

var exs = []string{"coincheck", "bitflyer", "bitbank", "btcbox", "fisco", "zaif", "quoine", "kraken"}

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

	dboption := redis.DialDatabase(0)
	con, err := redis.Dial("tcp", "127.0.0.1:6379", dboption)
	if err != nil {
		// handle error
	}
	defer con.Close()

	for {
		var keys_bid []interface{}
		var keys_ask []interface{}
		for _, k := range exs {
			keys_bid = append(keys_bid, k+":bid")
			keys_ask = append(keys_ask, k+":ask")
		}
		bids, _ := redis.Ints(con.Do("MGET", keys_bid...))
		asks, _ := redis.Ints(con.Do("MGET", keys_ask...))

		for ia, ex_a := range exs {
			for ib, ex_b := range exs {
				if ex_a == ex_b {
					continue
				}
				if bids[ib] != 0 && asks[ia] != 0 {
					rate := 100 * (float64(bids[ib])/float64(asks[ia]) - 1)
					con.Send("HSET", "alert", ex_a+"_"+ex_b, rate)

					_, err = db.Exec("INSERT INTO diff_log (ex_ask, ex_bid, ask, bid, rate, created_at) VALUES (?, ?, ?, ?, ?, ?)", ex_a, ex_b, asks[ia], bids[ib], rate, time.Now())

				} else {
					con.Send("HDEL", "alert", ex_a+"_"+ex_b)
				}
			}
		}
		con.Do("EXEC")

		time.Sleep(10 * time.Second)
	}
}
