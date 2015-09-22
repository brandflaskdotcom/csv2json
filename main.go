package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

type ImageRef struct {
	Fbobjid  string
	Parentid string
	Name     string
	Url      string
}

func main() {
	input_file := flag.String("f", "", "input file path")
	flag.Parse()
	f, _ := os.Open(*input_file)
	defer f.Close()

	r := csv.NewReader(f)

	record, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	var allRecord []ImageRef
	for _, each := range record {
		var tmp ImageRef

		tmp.Fbobjid = each[0]
		tmp.Parentid = each[1] // need to cast integer to string
		tmp.Name = each[2]
		tmp.Url = each[3]
		allRecord = append(allRecord, tmp)
	}
	jsondata, err := json.Marshal(allRecord) // convert to JSON

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(jsondata))
}
