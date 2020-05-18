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
	IBN   string `sha1:"ibn"`
	Words int
}

func wc(line string) int {
	tokens := strings.Split(line, " ")
	return len(tokens)
}

func hydrate(b interface{}) error {
	rb := reflect.Indirect(reflect.ValueOf(b))

	file := rb.FieldByName("Book").String()
	f, err := os.Open(file)
	if err != nil {
		return err
	}

	var count, line int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		count += wc(sc.Text())
		line++
	}
	rb.FieldByName("Words").SetInt(int64(count))

	ibnT, _ := reflect.Indirect(rb).Type().FieldByName("IBN")
	if _, ok := ibnT.Tag.Lookup("sha1"); ok {
		rb.FieldByName("IBN").SetString(fmt.Sprintf("%x", sha1.Sum([]byte(file))))
	} else if _, ok := ibnT.Tag.Lookup("md5"); ok {
		rb.FieldByName("IBN").SetString(fmt.Sprintf("%x", md5.Sum([]byte(file))))
	}
	rb.FieldByName("Lines").SetInt(int64(line))

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
