package depth

import (
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

type Quoine struct {
	Json *simplejson.Json
}

func (b *Quoine) GetDepth() int64 {
	a := b.Json.Get("last_traded_price").MustFloat64()
	return int64(a)
}

func (b *Quoine) GetTimestamp() int64 {
	return int64(time.Now().Unix())
}

func (b *Quoine) SetJson(json *simplejson.Json) {
	b.Json = json
}
