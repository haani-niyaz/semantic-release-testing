package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	gf "github.com/brianvoe/gofakeit/v6"
)

func main() {
	// Feel free to delete this file.
	fmt.Println("Hello Gophers")

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(body))

	fmt.Println(gf.Name())
}
