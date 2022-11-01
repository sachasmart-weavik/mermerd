package util

import (
	"fmt"
	"os/exec"

	"github.com/fatih/color"
)

func ShowSuccess(fileName string) {

	// execute sh command

	exec.Command("sh", "-c", "echo 'hello'").Run()

	color.Green(fmt.Sprintf(`

✓ Diagram was created successfully (%s)

`, fileName))
}

func ShowError() {
	color.Red(`

X Something went wrong!

`)
}
