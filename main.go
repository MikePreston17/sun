package main

import "net/http"

func main() {
	// fmt.Println("ready to go")

	res, err := http.Get("https://api.weatherapi.com/v1/current.json?key=4d31bbdd7f954dcb8b600541232310&q=Dallas&aqi=no")

	if err != nil {
		panic(err)
	}

}
