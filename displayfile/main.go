package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args[1:]

	count := 0
	for range arg {
		count++
	}

	if count == 1 {
		file, err := ioutil.ReadFile(arg[0])

		if err != nil {
			fmt.Println("File name missing")
		}

		fmt.Println(string(file))

	} else if count == 0 {
		fmt.Println("File name missing")
	} else {
		fmt.Println("Too many arguments")
	}

}
