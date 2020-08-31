package main

import (
	"context"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

const assetDir = "assets"

type digest struct {
	file string
	sum  [md5.Size]byte
}

func newDigest(f string, raw []byte) digest {
	return digest{
		file: f,
		sum:  md5.Sum(raw),
	}
}

func main() {
	var parallel bool
	flag.BoolVar(&parallel, "p", false, "Enable parallel mode")
	flag.Parse()

	if parallel {
		ctx, cancel := context.WithTimeout(context.Background(), 500*time.Microsecond)
		defer cancel()
		if err := md5ParallelControlled(ctx, os.Stdout, assetDir); err != nil {
			log.Fatal(err)
		}
		return
	}

	if err := md5Serial(os.Stdout, assetDir); err != nil {
		log.Fatal(err)
	}
}

func md5Serial(w io.Writer, dir string) error {
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
		d := newDigest(f.Name(), raw)
		mm[d.file] = d.sum
	}
	collate(w, mm)

	return nil
}

func md5ParallelFree(ctx context.Context, w io.Writer, dir string) error {
	fileChan, err1 := readDir(ctx, dir)
	out, err2 := computeDigestFree(ctx, dir, fileChan)
	select {
	case e, ok := <-err1:
		if ok {
			return e
		}
	case e, ok := <-err2:
		if ok {
			return e
		}
	}
	mm := make(map[string][md5.Size]byte)
	for d := range out {
		mm[d.file] = d.sum
	}
	collate(w, mm)

	return nil
}

func md5ParallelControlled(ctx context.Context, w io.Writer, dir string) error {
	fileChan, err1 := readDir(ctx, dir)
	out, err2 := computeDigestControlled(ctx, dir, fileChan)
	select {
	case e, ok := <-err1:
		if ok {
			return e
		}
	case e, ok := <-err2:
		if ok {
			return e
		}
	}
	mm := make(map[string][md5.Size]byte)
	for d := range out {
		mm[d.file] = d.sum
	}
	collate(w, mm)

	return nil
}

func computeDigestFree(ctx context.Context, dir string, in <-chan os.FileInfo) (<-chan digest, <-chan error) {
	out, errc := make(chan digest, 1), make(chan error, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for f := range in {
			select {
			case <-ctx.Done():
				errc <- ctx.Err()
				return
			default:
				wg.Add(1)
				go func(f string) {
					defer wg.Done()
					raw, err := ioutil.ReadFile(filepath.Join(dir, f))
					if err != nil {
						errc <- err
						return
					}
					out <- newDigest(f, raw)
				}(f.Name())
			}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
		close(errc)
	}()

	return out, errc
}

func computeDigestControlled(ctx context.Context, dir string, in <-chan os.FileInfo) (<-chan digest, <-chan error) {
	const numGOR = 3
	out, errc := make(chan digest, 1), make(chan error, 1)
	var wg sync.WaitGroup

	for i := 0; i < numGOR; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for f := range in {
				select {
				case <-ctx.Done():
					errc <- ctx.Err()
					return
				default:
					raw, err := ioutil.ReadFile(filepath.Join(dir, f.Name()))
					if err != nil {
						errc <- err
						return
					}
					out <- newDigest(f.Name(), raw)
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(out)
		close(errc)
	}()

	return out, errc
}

func readDir(ctx context.Context, dir string) (<-chan os.FileInfo, <-chan error) {
	out, errc := make(chan os.FileInfo, 2), make(chan error, 1)
	var wg sync.WaitGroup

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		errc <- err
		close(out)
		return out, errc
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, f := range files {
			select {
			case out <- f:
			case <-ctx.Done():
				errc <- ctx.Err()
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(out)
		close(errc)
	}()

	return out, errc
}

func collate(w io.Writer, mm map[string][md5.Size]byte) {
	keys := make([]string, 0, len(mm))
	for k := range mm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(w, "%x  %s\n", mm[k], k)
	}
}
