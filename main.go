package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type Problem struct {
	question string
	answer string
}

func main() {
	// Generate flag with -h for help
	filename := flag.String("csv", "file.csv", "CSV file in formation 'question,answer'")
	flag.Parse()

	// Read the file
	file, err := os.Open(*filename)
	if err != nil {
		exit(fmt.Sprintf("Could not open the file : %s\n", *filename))
	}

	// Create the csv reader
	reader := csv.NewReader(file)

	// Read all lines
	lines, err := reader.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the file"))
	}

	problems := parseLines(lines)
	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.question)
		var answer string
		// make sur it has a pointer value to it
		_, _ = fmt.Scanf("%s\n", &answer)
		if answer == problem.answer {
			correct++
		}
	}

	fmt.Printf("Scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []Problem {
	r := make([]Problem, len(lines))

	for i, line := range lines {
		r[i] = Problem{
			question: line[0],
			answer: line[1],
		}
	}

	return r
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}