package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	text := make([]map[string]string, 100, 1000)
	lineReader := bufio.NewReaderSize(file, 100)
	for line, _, e := lineReader.ReadLine(); e == nil; line, _, e = lineReader.ReadLine() {
		//		os.Stdout.Write(line)
		hash := splitLine(string(line))
		text = append(text, hash)
	}

	fmt.Print(text)
	fmt.Print("\n")
	file.Close()
}

func splitLine(line string) map[string]string {
	data := make(map[string]string)
	str := strings.SplitN(line, " ", 3)

	data["startTime"] = str[0]
	data["endTime"] = str[1]
	data["text"] = str[2]

	return data
}
