package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

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
	r := csv.NewReader(file)
	_ = r
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}