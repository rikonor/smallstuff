package zip

import (
	"archive/zip"
	"bytes"
	"io/ioutil"
	"testing"
)

func TestGetZIPFileReader(t *testing.T) {
	zfc := FileCollection{
		FileContent{"tmp1", []byte("1,2,3")},
		FileContent{"tmp2", []byte("5,6,7")},
	}

	f, err := GetZIPFileReader(zfc)
	if err != nil {
		t.Fatal(err)
	}

	r, err := zip.NewReader(f, int64(f.Len()))
	if err != nil {
		t.Fatal(err)
	}

	for i, zf := range r.File {
		if zf.Name != zfc[i].Name {
			t.Fatal(err)
		}

		zfr, err := zf.Open()
		if err != nil {
			t.Fatal(err)
		}

		zfrb, err := ioutil.ReadAll(zfr)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(zfrb, zfc[i].Body) {
			t.Fatal("ZIP file content corrupted")
		}
	}
}
