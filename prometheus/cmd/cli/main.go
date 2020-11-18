// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gopherland/labs2/prometheus/hangman"
)

func main() {
	var address string

	flag.StringVar(&address, "hm", "localhost:5000", "Specify a Hangman service host:port")
	flag.Parse()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigs
		fmt.Printf("\n\n 🙅‍♀️ No way!! Seriously??\n\n")
		os.Exit(0)
	}()

	player := hangman.NewPlayer(address)
	if err := player.Loop(); err != nil {
		panic(err)
	}
}
