package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args[1:]
	for _, a := range arg {
		if a == "01" || a == "galaxy" || a == "galaxy 01" {
			fmt.Println("Alert!!!")
			break
		}

	}
}
