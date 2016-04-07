package amazing

import (
	"net/http"
	"time"
)

func newTimeoutClient(timeout time.Duration) *http.Client {
	if timeout <= 0 {
		timeout = 20 * time.Second
	}
	return &http.Client{
		Timeout: timeout,
	}
}
