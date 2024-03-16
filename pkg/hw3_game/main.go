package hw3_game

import (
	"fmt"
	"os"
	"strings"
)

func Run() {
	scene, player := initialize()

	for {
		chosen := scene.run(&player)

		scene = scene.getChosenScene(chosen, &player)
		fmt.Printf("Ви обрали '%s'\n", scene.name)
	}
}

func ask(description string, question string, options []*Scene) int {
	fmt.Printf("\n%s\n", description)

	if len(options) == 0 {
		os.Exit(0)
	}
	fmt.Println(question)
	for i, variant := range options {
		fmt.Printf("\x1b[34m%d. %s\n\x1b[0m", i+1, variant.name)
	}
	fmt.Println("\x1b[90m\nВведіть номер відповідного варіанту або 'exit' для виходу:\u001B[0m")

	var input string
	_, err := fmt.Scanf("%v", &input)
	if err != nil {
		fmt.Println("Помилка зчитування вводу:", err)
		return ask("Помилка зчитування вводу. Будь ласка, спробуйте ще раз.", question, options)
	}

	if strings.ToLower(input) == "exit" {
		fmt.Println("Вихід з програми...")
		os.Exit(0)
	}

	choice := strings.TrimSpace(input)
	var choiceNum int
	_, err = fmt.Sscanf(choice, "%d", &choiceNum)
	if err != nil || choiceNum < 1 || choiceNum > len(options) {
		return ask("Вибір не вірний. Будь ласка, спробуйте ще раз.", question, options)
	}

	choiceNum = choiceNum - 1

	return choiceNum
}
