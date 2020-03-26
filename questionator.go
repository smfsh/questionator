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
	// Prepare our input flag variables.
	var (
		fileIn       string
		fileOut      string
		tabDelimiter bool
	)

	// Define some strings to be used for the help text.
	const (
		inUsage        = "Path and name of CSV for parsing (default: questions.csv)"
		outUsage       = "Path and name to output JSON file (default: gamedata.json)"
		delimiterUsage = "Boolean value for whether tabs should be used as the delimiter (commas are default)"
	)

	// Prepare to capture whichever flags were used on command execution.
	// We setup multiple flags for the same variables here to catch both
	// the longhand and the shorthand versions of flags.
	flag.StringVar(&fileIn, "input", "questions.csv", inUsage)
	flag.StringVar(&fileIn, "in", "questions.csv", inUsage)
	flag.StringVar(&fileIn, "i", "questions.csv", inUsage)

	flag.StringVar(&fileOut, "output", "gamedata.json", outUsage)
	flag.StringVar(&fileOut, "out", "gamedata.json", outUsage)
	flag.StringVar(&fileOut, "o", "gamedata.json", outUsage)

	flag.BoolVar(&tabDelimiter, "tabs", false, delimiterUsage)
	flag.BoolVar(&tabDelimiter, "tab", false, delimiterUsage)
	flag.BoolVar(&tabDelimiter, "t", false, delimiterUsage)

	// Set the previously defined variables to the actual passed value.
	flag.Parse()

	// Read in the CSV file.
	in, err := ioutil.ReadFile(fileIn)
	if err != nil {
		log.Fatal(err)
	}

	// Initiate a reader that allows us to parse the CSV byte by byte.
	r := csv.NewReader(bytes.NewReader(in))
	// Let the reader know our csv might have variable quantities of fields per line.
	r.FieldsPerRecord = -1
	// Check whether the file being read is tab-delimited.
	if tabDelimiter {
		r.Comma = 0x0009
	}

	// Read each line and value into an array of an array of strings.
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	// Instantiate our gamedata struct that will be returned at the end containing
	// all of the data we parsed.
	gd := gamedata{}

	// Loop through the records array, touching each record (or row in the
	// CSV as it were) individually.
	for _, record := range records {
		// Instantiate a new question type to represent this row.
		q := question{
			// Setting value to the first item in the array (0), the question text.
			Question: record[0],
		}

		// Instantiating a brand new answer that we will nest in the previously
		// declared question variable, q.
		a := answer{}

		// Loop through each of the answers. Start one value into the array,
		// notated by "1:" because that gives a new "array", or slice in
		// Golang nomenclature, that contains only answers and values.
		for i, v := range record[1:] {
			// If the CSV contains empty records, skip over them.
			if v == "" {
				continue
			}

			// Check if the answer item we're looking at is even or odd.
			// Even numbers mean the value is the answer text.
			// Odd numbers mean the value is the answer value.
			if i%2 == 0 {
				a.Answer = v
			} else {
				a.Value = v
				// If a value has been set, we know the full answer is complete
				// with both parts, the text and the value, so we append it to
				// the question variable, q.
				q.Answers = append(q.Answers, a)
			}
		}
		// As each question completes its loop, and each of the answers for that
		// question are added to the question, we add the completed question to
		// the gamedata array, gd.
		gd.Questions = append(gd.Questions, q)
	}

	// Take our gd struct, marshal it into json, and finally, format it
	// with some double space indentations.
	j, err := json.MarshalIndent(gd, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	// Print out the formatted and completed json to the console.
	fmt.Println(string(j))

	// Write the formatted and completed json to the file specified.
	err = ioutil.WriteFile(fileOut, j, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
