package hw5_text_editor

import (
	"strings"
)

type LineTextEditor struct {
	text []string
}

func (editor *LineTextEditor) Init(text []string) {
	editor.text = text
}

func (editor *LineTextEditor) Clear() {
	editor.text = []string{}
}

func (editor *LineTextEditor) FindWordInRows(search string) []string {
	var rows []string

	for _, row := range editor.text {
		if strings.Contains(strings.ToLower(row), strings.ToLower(search)) {
			rows = append(rows, row)
		}
	}

	return rows
}
