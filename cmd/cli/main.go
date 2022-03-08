package main

import (
	"github.com/digitalex70/celeritas"
)

const version = "0.0.1"

var cel celeritas.celeritas

func main(){
	arg1, arg2, arg3, err := validateInput()
	if err != nil {
		exitGracefully()
	}
}

func exitGracefully() {
}

func validateInput() (string, string, string, error) {
	return "", "", "", nil
}