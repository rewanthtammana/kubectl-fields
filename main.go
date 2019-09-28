package main

import (
	"fmt"
	"os"
	"io"
	"flag"
	"bufio"

	"github.com/rewanth1997/order/pkg/order"
)

const (
    Author      = "Rewanth Cool"    // Package author: first name - last name
    Version     = "0.1-alpha"       // Package version: major.minor.maintenance.revision
    ReleaseDate = "2019-25-09"      // Release date: year-month-day
)

// This method reads std input from kubectl explain recursive module
func getStdInput() string {
	reader := bufio.NewReader(os.Stdin)
	var output []rune

	for {
			input, _, err := reader.ReadRune()
			if err != nil && err == io.EOF {
					break
			}
			output = append(output, input)
	}
	return string(output)
}

func main() {
	ignoreCase := flag.Bool("i", false, "Ignore case and match")
	help := flag.Bool("h", false, "Help menu")
	flag.Parse()

	if *help {
		fmt.Println(`Usage of order:
	kubectl explain --recursive RESOURCE | order PATTERN
	kubectl explain --recursive RESOURCE | order -i PATTERN

A tool to parse kubectl explain recursive module output and format it in a hierarchical order.

This tool is very handy while attempting CKA/CKAD certification.

Options:
	-h Help module
	-i Ignore case distinctions
		`)
		return
	}
	
	input := getStdInput()
	
	order.Parse(input, *ignoreCase)
}
