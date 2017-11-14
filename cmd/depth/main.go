package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	simplejson "github.com/bitly/go-simplejson"
)

type Exchange interface {
	GetLast() int64
	GetTimestamp() int64
	SetJson(*simplejson.Json)
}

func main() {
	flag.Parse()
	argument := flag.Args()[0]

	var url string
	switch argument {
	case "bitflyer":
		url = "https://api.bitflyer.jp/v1/getboard"
	case "coincheck":
		url = "https://coincheck.com/api/order_books"
	case "zaif":
		url = "https://api.zaif.jp/api/1/ticker/btc_jpy"
	case "bitbank":
		url = "https://public.bitbank.cc/btc_jpy/ticker"
	case "kraken":
		url = "https://api.kraken.com/0/public/Ticker?pair=XBTJPY"
	case "quoine":
		url = "https://api.quoine.com/products/5"
	case "btcbox":
		url = "https://www.btcbox.co.jp/api/v1/ticker/"
	case "fisco":
		url = "https://api.fcce.jp/api/1/ticker/btc_jpy"
	default:
		fmt.Println("There is no exchanges...")
		return
	}

	var Etag string
	// for {
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
		// continue
	}

	if val, ok := resp.Header["Etag"]; ok {
		Etag = val[0]
	}

	body, _ := ioutil.ReadAll(resp.Body)
	json, _ := simplejson.NewJson(body)
	// fmt.Println(json)
	// asks := json.Get("asks").MustArray()
	bids := json.Get("bids").MustArray()
	fmt.Println(reflect.TypeOf(bids))
	fmt.Println(bids)

	// jsonString, err := json.Marshal(bids)
	// fmt.Println(jsonString)

	resp.Body.Close()
	// }
}
