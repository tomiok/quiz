package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

const (
	defaultFileName = "problems.csv"
)

func main() {
	file, err := openFile(defaultFileName)

	if err != nil {
		panic(err.Error())
	}

	statements := readCSV(file)
	scanner := bufio.NewScanner(os.Stdin)

	done := make(chan bool)
	go func() {
		for i, v := range statements {
			fmt.Printf("#%d: %s = ", i+1, v.q)
			_ = scanner.Text()
			fmt.Println()
		}
		done <- true
	}()

	select {
	case <-done:
	}
}

// open the file, returning the interface
func openFile(fileName string) (io.Reader, error) {
	return os.Open(fileName)
}

func readCSV(r io.Reader) []statement {
	recs, err := csv.NewReader(r).ReadAll()

	if err != nil {
		log.Println("cannot read reader, do'h")
		panic(err.Error())
	}

	total := len(recs)

	log.Printf("total: %d", total)

	var ss []statement
	for _, v := range recs {
		var s statement
		s.q = v[0]
		s.a = v[1]
		ss = append(ss, s)
	}

	return ss
}

type statement struct {
	q string
	a string
}
