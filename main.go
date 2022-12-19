package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: boucles-v2 [directory]")
		os.Exit(-1)
	}

	files, err := os.ReadDir(os.Args[1])
	if err != nil {
		fmt.Printf("Could not read directory %s\n", os.Args[1])
		os.Exit(-1)
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".xlsx") {
			processFile(file.Name())
		}
	}
}

func processFile(filename string) {
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

	fmt.Printf("Processing file %s\n", input.Path)
	outputPath := strings.Replace(filename, "xlsx", "csv", 1)
	output, err := os.Create(outputPath)
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
	fmt.Printf("Writing output to %s\n", outputPath)

	// Get all the rows in the Sheet1.
	rows, err := input.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	in := ""
	out := ""
	i := 1
	for _, row := range rows {
		firstCell := row[0]
		if isValidRow(firstCell) {
			out = row[1]
			if in != "" && out != "" {
				writeLine(w, i, in, out)
				i++
			}
			in = out
		}
	}
	writeLine(w, i, in, plusOneMinute(in))

}

func isValidRow(cell string) bool {
	match, _ := regexp.MatchString("[0-9]+", cell)
	return match
}

func writeLine(w *bufio.Writer, i int, in, out string) {
	line := fmt.Sprintf("%d,%s,%s\n", i, in, out)
	w.WriteString(line)
}

func plusOneMinute(timecode string) string {
	split := strings.Split(timecode, ":")
	if len(split) != 4 {
		return timecode
	}
	minutes, _ := strconv.Atoi(split[1])
	if minutes == 59 {
		hours, _ := strconv.Atoi(split[0])
		return fmt.Sprintf("%02d:00:%s:%s", hours+1, split[2], split[3])
	}
	return fmt.Sprintf("%s:%02d:%s:%s", split[0], minutes+1, split[2], split[3])
}
