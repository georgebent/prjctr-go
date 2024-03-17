package hw5_text_editor

import (
	"bufio"
	"fmt"
	"os"
)

type TextEditor interface {
	Init(text []string)
	FindWordInRows(search string) []string
	Clear()
}

func ReadFile(file string) []string {
	myFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}
	}
	defer myFile.Close()

	var text []string
	scanner := bufio.NewScanner(myFile)
	for scanner.Scan() {
		line := scanner.Text()

		text = append(text, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return text
}
