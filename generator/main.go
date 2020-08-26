// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"fmt"
	"log"

	"github.com/gookit/color"
	<<!!YOUR_CODE!!> -- Specify your generic stacks package
)

<!!YOUR_CODE!!> -- Add stacker generator to generate the generic stacks
func main() {
	fmt.Print("\033[H\033[2J")

	tryFloat()
	tryInt()
}

func tryFloat() {
	s := stacks.Float64{}
	for _, v := range []float64{10.5, 20.2, 42.25} {
		s.Push(v)
	}
	fmt.Printf("🥞 %-7s Pushed:%v -- ", "Float64", &s)
	v, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}

	cyan, green := color.FgCyan.Render, color.FgGreen.Render
	fmt.Printf("Pop:%v - Top:%s --> Peek:%v\n", cyan(v), green(s.Top()), &s)
}

func tryInt() {
	s := stacks.Int32{}
	for _, v := range []int32{200, 100, 300} {
		s.Push(v)
	}
	fmt.Printf("📚 %-7s Pushed:%v     -- ", "Int32", &s)
	v, err := s.Pop()
	if err != nil {
		log.Fatal(err)
	}

	cyan, green := color.FgCyan.Render, color.FgGreen.Render
	fmt.Printf("Pop:%v   - Top:%s  --> Peek:%v\n", cyan(v), green(s.Top()), &s)
}