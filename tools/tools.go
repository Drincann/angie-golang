package tools

import (
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) (stdout string, err error) {
	stdoutBytes, e := exec.Command(
		map[string]string{"windows": "cmd /c start", "darwin": "open"}[runtime.GOOS],
		url,
	).Output()

	if e != nil {
		stdout = ""
		err = e
		return
	}
	stdout = string(stdoutBytes)
	return
}
