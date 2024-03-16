package hw4_1_text_editor

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const defaultSearch = "feel"

func Run() {
	search := defaultSearch
	if len(os.Args) > 1 {
		search = os.Args[1]
	}

	myFile, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer myFile.Close()

	scanner := bufio.NewScanner(myFile)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Contains(line, search) {
			output := strings.Replace(line, search, fmt.Sprintf("\u001B[34m%v\u001B[0m", search), -1)
			fmt.Printf("%v\n\n", output)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}
