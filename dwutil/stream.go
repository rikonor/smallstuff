package dwutil

import (
	"fmt"
	"io"
	"net/http"
)

type StreamDownloader struct {
}

func (d *StreamDownloader) Download(url string) (io.Reader, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	// Check request succeeded with 200
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad status code: %d", res.StatusCode)
	}

	return res.Body, nil
}
