package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

const httpMaxConcurrency = 10
const httpSleepTimer = time.Millisecond * 250

var httpThrottle = make(chan bool, httpMaxConcurrency)

// Make a request with max concurrency
func makeRequest(url string) ([]byte, error) {

	httpThrottle <- true // wait for slot

	resp, err := makeActualRequest(url)

	time.Sleep(httpSleepTimer) // arbitrary delay

	<-httpThrottle // release slot

	return resp, err

}

// Actually make the API request
func makeActualRequest(url string) ([]byte, error) {

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
