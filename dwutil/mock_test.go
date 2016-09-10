package dwutil

import (
	"errors"
	"io"
	"io/ioutil"
	"strings"
	"testing"
)

func TestMockDownloader(t *testing.T) {
	mockErr := errors.New("error")
	mockRes := strings.NewReader("response")

	d := &MockDownloader{
		DownloadFn: func(url string) (io.Reader, error) {
			return mockRes, mockErr
		},
	}

	r, err := d.Download("fake_url")
	if err != mockErr {
		t.Error("Failed to mock error:", err)
	}
	if r != mockRes {
		t.Error("Failed to mock response")
	}
}

func TestTextDownloader(t *testing.T) {
	d := TextDownloader("test-response")

	r, err := d.Download("fake_url")
	if err != nil {
		t.Error("failed to return text response:", err)
	}

	if res, _ := ioutil.ReadAll(r); string(res) != "test-response" {
		t.Errorf("Failed to return correct response: %s", res)
	}
}
