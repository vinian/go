package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: os.Args[0] file[test.txt]")
		// read from stdin and output to stdout
        os.Exit(1)
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	text := make([]map[string]string, 0, 100)
	lineReader := bufio.NewReaderSize(file, 100)
	for line, _, e := lineReader.ReadLine(); e == nil; line, _, e = lineReader.ReadLine() {
		//		os.Stdout.Write(line)
		if e == nil {
			hash := splitLine(string(line))
			// append hash to the array
			text = append(text, hash)
		}
	}

	fmt.Print(text)
	fmt.Print("\n")
	file.Close()
}

func splitLine(line string) map[string]string {
	data := make(map[string]string)

	// split line in the three element
	str := strings.SplitN(line, " ", 3)

	data["startTime"] = str[0]
	data["endTime"] = str[1]
	data["text"] = str[2]

	// return hash
	return data
}
