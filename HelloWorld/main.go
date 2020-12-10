package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://api.coinbase.com/v2/prices/spot?currency=usd")
	if err != nil {
		fmt.Printf("Bad Request")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
