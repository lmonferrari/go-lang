package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	var url string
	flag.StringVar(&url, "url", "https://kickasstorrents.to/dungeons-and-dragons-honor-among-thieves-2023-1080p-amzn-webrip-1600mb-dd5-1-x264-galaxyrg-t5643547.html", "url to catch the hash")
	flag.Parse()

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)

	regex := regexp.MustCompile("Torrent hash: [0-9A-F]{40}")
	hash := regex.FindString(string(body))

	fmt.Println(hash)

	defer resp.Body.Close()
}
