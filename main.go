package main

import (
	"fmt"
	"io"
	"net/http"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:location`
}

func main() {
	res, err := http.Get("http://api.weatherapi.com/v1/forecast.json?key=a7bf715a8d8f401f9fe123814241501&q=Istanbul&days=1&aqi=no&alerts=no")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather api error")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
