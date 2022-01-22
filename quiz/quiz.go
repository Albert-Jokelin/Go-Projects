package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to the quiz app!")

	// Open the file
	file := flag.String("csv", "D:/Code/GitHub/Go-Projects/quiz/problems.csv", "A csv file in the format of 'question, answer'. ")
	timeLimit := flag.Int("limit", 10, "the time limit for the quiz")
	flag.Parse()

	csvfile, err := os.Open(*file)
	if err != nil {
		log.Fatalln("ERROR: Could not open CSV File.", err)
	}

	// Parse the file

	r := csv.NewReader(csvfile)

	// Interating through the records
	var corr, total int = 0, 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("ERROR Reading record. ", err)
			break
		}

		total++
		fmt.Println("What is", record[0])

		ansCh := make(chan string)
		go func() {
			var ans string
			fmt.Scanln(&ans)
			ansCh <- ans
		}()

		select {
		case <-timer.C:
			println("\nOh no, you ran out of time.")
			println("You got ", corr, " out of ", total, " correct. \n")
			return
		case answer := <-ansCh:
			if answer == record[1] {
				corr++
			}

		}

	}

	println("You got ", corr, " out of ", total, " correct. \n")
}
