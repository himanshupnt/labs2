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

<<!!YOUR_CODE!!>> -- Define and implement the hydrate function

func main() {
	b := BookInfo{
		Book: "assets/100west.txt",
	}
	if err := hydrate(&b); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", b)
}
