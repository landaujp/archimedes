package exchanges

import simplejson "github.com/bitly/go-simplejson"

type Coincheck struct {
	Json *simplejson.Json
}

func (c *Coincheck) GetLast() float64 {
	a := c.Json.Get("last").MustFloat64()
	return a
}

func (c *Coincheck) GetTimestamp() int64 {
	a := c.Json.Get("timestamp").MustInt64()
	return a
}

func (c *Coincheck) SetJson(json *simplejson.Json) {
	c.Json = json
}
