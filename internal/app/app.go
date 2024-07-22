package app

import (
	"fmt"
)

type MyApp struct {
}

func NewMyApp() *MyApp {
	return &MyApp{}
}

func (a *MyApp) Execute() {
	fmt.Println("Executing myApp...")
}
