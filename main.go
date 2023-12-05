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
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timelimit := flag.Int("limit", 30, "the time limit for quiz in seconds")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
		os.Exit(1)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file."))
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	correct := 0
	for i, p := range problems {
		fmt.Printf("question #%d: %s = ", i+1, p.q)

		answerCh := make(chan string) // channel, for a way to get the answer.
		go func() {                   // anonymous function to get/scan the user answer
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer // get the answer & send it to the channel(made it a closure by using outside var)
		}() // calling it rightaway

		select {
		case <-timer.C: // check if timer expired?
			fmt.Printf("\nYou scored %d out of %d.\n", correct, len(problems))
			return
		case answer := <-answerCh: // if we get the answer then check for correctness
			if answer == p.a {
				correct++
			}
			// Note: no default case, we always either wait for answer or the timer expires!!!
		}

	}
	fmt.Printf("You scored %d out of %d.\n", correct, len(problems))
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
