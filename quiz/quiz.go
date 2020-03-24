package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

type options struct {
	file      string
	timeLimit int
}

type question struct {
	q, a string
}

func (o options) String() string {
	return fmt.Sprintf("-csv %s -timeLimit %d\n", o.file, o.timeLimit)
}

func parseCmdLine() options {
	opt := options{}

	flag.StringVar(&opt.file, "csv", "problems.csv", "CSV file containing quiz questions.")
	flag.IntVar(&opt.timeLimit, "timeLimit", 600, "Time limit for quiz in seconds. Default to 5min.")
	flag.Parse()
	return opt
}

func readQuestions(questionFile string) []question {
	file, err := os.Open(questionFile)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	var questions []question
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error parsing questions file", err)
		}
		if cap(record) != 2 {
			log.Fatal("Question is in incorrect format.")
		}
		questions = append(questions, question{
			q: record[0],
			a: record[1],
		})
	}
	return questions

}

func main() {
	//Setup cmd flags
	opt := parseCmdLine()

	//Parse CSV file to create question set
	questions := readQuestions(opt.file)

	answersCorrect := 0
	answers := bufio.NewReader(os.Stdin)

	//Prompt to start timer
	fmt.Printf("%d second timer will start once you press [Enter]...", opt.timeLimit)
	answers.ReadString('\n')
	timer := time.NewTimer(time.Duration(opt.timeLimit) * time.Second)
	defer timer.Stop()

	for i, q := range questions {
		select {
		case <-timer.C:
			fmt.Printf("Time has expired! You answer %d/%d correctly.", answersCorrect, len(questions))
			os.Exit(1)
		default:
			fmt.Printf("%d) %s: ", i+1, q.q)
			a, _ := answers.ReadString('\n')
			a = strings.TrimSpace(a)
			if q.a == a {
				answersCorrect++
			}
		}
	}

	fmt.Printf("You answered %d/%d correctly!\n", answersCorrect, len(questions))
}
