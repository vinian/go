package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	lineReader := bufio.NewReaderSize(file, 100)
	for line, _, e := lineReader.ReadLine(); e == nil; line, _, e = lineReader.ReadLine() {
		os.Stdout.Write(line)
		fmt.Printf("\n")
	}

	file.Close()
}
