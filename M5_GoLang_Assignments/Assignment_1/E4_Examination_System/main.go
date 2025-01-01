package main

import exam "examination/examination"

func main() {
	exam.InitializeQuestions()
	exam.TakeQuiz()
	exam.ScoreCalculation()
}
