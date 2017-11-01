package exchanges

import (
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

type Zaif struct {
	Json *simplejson.Json
}

func (c *Zaif) GetLast() float64 {
	a := c.Json.Get("last").MustFloat64()
	return a
}

func (c *Zaif) GetTimestamp() int64 {
	return int64(time.Now().Unix())
}

func (c *Zaif) SetJson(json *simplejson.Json) {
	c.Json = json
}
