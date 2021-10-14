package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
	}
	//rader is a common interface used in go.can read in docs
	r := csv.NewReader(file)
	//read all lines in csv
	lines, err := r.ReadAll()
	if err != nil{
		exit("Failed to parse provided CSV file.")
	}


	problems := parseLines(lines)

	correct := 0
	for i, p := range problems{
		fmt.Printf("problem #%d: %s = \n", i+1, p.q)
		var answer string
		//good for this simple math quiz
		//& is a pointer value
		fmt.Scanf("%s\n", &answer)
		if answer == p.a{
			correct++
		}
		
}
fmt.Printf("you scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines{
		ret[i] =problem{
			q: line[0],
			//this trims the white space from the answers so if there are spaces it won't error them
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}


//struct for the problem, follow this even if not csv file
type problem struct {
	q string
	a string
}

func exit(msg string){
	fmt.Println(msg)
	os.Exit(1)
}
