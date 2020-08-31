// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type BookInfo struct {
	Book  string
	Lines int
	IBN   string `ibn:"md5"`
	Words int
}

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

func hydrate(b interface{}) error {
	rb := reflect.Indirect(reflect.ValueOf(b))

	book := rb.FieldByName("Book")
	if !book.IsValid() {
		return errors.New("Unable to find field Book")
	}

	path := book.String()
	lines, err := readLines(path)
	if err != nil {
		return err
	}
	var wordCount, lineCount int
	for _, line := range lines {
		wordCount += wc(line)
		lineCount++
	}
	rb.FieldByName("Words").SetInt(int64(wordCount))
	rb.FieldByName("Lines").SetInt(int64(lineCount))

	ibnT, _ := reflect.Indirect(rb).Type().FieldByName("IBN")
	if tag, ok := ibnT.Tag.Lookup("ibn"); ok {
		fmt.Printf("%#v\n", tag)
		rb.FieldByName("IBN").SetString(fmt.Sprintf("%x", sha1.Sum([]byte(path))))
	} else if _, ok := ibnT.Tag.Lookup("md5"); ok {
		rb.FieldByName("IBN").SetString(fmt.Sprintf("%x", md5.Sum([]byte(path))))
	}

	return nil
}

func main() {
	b := BookInfo{
		Book: "assets/100west.txt",
	}
	if err := hydrate(&b); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", b)
}
