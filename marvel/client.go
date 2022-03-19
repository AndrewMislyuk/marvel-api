package marvel

import (
	"errors"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout can not be zero")
	}

	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &webLogger{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}
