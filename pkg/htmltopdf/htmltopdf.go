package htmltopdf

import (
	"crypto/rand"
	"fmt"
	"io"
	"os/exec"
	"runtime"
)

func Convert(fp string) (string, error) {
	var o string
	var e error
	var com string
	ui, err := newUUID()
	if err != nil {
		return o, err
	}
	uid := fmt.Sprintf("%s", ui)
	a := []string{
		"-q",
		"-L",
		"0",
		"-R",
		"0",
		"-T",
		"0",
		fp,
		"/tmp/" + uid + ".pdf",
	}
	if runtime.GOOS == "linux" {
		com = "xvfb-run"
		a = append([]string{"wkhtmltopdf"}, a...)
	} else if runtime.GOOS == "darwin" {
		a = append([]string{"--minimum-font-size", "50"}, a...)
		com = "wkhtmltopdf"
	}
	cmd := exec.Command(com, a...)
	cerr := cmd.Run()
	if cerr != nil && cerr.Error() != "exit status 1" {
		return o, cerr
	}
	o = "/tmp/" + uid + ".pdf"
	return o, e
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
