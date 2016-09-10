package dwutil

import (
	"io"
	"time"
)

// RetryDownloader is a wrapper around a Downloader
// that will retry to Download in case of failure
type RetryDownloader struct {
	d       Downloader
	retries int           // number of retries
	delay   time.Duration // time between retries
}

// NewRetryDownloader creates a new RetryDownloader given a Downlaoder
func NewRetryDownloader(d Downloader, retries int, delay time.Duration) Downloader {
	return &RetryDownloader{d: d, retries: retries, delay: delay}
}

// Download tries to download the given url
func (d *RetryDownloader) Download(downloadURL string) (io.Reader, error) {
	var err error
	var r io.Reader

	for d.retries > 0 {
		r, err = d.d.Download(downloadURL)
		if err != nil {
			d.retries--
			time.Sleep(d.delay)
			continue
		}

		return r, nil
	}

	return nil, err
}
