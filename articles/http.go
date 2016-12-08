package main

import (
	"encoding/xml"
	"net/http"
)

type rssFeed struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		XMLName xml.Name `xml:"channel"`
		Items   []struct {
			XMLName xml.Name `xml:"item"`
			Title   string   `xml:"title"`
			GUID    string   `xml:"guid"`
			Link    string   `xml:"link"`
		} `xml:"item"`
	} `xml:"channel"`
}

func fetchRSSFeed(u string) (rssFeed, error) {
	var feed rssFeed

	resp, err := http.Get(u)
	if err != nil {
		return feed, err
	}
	defer resp.Body.Close()

	dec := xml.NewDecoder(resp.Body)
	if err := dec.Decode(&feed); err != nil {
		return feed, err
	}

	return feed, nil
}
