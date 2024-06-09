package questions

import (
	"encoding/json"
	"fmt"
	"os"
)

type Question struct {
	Country string `json:"country"`
	Capital string `json:"capital"`
}

func LoadQuestions() ([]Question, error) {
	jsonData, err := os.ReadFile("quiz.json")

	if err != nil {
		return nil, fmt.Errorf("failed to load questions: %w", err)
	}
	var questions []Question
	err = json.Unmarshal(jsonData, &questions)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal questions: %w", err)
	}

	return questions, nil
	

}
