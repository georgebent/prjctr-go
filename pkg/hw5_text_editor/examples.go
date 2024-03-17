package hw5_text_editor

import (
	"fmt"
	"time"
)

func findWordInMap(text []string, search string) []string {
	mapEditor := MapTextEditor{}
	mapEditor.Init(text)

	return mapEditor.FindWordInRows(search)
}

func findWordInLine(text []string, search string) []string {
	lineEditor := LineTextEditor{}
	lineEditor.Init(text)

	return lineEditor.FindWordInRows(search)
}

func RunSimpleExample() {
	text := []string{
		"Hash map data structures use a hash function, which turns a key into an index within an underlying array.",
		"The Hash function can be used to access an index when hash inserting a value or retrieving a value from a hash map.",
		"Hash map underlying data structure",
		"Hash maps are built on top of an underlying array data structure using an indexing system.",
		"Each index in the array can store one key-value pair.",
		"If the hash map is implemented using chaining for collision resolution, each index can store another data structure such as a linked list, which stores all values for multiple keys that hash to the same index.",
		"Each Hash Map key can be paired with only one value. However, different keys can be paired with the same value.",
		"Try hash.",
		"Try only map.",
	}

	search := "hash"

	for _, row := range findWordInMap(text, search) {
		fmt.Println(row)
	}

	fmt.Println("___comparing___")

	for _, row := range findWordInLine(text, search) {
		fmt.Println(row)
	}
}

func RunFilesTimeExample() {
	textFiles := map[string]string{
		"10":  "pkg/hw5_text_editor/texts/text10.txt",
		"100": "pkg/hw5_text_editor/texts/text100.txt",
		"10K": "pkg/hw5_text_editor/texts/text10K.txt",
	}

	search := "lorem"

	for name, file := range textFiles {
		text := ReadFile(file)

		start := time.Now()
		foundMap := findWordInMap(text, search)
		elapsed := time.Since(start).Seconds()
		fmt.Printf("Індексований Пошук серед %v рядків, тривав: %.8f секунд і знайшов %d рядків\n", name, elapsed, len(foundMap))

		start = time.Now()
		foundLine := findWordInLine(text, search)
		elapsed = time.Since(start).Seconds()
		fmt.Printf("Лінійний Пошук серед %v рядків, тривав: %.8f секунд і знайшов %d рядків\n\n", name, elapsed, len(foundLine))
	}
}
