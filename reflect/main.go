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

// readLines read lines from a file on disk and returns a chan containing each lines read or an error.
func readLines(path string) (chan string, chan error) {
	out, errC := make(chan string), make(chan error, 1)

	go func(path string, out chan<- string, errC chan<- error) {
		defer func() {
			close(out)
			close(errC)
		}()
		f, err := os.Open(path)
		if err != nil {
			errC <- err
			return
		}

		sc := bufio.NewScanner(f)
		for sc.Scan() {
			out <- sc.Text()
		}
		if err := sc.Err(); err != nil {
			errC <- err
		}
	}(path, out, errC)

	return out, errC
}

<<!!YOUR_CODE!!>> -- Define and implement the hydrate function

func main() {
	b := BookInfo{
		Book: "assets/100west.txt",
	}
	<<!!YOUR_CODE!!>> -- Call your hydrate function to populate a BookInfo instance
	fmt.Printf("%#v\n", b)
}
