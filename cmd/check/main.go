package main

import (
	"time"

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

var exs = []string{"coincheck", "bitflyer", "bitbank", "btcbox", "fisco", "zaif", "quoine", "kraken"}

func main() {

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
					rate := float64(bids[ib])/float64(asks[ia]) - 1
					con.Send("HSET", "alert", ex_a+"_"+ex_b, rate)
				} else {
					con.Send("HDEL", "alert", ex_a+"_"+ex_b)
				}
			}
		}
		con.Do("EXEC")

		time.Sleep(10 * time.Second)
	}
}
