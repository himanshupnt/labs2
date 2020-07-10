package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

type digest struct {
	file string
	sum  [md5.Size]byte
	err  error
}

func (d *digest) digest(raw []byte) {
	d.sum = md5.Sum(raw)
}

func main() {
	var parallel bool
	flag.BoolVar(&parallel, "p", false, "Toggle the run mode")
	flag.Parse()

	if parallel {
		if err := md5Pipe("assets"); err != nil {
			panic(err)
		}
		return
	}

	if err := md5Serial("assets"); err != nil {
		panic(err)
	}
}

func md5Serial(dir string) error {
	defer func(t time.Time) {
		fmt.Printf("Serial %v\n", time.Since(t))
	}(time.Now())

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	mm := make(map[string][md5.Size]byte)
	for _, f := range files {
		raw, err := ioutil.ReadFile(filepath.Join(dir, f.Name()))
		if err != nil {
			return err
		}
		d := digest{file: f.Name()}
		d.digest(raw)
		mm[d.file] = d.sum
	}
  collate(mm)

	return nil
}

func collate(mm map[string][md5.Size]byte) {
	keys := make([]string, 0, len(mm))
	for k := range mm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%x  %s\n", mm[k], k)
	}
}

func md5Pipe(dir string) error {
	<<!!YOUR_CODE!!>> -- implement a single stage pipeline to process digests
}
