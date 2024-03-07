package hw1_biography

import "fmt"

type bornPlace string

func (bp bornPlace) String() string {
	return fmt.Sprintf("Was born in %s.", string(bp))
}

func main() {
	name := "Albert Einstein"
	birthDay := 14
	birthMonth := "March"
	birthYear := 1879
	location := "Ulm, Kingdom of Württemberg, German Empire"
	facts := []string{
		"He was one of the greatest and most influential scientists of all time",
		"Best known for developing the theory of relativity",
		"He received the 1921 Nobel Prize in Physics",
	}

	bp := bornPlace(location)
	biography := makeBiography(name, birthDay, birthMonth, birthYear, bp, facts)

	fmt.Println(biography)
}

func makeBiography(name string, birthDay int, birthMonth string, birthYear int, location bornPlace, facts []string) string {
	factsStr := ""
	for _, fact := range facts {
		factsStr += fmt.Sprintf("%s.\n", fact)
	}

	return fmt.Sprintf("%s (%d %s %d)\n%s\n%s", name, birthDay, birthMonth, birthYear, location, factsStr)
}
