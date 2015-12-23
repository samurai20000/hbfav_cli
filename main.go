package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Item struct {
	Title string `xml:"title"`
	URL   string `xml:"link"`
	User  string `xml:"creator"`
}

type Entries struct {
	Entries []Item `xml:"item"`
}

func main() {
	user := "samurai20000"
	url := buildURL(user)
	feed := getRSSFeed(url)
	results := parseItems(feed)
	for _, entry := range results.Entries {
		fmt.Println(entry)
	}
	// contents, err := ioutil.ReadAll(results)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%s\n", string(contents))
}

func buildURL(user string) string {
	return fmt.Sprintf("http://b.hatena.ne.jp/%s/favorite.rss", url.QueryEscape(user))
}

func getRSSFeed(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	feed, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	return feed
}

func parseItems(feed []byte) Entries {
	entries := Entries{}

	err := xml.Unmarshal(feed, &entries)
	if err != nil {
		fmt.Println(err)
	}

	return entries
}
