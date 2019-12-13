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
)

func main() {
	const (
		fileDefault  = "problems.csv"
		fileUsage    = "path of csv file with the problems"
		limitDefault = 30
		limitUsage   = "number of second to do the quiz"
	)

	filePath := flag.String("file", fileDefault, fileUsage)
	limit := flag.Int("limit", limitDefault, limitUsage)

	flag.Parse()

	csv_file, err := os.Open(*filePath)

	if err != nil {
		fmt.Println("Could not read file")
	}
	reader := bufio.NewReader(os.Stdin)
	fr := csv.NewReader(csv_file)
	counter := 0
	score := 0

	fmt.Println("Time limit: ", *limit)
	for {
		record, err := fr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		counter++
		fmt.Print("Problem #", counter, ": ")
		fmt.Print(record[0], " = ")
		result, _ := reader.ReadString('\n')
		result = strings.TrimRight(result, "\n")
		if result == record[1] {
			score++
		}
	}
	fmt.Println("You scored ", score, "out of ", counter)
}
