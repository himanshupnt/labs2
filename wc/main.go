package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"unicode"
	// "github.com/pkg/profile"
)

func main() {
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1), profile.ProfilePath(".")).Stop()

	f, err := os.Open(filepath.Join(os.Args[1]))
	if err != nil {
		panic(err)
	}
	var (
		count  int
		inWord bool
		rr     = bufio.NewReader(f)
		bb     = make([]byte, 1)
	)

	for {
		_, err := rr.Read(bb)
		r := rune(bb[0])
		// r, err := readByte(b, bb)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("read failed %q -- %v", os.Args[1], err)
		}
		if unicode.IsSpace(r) && inWord {
			count, inWord = count+1, false
			continue
		}
		inWord = unicode.IsLetter(r) ||
			unicode.IsNumber(r) ||
			unicode.IsSymbol(r) ||
			unicode.IsPunct(r)
	}
	fmt.Printf("  %d %s\n", count, os.Args[1])
}

func readByte(r io.Reader, bb []byte) (rune, error) {
	// var b [1]byte
	// _, err := r.Read(b[:])
	_, err := r.Read(bb)

	return rune(bb[0]), err
}
