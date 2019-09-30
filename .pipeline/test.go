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

// This package is used only to run build tests on drone CI
package main

import (
	"fmt"
	"os"
	"io"
	"flag"
	"bufio"

	"github.com/rewanth1997/kubectl-fields/pkg/fields"
)

const (
    Author      = "Rewanth Cool"    // Package author: first name - last name
    Version     = "0.2"       // Package version: major.minor.maintenance.revision
    ReleaseDate = "2019-30-09"      // Release date: year-month-day
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
	
	fields.Parse(input, os.Args[1:], *ignoreCase)
}