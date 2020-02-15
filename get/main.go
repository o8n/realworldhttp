package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values {
		"query": {"Hello World"},
	}
	resp, _ := http.Get("http://localhost:18888" + "?" + values.Encode())
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

	// log.Println("Status:", resp.Status)
	// log.Println("StatusCode:", resp.Header)
}
