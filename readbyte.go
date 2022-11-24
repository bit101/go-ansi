// Package ansi allows for advanced terminal text manipulation.
package ansi

import (
	"os"
	"os/exec"
	"strconv"
)

// ReadBytes reads a single byte from stdn.
func ReadBytes(count int) []byte {
	// disable input buffering
	// stty changes terminal line settings
	// -F /dev/tty uses that device instead of stdin
	// cbreak enable special characters: erase, kill, werase, rprnt
	// min 1 = min chars for completed read.
	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", strconv.Itoa(count)).Run()

	// do not display entered characters on the screen
	// see above
	// -echo disables echo
	exec.Command("stty", "-F", "/dev/tty", "-echo").Run()

	var b = make([]byte, count)
	os.Stdin.Read(b)
	return b
}
