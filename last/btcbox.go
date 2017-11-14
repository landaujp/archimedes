package last

import (
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

type Btcbox struct {
	Json *simplejson.Json
}

func (b *Btcbox) GetLast() int64 {
	a := b.Json.Get("last").MustInt64()
	return a
}

func (b *Btcbox) GetTimestamp() int64 {
	return int64(time.Now().Unix())
}

func (b *Btcbox) SetJson(json *simplejson.Json) {
	b.Json = json
}
