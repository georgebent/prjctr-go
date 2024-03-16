package l5_map

import (
	"fmt"
	"strings"
)

func FindWordRows(text []string) map[string][]int {
	stringMap := make(map[string][]int)

	for rowKey, row := range text {
		split := strings.Split(row, " ")
		for _, word := range split {
			stringMap[word] = append(stringMap[word], rowKey)
		}
	}

	return stringMap
}

func example() {
	text := []string{
		"Hash map data structures use a hash function, which turns a key into an index within an underlying array.",
		"The hash function can be used to access an index when inserting a value or retrieving a value from a hash map.",
		"Hash map underlying data structure",
		"Hash maps are built on top of an underlying array data structure using an indexing system.",
		"Each index in the array can store one key-value pair.",
		"If the hash map is implemented using chaining for collision resolution, each index can store another data structure such as a linked list, which stores all values for multiple keys that hash to the same index.",
		"Each Hash Map key can be paired with only one value. However, different keys can be paired with the same value.",
	}

	stringMap := FindWordRows(text)

	for key, value := range stringMap {
		fmt.Printf("\n%s = %q\n", key, value)
		for _, found := range value {
			fmt.Printf("%s = %q\n", key, text[found])
		}

	}
}
