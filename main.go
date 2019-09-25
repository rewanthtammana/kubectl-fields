package main

import (
	"fmt"
	"os"
	"io"
	"flag"
	"bufio"
	"strings"
	"regexp"
)

const (
    Author      = "Rewanth Cool"    // Package author: first name - last name
    Version     = "0.1"       // Package version: major.minor.maintenance.revision
    ReleaseDate = "2019-25-09"      // Release date: year-month-day
)

// This method uses huffman compression analogy to compress kubectl explain recursive output for further processing
func process(data string) ([]int, []string) {
	var spaces []int
	var fields []string
	counter := 0
	substring := ""
	
	for i:=0; i<len(data); i++ {
		if data[i] == 32 {
			if substring != "" {
				spaces = append(spaces, counter)
				fields = append(fields, substring)
				substring = ""
				counter = 0
			}
			counter++
		} else {
			substring+=string(data[i])
		}
	}

	spaces = append(spaces, counter)
	fields = append(fields, substring)

	return spaces, fields
}

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

func getIndex(data string, substring string) int {
	return strings.Index(data, substring)
}

// This method finds the given regex and deletes it from memory
func findAndDelete(data string, regex string) string {
	var re = regexp.MustCompile(regex)
	return re.ReplaceAllString(data, "")
}

// This method returns the parent id of given child id
func findParentIndex(spaces []int, index int, tabLength int) int {
	for i:=index-1; i>=0; i-- {
		if spaces[i] == spaces[index] - tabLength {
			return i
		}
	}
	return -1
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

	const Seperator = "FIELDS:"
	const TabLength = 3
	
	input := getStdInput()
	start := getIndex(input, Seperator)
	if start == -1 {
		fmt.Println(input)
		return
	}

	// Take only data under fields section
	data := input[start+len(Seperator):]

	// Deletes all tabs, data type information and new lines from input for compressed processing
	const Trash = string(9)+"|"+`<[^>]*>|\n`
	data = findAndDelete(data, Trash)

	var spaces []int
	var fields []string

	// Process data for further analysis
	spaces, fields = process(data)

	var pattern string
	var hierarchy string
	var index int

	for j:=1; j<len(os.Args); j++ {
		pattern = os.Args[j]
		for i:=0; i<len(fields) ;i++ {
			// Single conditional statement to check status of ignore case flag
			// The first condition refers to case sensitive pattern match
			// The second condition refers to case insensitive pattern match
			if (!(*ignoreCase) && strings.Contains(fields[i], pattern)) || (*ignoreCase && strings.Contains(strings.ToLower(fields[i]), strings.ToLower(pattern))) {
				hierarchy = fields[i]
				index = i
				for index != -1 {
					index = findParentIndex(spaces, index, TabLength)
					if index != -1 {
						hierarchy = fields[index] + "." + hierarchy
					}
				}
				fmt.Println(hierarchy)
			} 
		}
	}
}
