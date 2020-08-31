// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type BookInfo struct {
	Book  string
	Lines int
	<<!!YOUR_CODE!!>> Define the IBN field and corresponding struct tag
	Words int
}

// wc counts the number of words in a line of text.
func wc(line string) int {
	tokens := strings.Split(line, " ")

	return len(tokens)
}

func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var lines []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

<<!!YOUR_CODE!!>> -- Define and implement the hydrate function

func main() {
	b := BookInfo{
		Book: "assets/100west.txt",
	}
	<<!!YOUR_CODE!!>> -- Call your hydrate function to populate a BookInfo instance
	fmt.Printf("%#v\n", b)
}
