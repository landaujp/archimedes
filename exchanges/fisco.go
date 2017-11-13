package exchanges

import (
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

type Fisco struct {
	Json *simplejson.Json
}

func (b *Fisco) GetLast() int64 {
	a := b.Json.Get("last").MustFloat64()
	return int64(a)
}

func (b *Fisco) GetTimestamp() int64 {
	return int64(time.Now().Unix())
}

func (b *Fisco) SetJson(json *simplejson.Json) {
	b.Json = json
}
