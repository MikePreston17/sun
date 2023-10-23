package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	res, err := http.Get("https://api.weatherapi.com/v1/current.json?key=4d31bbdd7f954dcb8b600541232310&q=Dallas&aqi=no")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
