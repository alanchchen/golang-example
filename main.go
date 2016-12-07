package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	url = flag.String("url", "", "The URL")
	port = flag.String("port", "8080", "The listening port")
)

func handler(w http.ResponseWriter, r *http.Request) {
	getURL := strings.TrimSpace(*url)
	if getURL == "" {
		fmt.Println("URL is empty")
		return
	}

	resp, err := http.Get(getURL)
	if err != nil {
		fmt.Printf("Failed to GET %s, %v\n", getURL, err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body,", err)
		return
	}

	fmt.Println(string(body))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
