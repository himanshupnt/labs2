package main

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

const expected = `8d0c9e284adb61f1ce136507b97c01d0  100west.txt
e6011247f142af47a8c2d7fdb930c931  13chil.txt
a4ac47adb5f1b06052056a29b87b694f  3lpigs.txt
6af2ad751fd500d501202181b1c843d8  3student.txt
`

func TestSerial(t *testing.T) {
	var b bytes.Buffer
	assert.Nil(t, md5Serial(&b, assetDir))
	assert.Equal(t, expected, b.String())
}

func TestParallelFree(t *testing.T) {
	var b bytes.Buffer
	assert.Nil(t, md5ParallelFree(context.Background(), &b, assetDir))
	assert.Equal(t, expected, b.String())
}

func TestParallelControlled(t *testing.T) {
	var b bytes.Buffer
	assert.Nil(t, md5ParallelControlled(context.Background(), &b, assetDir))
	assert.Equal(t, expected, b.String())
}

func BenchmarkSerial(b *testing.B) {
	for n := 0; n < b.N; n++ {
		md5Serial(ioutil.Discard, assetDir)
	}
}

func BenchmarkParallelFree(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		md5ParallelFree(ctx, ioutil.Discard, assetDir)
	}
}

func BenchmarkParallelControlled(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		md5ParallelControlled(ctx, ioutil.Discard, assetDir)
	}
}
