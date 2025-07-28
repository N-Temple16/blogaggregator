package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("Error fetching RSS feed: %v\n", err)
	}	
	
	req.Header.Set("User-Agent", "gator")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error sending request: %v\n", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v\n", err)
	}

	var r RSSFeed
    err = xml.Unmarshal(body, &r)
    if err != nil {
        return nil, fmt.Errorf("Error unmarshalling data: %v\n", err)
    }

	unescapedTitle := html.UnescapeString(r.Channel.Title)
	unescapedDescription := html.UnescapeString(r.Channel.Description)

	r.Channel.Title = unescapedTitle
	r.Channel.Description = unescapedDescription

	for idx, item := range r.Channel.Item {
		title := html.UnescapeString(item.Title)
		desc := html.UnescapeString(item.Description)

		r.Channel.Item[idx].Title = title
    	r.Channel.Item[idx].Description = desc
	}

	return &r, nil
}