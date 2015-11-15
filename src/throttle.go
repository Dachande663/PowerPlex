package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

const maxConcurrency = 10

var httpThrottle = make(chan bool, maxConcurrency)

// Make a request with max concurrency
func makeRequest(url string) ([]byte, error) {

	httpThrottle <- true // wait for slot

	resp, err := makeActualRequest(url)

	time.Sleep(time.Second)

	<-httpThrottle // release slot

	return resp, err

}

// Make an API request
func makeActualRequest(url string) ([]byte, error) {

	// time.Sleep(time.Second)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil

}
