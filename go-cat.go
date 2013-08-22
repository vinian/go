package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("The argument is least than 2!\n")
		// read from stdin and output to stdout
	} else {
		fileCat(os.Args[1])
	}
}

func fileCat(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	buffer := make([]byte, 100)

	for n, e := file.Read(buffer); e == nil; n, e = file.Read(buffer) {
		if n > 0 {
			os.Stdout.Write(buffer[0:n])
		}
	}

	file.Close()
}
