package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	values := url.Values{
		"query": {"Hello World"},
	}
	resp, err := http.Get("http://localhost:18888" + "?" + values.Encode())
	if err != nil {
		panic(err)
	}

	log.Println("status: ", resp.Status)
	log.Println("statusCode: ", resp.StatusCode)
	log.Println("Header: ", resp.Header)
	log.Println("Header: ", resp.Header.Get("Content-Length"))

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	log.Println(string(body))

}
