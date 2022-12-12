package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: boucles-v2 input-file.xlsx")
		os.Exit(-1)
	}

	filename := os.Args[1]

	if !strings.HasSuffix(filename, ".xlsx") {
		fmt.Println("Input file should be a valid excel document")
		os.Exit(-1)
	}

	input, err := excelize.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := input.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	output, err := os.Create(strings.Replace(filename, "xlsx", "csv", 1))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the output file.
		if err := output.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	w := bufio.NewWriter(output)
	defer w.Flush()

	// Get all the rows in the Sheet1.
	rows, err := input.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	in := ""
	out := ""
	i := 2
	w.WriteString("Boucle,Timecode In,Timecode Out\n")
	w.WriteString("1,,\n")
	for _, row := range rows {
		firstCell := row[0]
		if isValidRow(firstCell) {
			out = row[1]
			line := fmt.Sprintf("%d,%s,%s\n", i, in, out)
			w.WriteString(line)
			i++
			in = out
		}
	}
}

func isValidRow(cell string) bool {
	match, _ := regexp.MatchString("[0-9]+", cell)
	return match
}
