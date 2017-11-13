package exchanges

import (
	"strconv"
	"strings"
	"time"

	simplejson "github.com/bitly/go-simplejson"
)

type Kraken struct {
	Json *simplejson.Json
}

func (b *Kraken) GetLast() int64 {
	a, _ := b.Json.Get("result").Get("XXBTZJPY").Get("c").StringArray()
	str := strings.Split(a[0], ".")[0]
	res, _ := strconv.ParseInt(str, 10, 64)
	return res
}

func (b *Kraken) GetTimestamp() int64 {
	return int64(time.Now().Unix())
}

func (b *Kraken) SetJson(json *simplejson.Json) {
	b.Json = json
}
