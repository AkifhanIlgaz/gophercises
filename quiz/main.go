package main

import "flag"

func main() {
	filePath := flag.String("path", "problems.csv", "File path to a CSV file ")
	timer := *flag.Int("timer", 30, "Time limit")
	flag.Parse()

	quiz := newQuiz(*filePath, timer)

	quiz.startQuiz()

}
