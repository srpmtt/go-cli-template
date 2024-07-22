package main

import (
	"fmt"
	"myApp/internal/app"
	"os"
)

func main() {
	myApp := app.NewMyApp()

	if len(os.Args) < 2 {
		fmt.Println("Please provide a command (e.g., command1)")
		return
	}

	command := os.Args[1]
	myApp.HandleCommand(command)
}
