package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"

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
	jsonObj, _ := simplejson.NewJson(body)
	// asks := json.Get("asks").MustArray()
	bids, _ := jsonObj.Get("bids").Array()

	sort_bids := make(map[int]float64)

	for _, arr := range bids {
		v := arr.([]interface{})
		v1 := strings.Split(v[0].(string), ".")[0]
		v2 := v[1].(string)
		vv1, _ := strconv.Atoi(v1)
		vv2, _ := strconv.ParseFloat(v2, 64)
		sort_bids[vv1] = vv2
	}
	// sort desc
	var keys []int
	for k := range sort_bids {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	keys = keys[len(keys)-10:]

	type Pair struct {
		Price int     `json:"price"`
		Size  float64 `json:"size"`
	}
	var res_bids []Pair

	for _, s := range keys {
		res_bids = append(res_bids, Pair{s, sort_bids[s]})
	}
	outputJson, err := json.Marshal(res_bids)
	fmt.Println(string(outputJson))
	resp.Body.Close()
	// }
}
