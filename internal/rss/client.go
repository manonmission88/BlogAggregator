package rss

import "net/http"

type Client struct {
	httpClient http.Client
}
