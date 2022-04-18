package spigot_test

import (
	"fmt"
	"testing"

	spigot "github.com/momonaguilar/go-spigot-algo"
)

func TestGetPi(t *testing.T) {
	fmt.Print("INFO: Starting test...")
	spigot.Precision = 10
	pi := spigot.GetPi()
	fmt.Println("TEST: Pi value is: ", pi)
}
