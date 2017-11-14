package depth

import (
	"encoding/json"
	"sort"
	"strconv"
	"strings"

	simplejson "github.com/bitly/go-simplejson"
)

type Coincheck struct {
	SimpleJson *simplejson.Json
}

func (c *Coincheck) GetDepth() string {

	type Pair struct {
		Price int     `json:"price"`
		Size  float64 `json:"size"`
	}

	type SaveJson struct {
		Bids []Pair `json:"bids"`
		Asks []Pair `json:"asks"`
	}

	/*
	 * bid
	 */
	bids, _ := c.SimpleJson.Get("bids").Array()

	map_bids := make(map[int]float64)

	for _, arr := range bids {
		v := arr.([]interface{})
		v1 := strings.Split(v[0].(string), ".")[0]
		v2 := v[1].(string)
		vv1, _ := strconv.Atoi(v1)
		vv2, _ := strconv.ParseFloat(v2, 64)
		map_bids[vv1] = vv2
	}
	// sort desc
	var keys []int
	for k := range map_bids {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	keys = keys[0:10]

	var res_bids []Pair

	for _, s := range keys {
		res_bids = append(res_bids, Pair{s, map_bids[s]})
	}

	/*
	 * ask
	 */
	asks, _ := c.SimpleJson.Get("asks").Array()

	map_asks := make(map[int]float64)

	for _, arr := range asks {
		v := arr.([]interface{})
		v1 := strings.Split(v[0].(string), ".")[0]
		v2 := v[1].(string)
		vv1, _ := strconv.Atoi(v1)
		vv2, _ := strconv.ParseFloat(v2, 64)
		map_asks[vv1] = vv2
	}
	// sort desc
	var a_keys []int
	for k := range map_asks {
		a_keys = append(a_keys, k)
	}
	sort.Ints(a_keys)
	a_keys = a_keys[0:10]
	sort.Sort(sort.Reverse(sort.IntSlice(a_keys)))

	var res_asks []Pair

	for _, s := range a_keys {
		res_asks = append(res_asks, Pair{s, map_asks[s]})
	}

	saveJson := SaveJson{res_bids, res_asks}
	outputJson, _ := json.Marshal(saveJson)

	return string(outputJson)
}

func (c *Coincheck) SetJson(json *simplejson.Json) {
	c.SimpleJson = json
}
