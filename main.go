package main

import (
	"fmt"
	"time"
	"net/http"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Jsondata struct {
	Last, Bid, Ask, High, Low, Volume float32
	Timestamp int
}

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/market")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される

	cc := "https://coincheck.com/api/ticker"

	for {
		resp, err := http.Get(cc)

		if err != nil {
			fmt.Println(err)
			return
		}

		data := Jsondata{}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			panic(err)
		}
		_, err = db.Exec("INSERT INTO coincheck (last) VALUES (?)", data.Last)
		if err != nil {
			panic(err.Error())
		}
		resp.Body.Close()

		time.Sleep(1 * time.Second) // 1秒待つ
	}
}
