package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

type gamedata struct {
	Questions []question
}

type question struct {
	Question string
	Answers  []answer
}

type answer struct {
	Answer string
	Value  string
}

func main() {
	var (
		fileIn       string
		fileOut      string
		tabDelimiter bool
	)

	const (
		inUsage        = "Path and name of CSV for parsing (default: questions.csv)"
		outUsage       = "Path and name to output JSON file (default: gamedata.json)"
		delimiterUsage = "Boolean value for whether tabs should be used as the delimiter (commas are default)"
	)

	flag.StringVar(&fileIn, "input", "questions.csv", inUsage)
	flag.StringVar(&fileIn, "in", "questions.csv", inUsage)
	flag.StringVar(&fileIn, "i", "questions.csv", inUsage)

	flag.StringVar(&fileOut, "output", "gamedata.json", outUsage)
	flag.StringVar(&fileOut, "out", "gamedata.json", outUsage)
	flag.StringVar(&fileOut, "o", "gamedata.json", outUsage)

	flag.BoolVar(&tabDelimiter, "tabs", false, delimiterUsage)
	flag.BoolVar(&tabDelimiter, "tab", false, delimiterUsage)
	flag.BoolVar(&tabDelimiter, "t", false, delimiterUsage)

	flag.Parse()

	in, err := ioutil.ReadFile(fileIn)

	r := csv.NewReader(bytes.NewReader(in))
	r.FieldsPerRecord = -1
	if tabDelimiter {
		r.Comma = 0x0009
	}

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	gd := gamedata{}

	for _, record := range records {
		q := question{
			Question: record[0],
		}

		a := answer{}

		for i, v := range record[1:] {
			if v == "" {
				continue
			}

			if i%2 == 0 {
				a.Answer = v
			} else {
				a.Value = v
				q.Answers = append(q.Answers, a)
			}
		}
		gd.Questions = append(gd.Questions, q)
	}

	j, err := json.MarshalIndent(gd, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(j))
	err = ioutil.WriteFile(fileOut, j, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
