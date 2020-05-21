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

	keys := make([]string, 0, len(mm))
	for k := range mm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%x  %s\n", mm[k], k)
	}

	return nil
}

func md5Pipe(dir string) error {
	defer func(t time.Time) {
		fmt.Printf("Pipe %v\n", time.Since(t))
	}(time.Now())

	out, errc := md5Walk(dir)

	mm := make(map[string][md5.Size]byte)
	for d := range out {
		mm[d.file] = d.sum
	}
	keys := make([]string, 0, len(mm))
	for k := range mm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Printf("%x  %s\n", mm[k], k)
	}
	if err, ok := <-errc; ok {
		return err
	}

	return nil
}

func md5Walk(dir string) (<-chan digest, <-chan error) {
	out, errc := make(chan digest), make(chan error, 1)
	var wg sync.WaitGroup

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		errc <- err
		return out, errc
	}
	for _, f := range files {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			raw, err := ioutil.ReadFile(filepath.Join(dir, f))
			if err != nil {
				fmt.Println(err)
				errc <- err
				return
			}
			d := digest{
				file: f,
				err:  err,
			}
			d.digest(raw)
			out <- d
		}(f.Name())
	}

	go func() {
		wg.Wait()
		close(out)
		close(errc)
	}()

	return out, errc
}
