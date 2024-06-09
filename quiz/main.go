package main

import (
	"fmt"
	"os"
	"quiz/game"
	"quiz/questions"
	"quiz/shuffler"
)

func main() {
	questions, err := questions.LoadQuestions()
	if err != nil {
		fmt.Fprintf(os.Stderr, "An error occurred: %s\n", err)
		os.Exit(1)
	}
	shuffler.Shuffle(questions)
	correctAnswers := game.Run(questions)
	fmt.Printf("You got %d out of %d correct!\n", correctAnswers, len(questions))
}
