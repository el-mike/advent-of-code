package common

import (
	"os"
	"os/exec"
)

func ClearTerminal() {
	cmd := exec.Command("clear")

	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		panic(err)
	}
}
