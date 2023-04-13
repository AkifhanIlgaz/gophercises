package main

import (
	"bufio"
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Quiz struct {
	problems           []problem
	userCorrectAnswers int
	timer              time.Duration
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func newQuiz(filePath string, timer int) Quiz {
	csvFile, err := os.Open(filePath)
	if err != nil {
		exit(fmt.Sprintf("Unable to open %s", filePath))
	}

	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	lines, err := csvReader.ReadAll()
	if err != nil {
		exit("Unable to read CSV file")
	}
	problems := make([]problem, len(lines))

	for i, line := range lines {
		newProblem := newproblem(line[0], strings.TrimSpace(line[1]))
		problems[i] = newProblem
	}

	return Quiz{problems: problems, timer: time.Duration(timer * int(time.Second))}
}

func (q *Quiz) startQuiz() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(color.HiBlueString("Press enter to start the quiz"))
	scanner.Scan()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(q.timer))
	defer cancel()

	resultChan := make(chan int)
	finished := make(chan struct{})
	go askProblems(ctx, q.problems, resultChan, finished, scanner)

out:
	for {
		select {
		case <-ctx.Done():
			color.Red("\nTime is up!")
			break out
		case result := <-resultChan:
			q.userCorrectAnswers += result
		case <-finished:
			break out
		}
	}

	q.printResult()

}

func askProblems(ctx context.Context, problems []problem, resultChan chan int, finished chan<- struct{}, scanner *bufio.Scanner) {
	for i, problem := range problems {
		problem.printProblem(i + 1)

		scanner.Scan()
		userAnswer := strings.TrimSpace(scanner.Text())

		_, err := strconv.Atoi(userAnswer)
		if err != nil {
			color.Red("Error: Answer is not a number")
			continue
		}

		resultChan <- problem.checkAnswer(userAnswer)
	}

	finished <- struct{}{}
}

func (q Quiz) printResult() {
	fmt.Printf("Score: %s/%s\n", color.GreenString("%d", q.userCorrectAnswers), color.BlueString("%d", len(q.problems)))
}

type problem struct {
	question string
	answer   string
}

func newproblem(question, answer string) problem {
	return problem{question, answer}
}

func (p problem) printProblem(index int) {
	equalSignColor := color.Set(color.FgMagenta, color.Bold)
	indexColor := color.Set(color.FgWhite, color.Bold)
	questionColor := color.Set(color.FgYellow)

	indexColor.Printf("#%d ", index)
	questionColor.Printf("%s", p.question)
	equalSignColor.Print(" = ")
}

func (p problem) checkAnswer(userAnswer string) int {
	if p.answer == userAnswer {
		return 1
	}
	return 0
}
