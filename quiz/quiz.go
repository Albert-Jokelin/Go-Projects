package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Welcome to the quiz app!")

	// Open the file
	csvfile, err := os.Open("D:/Code/GitHub/Go-Projects/quiz/problems.csv")

	if err != nil {
		log.Fatalln("ERROR: Could not open CSV File.", err)
	}

	// Parse the file

	r := csv.NewReader(csvfile)

	// Interating through the records
	var corr, total int = 0, 0
	ans := ""
	timer1 := time.NewTimer(2 * time.Second)

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

		fmt.Println("What is ", record[0])
		fmt.Scanln(&ans)

		if ans == record[1] {
			corr++
		}
	}

	println("You got ", corr, " out of ", total, " correct. \n")
}
