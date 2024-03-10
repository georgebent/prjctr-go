package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	description := "Вітаємо вас у грі"
	question := "Виберіть один з варіантів:"
	var variants = []string{"continue", "new game", "settings", "back"}
	chosen := ask(description, question, variants)

	fmt.Printf("Ви обрали '%s'\n", variants[chosen])
}

func ask(description string, question string, options []string) int {
	fmt.Println(description)

	if len(options) == 0 {
		os.Exit(0)
	}
	fmt.Println(question)
	for i, variant := range options {
		fmt.Printf("%d. %s\n", i+1, variant)
	}
	fmt.Println("Введіть номер відповідного варіанту або 'exit' для виходу:")

	var input string
	_, err := fmt.Scanf("%v", &input)
	if err != nil {
		fmt.Println("Помилка зчитування вводу:", err)
		ask("Помилка зчитування вводу. Будь ласка, спробуйте ще раз.", question, options)
	}

	if strings.ToLower(input) == "exit" {
		fmt.Println("Вихід з програми...")
		os.Exit(0)
	}

	choice := strings.TrimSpace(input)
	var choiceNum int
	_, err = fmt.Sscanf(choice, "%d", &choiceNum)
	if err != nil || choiceNum < 1 || choiceNum > len(options) {
		ask("Вибір не вірний. Будь ласка, спробуйте ще раз.", question, options)
	}

	choiceNum = choiceNum - 1

	return choiceNum
}
