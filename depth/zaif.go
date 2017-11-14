package depth

import (
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

type Zaif struct {
	Json *simplejson.Json
}

func (z *Zaif) GetDepth() int64 {
	a := z.Json.Get("last").MustFloat64()
	return int64(a)
}

func (z *Zaif) GetTimestamp() int64 {
	return int64(time.Now().Unix())
}

func (z *Zaif) SetJson(json *simplejson.Json) {
	z.Json = json
}
