// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"bytes"
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type templateFn func() string

func main() {
	var pkg, types string
	flag.StringVar(&types, "t", "", "Specify one or more types column separated")
	flag.StringVar(&pkg, "p", "stacks", "Specify the stack package name")
	flag.Parse()

	generate(pkg, types)
}

func generate(pkg, types string) {
	tt := strings.Split(types, ",")
	if len(tt) == 0 {
		panic("you must specify some stack types")
	}
	for _, t := range tt {
		log.Printf("Generating Stack for %s::%s", pkg, t)
		if err := generateFiles(pkg, t); err != nil {
			log.Panic(err)
		}
	}
}

func generateFiles(pkg, ztype string) error {
	if err := genFile(pkg, ztype, false, goFile); err != nil {
		return err
	}

	return genFile(pkg, ztype, true, goTestFile)
}

func genFile(pkg, ztype string, isTest bool, fn templateFn) error {
	if err := ensureDir(pkg); err != nil {
		return err
	}

	<<!!YOUR_CODE!!>>
	  - Hydrate your template using the provided template function (refer to the template lab!)
		- Note: this time we will need to hydrate the template into a bytes.Buffer to store the generated template
		- Scan the stack template (tpl.go) and see what data structure and template function the provided template needs
		- Setup your data structure and template function map
		- Execute the template so that your bytes buffer (buff) gets populated
		- The provided code will create the file and save it into the provided package directory.

	fileName := ztype
	if isTest {
		fileName = ztype + "_test"
	}
	fileName += ".go"

	return ioutil.WriteFile(filepath.Join(pkg, fileName), buff.Bytes(), 0644)
}

// Helpers...

func ensureDir(dir string) error {
	if err := os.Mkdir(dir, 0755); err != nil {
		if err == os.ErrExist {
			return err
		}
	}

	return nil
}

func toTypeName(t string) string {
	return strings.Title(t)
}
