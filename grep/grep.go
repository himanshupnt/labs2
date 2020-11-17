// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package grep

import (
	"regexp"
	"strings"
)

// WordCount returns the number of occurrence of a word in a line.
func WordCount(word, line string) int {
	var rx = regexp.MustCompile(`(?i)` + word)
	ss := rx.FindAllStringIndex(line, -1)

	return len(ss)
}

// WordCountBytes returns the number of occurrence of a word in a line.
func WordCountBytes(word, line string) (count int) {
	var index int
	for _, b := range []byte(strings.ToLower(line)) {
		if b != word[index] {
			index = 0
			continue
		}
		index++
		if index == len(word) {
			count++
			index = 0
		}
	}
	return
}
