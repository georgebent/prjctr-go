package hw5_text_editor

import (
	"regexp"
	"strings"
)

type MapTextEditor struct {
	text []string
}

func (editor *MapTextEditor) Init(text []string) {
	editor.text = text
}

func (editor *MapTextEditor) Clear() {
	editor.text = []string{}
}

func (editor *MapTextEditor) AddRow(row string) {
	editor.text = append(editor.text, row)
}

func (editor *MapTextEditor) FindWordInRows(search string) []string {
	mappedWords := editor.IndexWordRows()
	var rows []string

	for _, rowKey := range mappedWords[strings.ToLower(search)] {
		rows = append(rows, editor.text[rowKey])
	}

	return rows
}

func (editor *MapTextEditor) IndexWordRows() map[string][]int {
	stringMap := make(map[string][]int)

	for rowKey, row := range editor.text {
		re := regexp.MustCompile(`[[:alnum:]]+`)
		words := re.FindAllString(row, -1)

		for _, word := range words {
			if isExist(stringMap[strings.ToLower(word)], rowKey) {
				continue
			}

			stringMap[strings.ToLower(word)] = append(stringMap[strings.ToLower(word)], rowKey)
		}
	}

	return stringMap
}

func isExist(numbers []int, search int) bool {
	for _, number := range numbers {
		if number == search {
			return true
		}
	}

	return false
}
