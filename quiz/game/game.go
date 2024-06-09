package game

import (
	"bufio"
	"fmt"
	"os"
	"quiz/questions"
	"strings"
)

func Run(questions []questions.Question) (correctAnswers uint) {
	fmt.Println("Welcome to the quiz!")
	for _, question := range questions {
		if askQuestion(question) {
			correctAnswers++
		}
	}
	return correctAnswers
}

func askQuestion(question questions.Question) bool {
	fmt.Printf("\nEnter the capital of %s: ", question.Country)

	if getUserInput() == strings.ToLower(question.Capital) {
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Printf("Wrong! The capital of %s is %s\n", question.Country, question.Capital)
		return false
	}

}

func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter your answer: ")
		result, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred", err)
			continue
		}
		return strings.ToLower(strings.TrimRight(result, "\n"))
	}
}
