package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

type problem struct {
	question string
	answer   string
}

func parseCSVLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for index, line := range lines {
		// The parsed line format is `question,answer`
		// First index (0) is the question, and the latter is the answer
		ret[index] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]), // Trim possible preceeding/ trailing spaces
		}
	}
	return ret
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func initScanner() *bufio.Scanner {
	return bufio.NewScanner(os.Stdin)
}

func main() {
	fmt.Println("Welcome to Quiz Game")

	// Define flag for running this cli apps
	// To access help, type -h or --help in console
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of `question,answer`")

	// Parse all defined flags before being used in the program
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// Parse the CSV file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the CSV file.")
	}

	problems := parseCSVLines(lines)
	fmt.Println(problems)

	// Start the quiz
	scanner := initScanner()
	correctAnswers := 0
	fmt.Println("Answer as many questions as you can.")
	for index, problem := range problems {
		// Display the problem
		fmt.Printf("%v. %s = ", index+1, problem.question)

		// Retrieve user input from console
		scanner.Scan()
		userAnswer := scanner.Text()

		// Check for answer's correctness
		if userAnswer == problem.answer {
			correctAnswers++
		}
	}

	fmt.Println()
	fmt.Printf("You scored %v out of %v problems.\n", correctAnswers, len(problems))
}
