package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in format of 'question,answer")

	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file: %s\n", *csvFilename))
	}
	//reader is a common interface used in go.can read in docs
	r := csv.NewReader(file)
	//read all lines in csv
	lines, err := r.ReadAll()
	if err != nil{
		exit("Failed to parse provided CSV file.")
	}


	problems := parseLines(lines)

	//do timer after set up as could lose time
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//label, can use with select to break out of loop
	problemloop:
	
	for i, p := range problems{
		correct := 0
		//displays problem
		fmt.Printf("problem #%d: %s = \n", i+1, p.q)

		//answer channel
		answerCh := make(chan string)

		//go routine anonymous function
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			//the below code turns this annonymous function into a clsoure, closure uses data defined outside it
			// <- = sends answer into channel, arrow points to way data is moving
			answerCh <- answer
		}() //call the above function

		select {
			//waits for message from timer channel
		case <-timer.C:
			fmt.Printf("you scored %d out of %d.\n", correct, len(problems))
			break problemloop

			// will now time out
		case answer := <-answerCh:

		if answer == p.a{
			correct++
		}
		}	
}

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
