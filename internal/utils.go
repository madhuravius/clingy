package internal

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func CheckMagickBinary() error {
	if _, err := exec.LookPath("magick"); os.IsNotExist(err) {
		return errors.New("error: magick binary not found")
	}
	return nil
}

func ClearTerminal() {
	fmt.Print("\033[H\033[2J") // taken from: https://stackoverflow.com/a/22892171
}
