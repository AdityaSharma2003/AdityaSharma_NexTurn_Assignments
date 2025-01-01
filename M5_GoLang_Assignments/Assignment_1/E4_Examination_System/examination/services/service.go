package services

import (
	"bufio"
	"examination/examination/models"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var Questions []models.Question
var Score int = 0
var TotalScore int = 0

func InitializeQuestions() {
	question1 := models.Question{
		ID:       2,
		Question: "What is the capital of India?",
		Options: map[int]string{
			1: "Mumbai",
			2: "Delhi",
			3: "Kolkata",
			4: "Chennai",
		},
		Answer:    2,
		Weightage: 5,
	}
	question2 := models.Question{
		ID:       1,
		Question: "What kind of language is Python?",
		Options: map[int]string{
			1: "Interpreted",
			2: "Compiled",
			3: "Both",
			4: "None of the above",
		},
		Answer:    1,
		Weightage: 5,
	}
	question3 := models.Question{
		ID:       3,
		Question: "Which data structure uses LIFO?",
		Options: map[int]string{
			1: "Queue",
			2: "Stack",
			3: "Array",
			4: "Linked List",
		},
		Answer:    2,
		Weightage: 5,
	}
	question4 := models.Question{
		ID:       4,
		Question: "Which planet is known as the Red Planet?",
		Options: map[int]string{
			1: "Mars",
			2: "Jupiter",
			3: "Venus",
			4: "Mercury",
		},
		Answer:    1,
		Weightage: 5,
	}
	question5 := models.Question{
		ID:       5,
		Question: "Who developed the theory of relativity?",
		Options: map[int]string{
			1: "Isaac Newton",
			2: "Albert Einstein",
			3: "Nikola Tesla",
			4: "Marie Curie",
		},
		Answer:    2,
		Weightage: 5,
	}
	question6 := models.Question{
		ID:       6,
		Question: "Which is the largest mammal in the world?",
		Options: map[int]string{
			1: "Elephant",
			2: "Blue Whale",
			3: "Giraffe",
			4: "Hippopotamus",
		},
		Answer:    2,
		Weightage: 5,
	}
	question7 := models.Question{
		ID:       7,
		Question: "Which programming language is known as the language of the web?",
		Options: map[int]string{
			1: "C++",
			2: "Python",
			3: "JavaScript",
			4: "Java",
		},
		Answer:    3,
		Weightage: 5,
	}

	Questions = append(Questions, question1, question2, question3, question4, question5, question6, question7)

	for _, question := range Questions {
		TotalScore += question.Weightage
	}
}

func TakeQuiz() {
	const questionTimeLimit = 10
	reader := bufio.NewReader(os.Stdin)

	for _, question := range Questions {
		fmt.Println(question.Question)
		for option, value := range question.Options {
			fmt.Printf("%d: %s\n", option, value)
		}

		var choice int
		startTime := time.Now()
		inputCaptured := false

		for time.Since(startTime).Seconds() < questionTimeLimit {
			remainingTime := questionTimeLimit - int(time.Since(startTime).Seconds())
			if remainingTime <= 0 {
				break
			}

			fmt.Printf("Enter your choice (1-4) or 5 to exit (%d seconds left): ", remainingTime)

			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input. Please try again.")
				continue
			}

			input = strings.TrimSpace(input)
			choice, err = strconv.Atoi(input)
			if err == nil && choice >= 1 && choice <= 5 {
				inputCaptured = true
				break
			}

			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
		}

		if !inputCaptured {
			fmt.Println("\nTime's up for this question!")
			continue
		}

		if choice == 5 {
			fmt.Println("Exiting the quiz. Thank you!")
			return
		}

		if choice == question.Answer {
			Score += question.Weightage
		}
	}
}

func ScoreCalculation() {
	if TotalScore == 0 {
		fmt.Println("Total score cannot be zero.")
		return
	}

	percentage := float64(Score) / float64(TotalScore) * 100

	if percentage >= 90 {
		fmt.Println("Performance: Excellent")
	} else if percentage >= 70 {
		fmt.Println("Performance: Good")
	} else if percentage >= 50 {
		fmt.Println("Performance: Average")
	} else {
		fmt.Println("Performance: Needs Improvement")
	}

	fmt.Printf("Your score: %d/%d (%.2f%%)\n", Score, TotalScore, percentage)
}
