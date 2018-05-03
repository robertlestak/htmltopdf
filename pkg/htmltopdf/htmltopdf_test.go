package htmltopdf

import (
	"os"
	"testing"
)

func TestConvert(t *testing.T) {
	f, err := Convert("../../examples/hello.html")
	if err != nil {
		t.Error(err)
	}
	os.Remove(f)
}
