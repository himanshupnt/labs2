package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var R rune

func TestReadByte(t *testing.T) {
	buff := bytes.NewBuffer([]byte("a"))

	bb := make([]byte, 1)
	b, err := readByte(buff, bb)
	assert.Nil(t, err)
	assert.Equal(t, 'a', b)
}

func BenchmarkReadByte(b *testing.B) {
	buff := bytes.NewBuffer([]byte("a"))

	bb := make([]byte, 1)

	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		R, _ = readByte(buff, bb)
	}
}
