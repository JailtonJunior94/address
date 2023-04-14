package interfaces

import (
	"net/http"
	"time"

	"github.com/jailtonjunior94/address/configs"
)

type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func NewHttpClient(config *configs.Config) HttpClient {
	client := &http.Client{
		Timeout: time.Duration(config.HttpClientTimeout) * time.Second,
	}
	return client
}
