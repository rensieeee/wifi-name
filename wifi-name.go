package wifiname

import (
	"io"
	"os/exec"
	"runtime"
	"strings"
)

const osxCmd = "networksetup -listpreferredwirelessnetworks en0 | sed -n '2 p' | tr -d '\t'"
const linuxCmd = "iwgetid"
const linuxArgs = "--raw"

func WifiName() string {
	platform := runtime.GOOS
	if platform == "darwin" {
		return forOSX()
	} else if platform == "win32" {
		// TODO for Windows
		return ""
	} else {
		// TODO for Linux
		return forLinux()
	}
}

func forLinux() string {
	cmd := exec.Command(linuxCmd, linuxArgs)
	stdout, err := cmd.StdoutPipe()
	panicIf(err)

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	defer cmd.Wait()

	var str string

	if b, err := io.ReadAll(stdout); err == nil {
		str += (string(b) + "\n")
	}

	name := strings.Replace(str, "\n", "", -1)
	return name
}

func forOSX() string {
	cmd := exec.Command("bash", "-c", osxCmd)

	stdout, err := cmd.StdoutPipe()
	panicIf(err)

	// start the command after having set up the pipe
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	defer cmd.Wait()

	var str string
	if b, err := io.ReadAll(stdout); err == nil {
		str = string(b)
	}

	return strings.TrimSpace(str)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
