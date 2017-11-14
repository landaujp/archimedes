package last

import (
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

type Bitflyer struct {
	Json *simplejson.Json
}

func (b *Bitflyer) GetLast() int64 {
	a := b.Json.Get("ltp").MustFloat64()
	return int64(a)
}

func (b *Bitflyer) GetTimestamp() int64 {
	timestamp, _ := b.Json.Get("timestamp").String()
	datetime := strings.Split(timestamp, ".")[0]
	t, _ := time.Parse("2006-01-02T15:04:05", datetime)
	return t.Unix()
}

func (b *Bitflyer) SetJson(json *simplejson.Json) {
	b.Json = json
}
