package spigot_test

import (
	"fmt"
	"testing"

	spigot "github.com/momonaguilar/go-spigot-algo"
)

func TestGetPiPrecision10(t *testing.T) {
	fmt.Print("INFO: Starting test...")
	spigot.Precision = 10
	pi := spigot.GetPi()
	fmt.Println("TEST: Pi value is: ", pi)
}

func TestGetPiPrecision20(t *testing.T) {
	fmt.Print("INFO: Starting test...")
	spigot.Precision = 20
	pi := spigot.GetPi()
	fmt.Println("TEST: Pi value is: ", pi)
}
