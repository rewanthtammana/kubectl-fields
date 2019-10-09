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

// Package color is used to add colors to stdout
package color

import (
	"log"
	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
	"index/suffixarray"
	"regexp"
)

// Sets colorable output for all terminals
func init() {
	log.SetOutput(colorable.NewColorableStdout())
	log.SetFlags(0)
}

// Fill function colors the matching patterns in a string
/*
s: Base string
r: Compiled regex pattern for matching
*/
func Fill(s string, r *regexp.Regexp) {
	red := color.New(color.FgRed).SprintFunc()

	index := suffixarray.New([]byte(s))
	res := index.FindAllIndex(r, -1)

	newstr := ""
	old := 0

	for _, v := range res {
		newstr = newstr + s[old:v[0]]
		newstr = newstr + red(s[v[0]:v[1]])
		old = v[1]
	}
	newstr = newstr + s[old:]

	log.Print(newstr)
}
