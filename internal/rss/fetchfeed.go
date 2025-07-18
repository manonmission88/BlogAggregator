package rss

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
)

// fetch the feed from the given url
func FetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedUrl, nil) // requests -> can terminate early
	if err != nil {
		return &RSSFeed{}, err
	}
	// making sure right user calling the web
	req.Header.Set("User-Agent", "gator")
	client := Client{
		httpClient: http.Client{},
	}
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return &RSSFeed{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &RSSFeed{}, err
	}
	// store the data here
	feedData := &RSSFeed{}
	err = xml.Unmarshal(data, &feedData) // bytes into xml
	if err != nil {
		return &RSSFeed{}, err
	}
	return feedData, nil

}
