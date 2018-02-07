package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const scrapedFeed = "https://www.washingtonpost.com/news-sitemap-index.xml"

func main() {
	resp, err := http.Get(scrapedFeed)

	if err != nil {
		panic(err.Error)
	}

	defer resp.Body.Close()

	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)

	fmt.Println(stringBody)
}
