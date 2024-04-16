package Quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func Start() {
	fmt.Println("Starting the Quiz")
	
	csvFileName, timeLimit := defineFlags()
	lines := handleOpenCsvFile(csvFileName)

	questions := formatQuestions(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	corrects := 0

	corrects = startQuiz(questions, timer, corrects)

	fmt.Printf("Your score is %d of %d \n", corrects, len(questions))
}

func defineFlags() (*string, *int64) {
	csvFileName := flag.String("csv", "Packages/quiz-game/Questions/problems.csv", "give a path for a csv file on the format 'Question,Answer'")
	timeLimit := flag.Int64("time", 30, "the time limit to anser the questions in seconds")
	flag.Parse()
	return csvFileName, timeLimit
}

func handleOpenCsvFile(csvFileName *string) [][]string {
	file, err := os.Open(*csvFileName)

	if err != nil {
		exit(fmt.Sprintf("Failed to open CSV file: %s \n", *csvFileName))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()

	if err != nil {
		exit("Failed to parse the CSV")
	}
	return lines
}

func formatQuestions(lines [][]string) []question {
	questionsSlice := make([]question, len(lines))
	for index, line := range lines {
		questionsSlice[index] = question{
			task:   line[0],
			answer: line[1],
		}
	}

	return questionsSlice
}

func startQuiz(questions []question, timer *time.Timer, corrects int) int {
quizLoop:
	for index, question := range questions {
		fmt.Printf("Question %d: %s is? \n", index+1, question.task)

		userAnswerChannel := make(chan string)

		go func() {
			var userAnswer string
			fmt.Scanf("%s\n", &userAnswer)
			userAnswerChannel <- userAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break quizLoop
		case userAnswer := <-userAnswerChannel:
			if userAnswer == question.answer {
				corrects++
			}
		}
	}

	return corrects
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

type question struct {
	task   string
	answer string
}
