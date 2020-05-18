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
	IBN   string `md5:"ibn"`
	Words int
}

func wc(line string) int {
	tokens := strings.Split(line, " ")
	return len(tokens)
}

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

func hydrate(b interface{}) error {
	rb := reflect.Indirect(reflect.ValueOf(b))

	book := rb.FieldByName("Book")
	if !book.IsValid() {
		return errors.New("Unable to find field Book")
	}

	path := book.String()
	out, errC := readLines(path)
	var words, lines int
	for line := range out {
		words += wc(line)
		lines++
	}
	if err, ok := <-errC; ok && err != nil {
		return err
	}
	rb.FieldByName("Words").SetInt(int64(words))

	ibnT, _ := reflect.Indirect(rb).Type().FieldByName("IBN")
	if _, ok := ibnT.Tag.Lookup("sha1"); ok {
		rb.FieldByName("IBN").SetString(fmt.Sprintf("%x", sha1.Sum([]byte(path))))
	} else if _, ok := ibnT.Tag.Lookup("md5"); ok {
		rb.FieldByName("IBN").SetString(fmt.Sprintf("%x", md5.Sum([]byte(path))))
	}
	rb.FieldByName("Lines").SetInt(int64(lines))

	return nil
}

func main() {
	b := BookInfo{
		Book: "assets/100west.txt",
	}
	if err := hydrate(&b); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", b)
}
