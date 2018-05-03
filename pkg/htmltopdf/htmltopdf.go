package htmltopdf

import (
	"fmt"
	"github.com/satori/go.uuid"
	"os/exec"
	"runtime"
)

func Convert(fp string) (string, error) {
	var o string
	var e error
	var com string
	ui := uuid.NewV4()
	uid := fmt.Sprintf("%s", ui)
	a := []string{
		"-q",
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
