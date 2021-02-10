package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Flag
	// Define flag to be added when running the program's build (add --h for flag's help)
	// All flag is a pointer to T (pointer to Type value stored in memory)
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")

	// Must be called after defining all flags
	flag.Parse()

	// Open the file with name
	file, err := os.Open(*csvFilename)
	if err != nil {
		// Notes: Sprintf format string based on format specifier and return the resulting string
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	// Create reader to read the csv file -> csv.NewReader takes an io.Reader instance
	// io.Reader instance represent an entity in which we can read a stream of bytes from
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}

	problems := parseLines(lines)
	correctProblemsCount := 0

	// Keep track of the number of problem correctly answered by user
	// Loop through every problems, present it to user, and ask for answer
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.question)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.answer {
			correctProblemsCount++
		}
	}

	fmt.Printf("You correctly answered %d out of %d problems.\n", correctProblemsCount, len(problems))
	stdInputReader := bufio.NewReader(os.Stdin)
	stdInputReader.ReadString('\n')
}

// Function to parse slices of slice of math problem from the read csv file
func parseLines(lines [][]string) []Problem {
	ret := make([]Problem, len(lines))
	for i, v := range lines {
		// Every v is a slice of string, every i is the index of that value
		// Ex. [5+2 7] -> the left side is the question (index 0), the right side is the answer (index 1)
		ret[i] = Problem{
			question: v[0],
			answer:   strings.TrimSpace(v[1]),
		}
	}
	return ret
}

// Problem is a struct which represent a single math question and answer
type Problem struct {
	question string
	answer   string
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
