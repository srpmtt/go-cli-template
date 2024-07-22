package app

import "fmt"

func (a *MyApp) HandleCommand(command string) {
	switch command {
	case "command1":
		a.Command1()
	case "command2":
		a.Command2()
	default:
		fmt.Println("Unknown command")
	}
}
