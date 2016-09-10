package dwutil

import "io"

// RetryDownloader is a wrapper around a Downloader
// that will retry to Download in case of failure
type RetryDownloader struct {
	d       Downloader
	retries int
}

// NewRetryDownloader creates a new RetryDownloader given a Downlaoder
func NewRetryDownloader(d Downloader, retries int) Downloader {
	return &RetryDownloader{d: d, retries: retries}
}

// Download tries to download the given url
func (d *RetryDownloader) Download(downloadURL string) (io.Reader, error) {
	var err error
	var r io.Reader

	for d.retries > 0 {
		r, err = d.d.Download(downloadURL)
		if err != nil {
			d.retries--
			continue
		}

		return r, nil
	}

	return nil, err
}
