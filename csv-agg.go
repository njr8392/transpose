package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

/*the goal is to read a bunch of tables from a csv file and write them to a single table.
adobe report builder forces us to put data in blocks even if we eventually need to combine
all the data. this will read every block and combine them into one block. */
var WRITER *os.File = os.Stdout

//number of columns within each block
const STEPSIZE = 6

func main() {
	file := readcsv(os.Args[1])
	cleansed := flatten(file)
	splt := split(cleansed, STEPSIZE)
	writer := csv.NewWriter(WRITER)
	err := writer.WriteAll(splt)
	if err != nil {
		fmt.Println(err)
	}

}

//split the line into equal chunks based on the length of each block
//Done after removing whitespace
func split(buf [][]string, step int) [][]string {
	var cleaned [][]string
	var tmp []string
	for _, line := range buf {
		for cell := range line {
			tmp = append(tmp, line[cell])

			if len(tmp) == STEPSIZE {
				cleaned = append(cleaned, tmp)
				tmp = nil
			}
		}
	}
	return cleaned
}

//remove empty cells from the file
func flatten(data [][]string) [][]string {
	var flattened [][]string
	for _, line := range data {
		var clean []string
		for _, item := range line {
			if item != "" {
				clean = append(clean, item)
			}
		}
		flattened = append(flattened, clean)
	}
	return flattened
}

func readcsv(file string) [][]string {
	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()
	reader := csv.NewReader(f)
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

