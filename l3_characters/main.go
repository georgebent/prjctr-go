package main

import "fmt"

func main() {
	char1, char2 := 'a', 'c'
	text := "abcd abc cd"
	count1, count2 := countTwoCharacters(text, char1, char2)
	fmt.Printf("character `%c` occurred %d times, character `%c` occurred %d times\n", char1, count1, char2, count2)
}
func countTwoCharacters(text string, char1, char2 rune) (int, int) {
	char1Count := 0
	char2Count := 0

	for _, char := range text {
		switch char {
		case char1:
			char1Count++
		case char2:
			char2Count++
		}
	}

	return char1Count, char2Count
}
