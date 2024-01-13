package main

import (
	"encoding/xml"
	"errors"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)

	if err != nil {
		return RSSFeed{}, errors.New("failed to fetch rss feed")
	}

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	if err != nil {
		return RSSFeed{}, err
	}

	rssFeed := RSSFeed{}

	err = xml.Unmarshal(data, &rssFeed)

	if err != nil {
		return RSSFeed{}, err
	}

	return rssFeed, nil
}
