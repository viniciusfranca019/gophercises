package main

import (
	"bufio"
	"fmt"
	Quiz "github.com/viniciusfranca019/gophercises/Packages/quiz-game"
	"os"
	"strconv"
	"strings"
)

const (
	QUIZ = "quiz"
)

func main() {
	fmt.Println("Choose an APP to start:")
	apps := []string{QUIZ}

	selectedAppIndex := selectApp(apps)
	selectedApp := handleAppIndex(selectedAppIndex, apps)

	fmt.Printf("You chose %s\n", selectedApp)

	runApp(selectedApp)
}

func selectApp(apps []string) int {
	for index, option := range apps {
		fmt.Printf("%d. %s\n", index+1, option)
	}

	bufioReader := bufio.NewReader(os.Stdin)

	var selectedAppIndex int

	for {
		fmt.Print("Enter the choosen app (number): ")
		input, err := bufioReader.ReadString('\n')
		input = strings.ReplaceAll(input, "\n", "")

		if input == "" {
			input = "0"
		}

		selectedAppIndex, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Failed to read input:", err)
			continue
		}

		break
	}
	return selectedAppIndex
}

func handleAppIndex(selectedAppIndex int, apps []string) string {
	selectedApp := "none"

	if selectedAppIndex != 0 {
		selectedApp = apps[selectedAppIndex-1]
	}
	return selectedApp
}

func runApp(selectedApp string) {
	switch selectedApp {
	case QUIZ:
		Quiz.Start()
		break
	default:
		fmt.Println("hello world")
	}
}
