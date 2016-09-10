package dwutil

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Downloader interface {
	Download(url string) (io.Reader, error)
}

type HTTPDownloader struct {
}

func (d *HTTPDownloader) Download(downloadURL string) (io.Reader, error) {
	u, err := url.Parse(downloadURL)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(u.Path, "/")
	fName := parts[len(parts)-1]

	res, err := http.Get(downloadURL)
	if err != nil {
		return nil, err
	}

	// Check request succeeded with 200
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad status code: %d", res.StatusCode)
	}

	// Stream the response body into a file
	f, err := os.Create(fName)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(f, res.Body)
	if err != nil {
		return nil, err
	}

	res.Body.Close()
	f.Close()

	f, err = os.Open(fName)
	if err != nil {
		return nil, err
	}

	return f, nil
}
