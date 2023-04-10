package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

/*the goal is to read a bunch of tables from a csv file and write them to a single table.
adobe report builder forces us to put data in blocks even if we eventually need to combine
all the data. this will read each all blocks and combine them into one block. */
var WRITER *os.File = os.Stdout

//number of columns within each block
const STEPSIZE = 4

func main() {
	file := readcsv(os.Args[1])
	cleansed := flatten(file)
	fmt.Println(cleansed)
	splt := split(cleansed, STEPSIZE)
	writer := csv.NewWriter(WRITER)
	err := writer.WriteAll(splt)
	if err != nil {
		fmt.Println(err)
	}

}

//split the line into equal chunks. Done after removing whitespace
func split(buf [][]string, step int) [][]string {
	var cleaned [][]string
	var tmp []string
	for _, line := range buf {
		for i := 0; i < len(line); i++ {
			tmp = append(tmp, line[i])

			if len(tmp) == STEPSIZE {
				cleaned = append(cleaned, tmp)
				tmp = nil
			}
		}
	}
	return cleaned
}

func flatten(data [][]string) [][]string {
	for i, line := range data {
		var clean []string
		for _, item := range line {
			if item != "" {
				clean = append(clean, item)
			}
		}
		data[i] = clean
	}
	return data
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

