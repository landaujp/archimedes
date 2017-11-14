package main

import (
	"database/sql"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/BurntSushi/toml"
	simplejson "github.com/bitly/go-simplejson"
	_ "github.com/go-sql-driver/mysql"
	"github.com/landaujp/archimedes/depth"
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
	GetDepth() string
	SetJson(*simplejson.Json)
}

func main() {

	var config Config

	data, _ := Asset("config/config.toml")
	_, err := toml.Decode(string(data), &config)
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

	var ex Exchange
	var url string
	switch argument {
	case "bitflyer":
		ex = &depth.Bitflyer{}
		url = "https://api.bitflyer.jp/v1/getboard"
	case "coincheck":
		ex = &depth.Coincheck{}
		url = "https://coincheck.com/api/order_books"
	case "zaif":
		ex = &depth.Zaif{}
		url = "https://api.zaif.jp/api/1/depth/btc_jpy"
	case "bitbank":
		ex = &depth.Bitbank{}
		url = "https://public.bitbank.cc/btc_jpy/depth"
	case "kraken":
		ex = &depth.Kraken{}
		url = "https://api.kraken.com/0/public/Depth?pair=XBTJPY"
	case "quoine":
		ex = &depth.Quoine{}
		url = "https://api.quoine.com/products/5/price_levels"
	case "btcbox":
		ex = &depth.Btcbox{}
		url = "https://www.btcbox.co.jp/api/v1/depth/"
	case "fisco":
		ex = &depth.Fisco{}
		url = "https://api.fcce.jp/api/1/depth/btc_jpy"
	default:
		fmt.Println("There is no exchanges...")
		return
	}

	var Etag string
	for {
		// time.Sleep(2 * time.Second) // 2秒待つ

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
		jsonObj, _ := simplejson.NewJson(body)
		ex.SetJson(jsonObj)
		jsonString := ex.GetDepth()
		_, err = db.Exec("UPDATE market SET depth = ? WHERE exchange = '"+argument+"' ", jsonString)
		if err != nil {
			panic(err.Error())
		}
		break

		// // asks := json.Get("asks").MustArray()
		// bids, _ := jsonObj.Get("bids").Array()

		// sort_bids := make(map[int]float64)

		// for _, arr := range bids {
		// 	v := arr.([]interface{})
		// 	v1 := strings.Split(v[0].(string), ".")[0]
		// 	v2 := v[1].(string)
		// 	vv1, _ := strconv.Atoi(v1)
		// 	vv2, _ := strconv.ParseFloat(v2, 64)
		// 	sort_bids[vv1] = vv2
		// }
		// // sort desc
		// var keys []int
		// for k := range sort_bids {
		// 	keys = append(keys, k)
		// }
		// sort.Sort(sort.Reverse(sort.IntSlice(keys)))
		// keys = keys[len(keys)-10:]

		// type Pair struct {
		// 	Price int     `json:"price"`
		// 	Size  float64 `json:"size"`
		// }
		// var res_bids []Pair

		// for _, s := range keys {
		// 	res_bids = append(res_bids, Pair{s, sort_bids[s]})
		// }
		// outputJson, err := json.Marshal(res_bids)
		// fmt.Println(string(outputJson))
		resp.Body.Close()
	}
}
