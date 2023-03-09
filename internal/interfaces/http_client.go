package interfaces

import (
	"net/http"
	"time"
)

type IHttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHttpClient() IHttpClient {
	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	return client
}
