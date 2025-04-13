package utils

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/common-nighthawk/go-figure"
)

func ClearConsole() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}

func Header() {
	header := figure.NewColorFigure("MYAPP", "doom", "cyan", true)
	header.Print()
	println("\n")
}
