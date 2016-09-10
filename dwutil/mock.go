package dwutil

import (
	"io"
	"strings"
)

// MockDownloader returns injected responses
type MockDownloader struct {
	DownloadFn        func(url string) (io.Reader, error)
	DownloadFnInvoked bool
}

// Download invokes the wrapped DownloadFn
func (d *MockDownloader) Download(url string) (io.Reader, error) {
	d.DownloadFnInvoked = true
	return d.DownloadFn(url)
}

// TextDownloader is a mock downloader returning a text response
func TextDownloader(text string) MockDownloader {
	return MockDownloader{
		DownloadFn: func(url string) (io.Reader, error) {
			return strings.NewReader(text), nil
		},
	}
}

// NOOPDownloader returns an empty response
var NOOPDownloader = TextDownloader("")
