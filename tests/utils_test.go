package utils

import (
	"myApp/pkg/utils"
	"testing"
)

func TestHello(t *testing.T) {
	expectedOutput := "Hello from utils package!"
	if got := utils.Hello(); got != expectedOutput {
		t.Errorf("SayHello() = %s; want %s", got, expectedOutput)
	}
}
