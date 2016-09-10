package gzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"testing"
)

func TestGetGZIPFileReader(t *testing.T) {
	testCase := []byte("test-content")

	f, err := GetGZIPFileReader(testCase)
	if err != nil {
		t.Fatal(err)
	}

	r, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(b, testCase) {
		t.Fatal("gzip file content corrupted")
	}
}
