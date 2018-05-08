package server

import (
	"crypto/rand"
	"fmt"
	"github.com/robertlestak/htmltopdf/pkg/htmltopdf"
	"io"
	"net/http"
	"os"
)

func Convert(w http.ResponseWriter, r *http.Request) {
	ff, _, ferr := r.FormFile("html")
	if ferr != nil {
		http.Error(w, ferr.Error(), http.StatusBadRequest)
		return
	}
	defer ff.Close()
	ui, uerr := newUUID()
	if uerr != nil {
		http.Error(w, uerr.Error(), http.StatusBadRequest)
		return
	}
	fp := "/tmp/" + ui + ".html"
	wf, cerr := os.Create(fp)
	if cerr != nil {
		http.Error(w, cerr.Error(), http.StatusBadRequest)
		return
	}
	io.Copy(wf, ff)
	pf, err := htmltopdf.Convert(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pd, oerr := os.Open(pf)
	if oerr != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer pd.Close()
	io.Copy(w, pd)
	os.Remove(fp)
	os.Remove(pf)
}

func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	uuid[8] = uuid[8]&^0xc0 | 0x80
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}
