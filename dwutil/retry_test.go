package dwutil

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestRetryDownloader(t *testing.T) {
	testCases := []int{
		0, // timesToFail
		3,
	}

	for _, timesToFail := range testCases {
		numOfFailures := 0

		// Build a MockDownloader that will fail `timesToFail` times
		md := &MockDownloader{
			DownloadFn: func(url string) (io.Reader, error) {
				if numOfFailures < timesToFail {
					numOfFailures++
					return nil, errors.New("bad")
				}
				return strings.NewReader("good"), nil
			},
		}

		// Create a RetryDownloader that will retry `timesToFail + 1` times before giving up
		d := NewRetryDownloader(md, timesToFail+1, 0)

		r, err := d.Download("http://example.com")
		if err != nil {
			t.Errorf("Exceeded number of expected failures: %d failures", numOfFailures)
		}

		if numOfFailures != timesToFail {
			t.Errorf("Wrong number of failures. Expected %d but got %d", timesToFail, numOfFailures)
		}

		if res, _ := ioutil.ReadAll(r); string(res) != "good" {
			t.Errorf("Failed to read correct response: %s", res)
		}
	}
}
