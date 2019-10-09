/*
Copyright © 2019 Rewanth Cool

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package fields is used to extract the parental fields for given kubectl resource
package fields

import (
	"fmt"
	"github.com/rewanth1997/kubectl-fields/pkg/color"
	"regexp"
	"strings"
)

// This method uses huffman compression analogy to compress kubectl explain recursive output for further processing
func analyze(data string) ([]int, []string) {
	var spaces []int
	var fields []string
	counter := 0
	substring := ""

	for i := 0; i < len(data); i++ {
		if data[i] == 32 {
			if substring != "" {
				spaces = append(spaces, counter)
				fields = append(fields, substring)
				substring = ""
				counter = 0
			}
			counter++
		} else {
			substring += string(data[i])
		}
	}

	spaces = append(spaces, counter)
	fields = append(fields, substring)

	return spaces, fields
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
	for i := index - 1; i >= 0; i-- {
		if spaces[i] == spaces[index]-tabLength {
			return i
		}
	}
	return -1
}

// Parse function parses given input and prints one liner hierarchy structures
/*
input: Expects kubectl explain --recursive output
patterns: Hierarchy to be computed for given patterns
caseSensitive: Case sensitive pattern match
noColor: If set, it won't print the colored output
*/
func Parse(input string, patterns []string, caseSensitive bool, noColor bool) {
	const Separator = "FIELDS:"
	const TabLength = 3

	start := getIndex(input, Separator)
	if start == -1 {
		fmt.Println(input)
		return
	}

	// Take only data under fields section
	data := input[start+len(Separator):]

	// Deletes all tabs, data type information and new lines from input for compressed processing
	const Trash = string(9) + "|" + `<[^>]*>|\n`
	data = findAndDelete(data, Trash)

	var spaces []int
	var fields []string

	// Process data for further analysis
	spaces, fields = analyze(data)

	// var pattern string
	var hierarchy string
	var index int

	// Compiled regex pattern for matching
	r := regexp.MustCompile(`(?i)` + strings.Join(patterns, "|"))
	if caseSensitive {
		r = regexp.MustCompile(strings.Join(patterns, "|"))
	}

	for i := 0; i < len(fields); i++ {
		if r.FindString(fields[i]) != "" {
			hierarchy = fields[i]
			index = i
			for index != -1 {
				index = findParentIndex(spaces, index, TabLength)
				if index != -1 {
					hierarchy = fields[index] + "." + hierarchy
				}
			}
			if noColor {
				fmt.Println(hierarchy)
			} else {
				color.Fill(hierarchy, r)
			}
		}
	}
}
