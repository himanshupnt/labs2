// Copyright 2020 Imhotep Software
// All material is licensed under the Apache License Version 2.0
// http://www.apache.org/licenses/LICENSE-2.0

package grep_test

import (
	"io/ioutil"
	"testing"

	"github.com/gopherland/labs2/grep"
	"github.com/stretchr/testify/assert"
)

func TestWordCountV1(t *testing.T) {
	uu := usecases(t)
	t.Parallel()
	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, grep.WordCount("moby", u.text))
		})
	}
}

func TestWordCountV2(t *testing.T) {
	uu := usecases(t)
	t.Parallel()
	for k := range uu {
		u := uu[k]
		t.Run(k, func(t *testing.T) {
			assert.Equal(t, u.e, grep.WordCountBytes("moby", u.text))
		})
	}
}

func BenchmarkWordCountV1(b *testing.B) {
	raw, _ := ioutil.ReadFile("./assets/moby.txt")
	txt := string(raw)
	b.ReportAllocs()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		grep.WordCount("moby", txt)
	}
}

func BenchmarkWordCountV2(b *testing.B) {
	raw, _ := ioutil.ReadFile("./assets/moby.txt")
	txt := string(raw)
	b.ReportAllocs()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		grep.WordCountBytes("moby", txt)
	}
}

// Helpers...

type useCase struct {
	text string
	e    int64
}

type useCases map[string]useCase

func usecases(t testing.TB) useCases {
	samples := genSamples(t)
	return map[string]useCase{
		"semi-cols": {
			text: samples[0],
			e:    1,
		},
		"dash": {
			text: samples[1],
			e:    1,
		},
		"quotes": {
			text: samples[2],
			e:    1,
		},
		"special-dash": {
			text: samples[3],
			e:    2,
		},
	}
}

func genSamples(t testing.TB) []string {
	t.Helper()

	return []string{
		"*** START OF THIS PROJECT GUTENBERG EBOOK MOBY DICK; OR THE WHALE ***",
		"MOBY-DICK;",
		`“Moby Dick?” shouted Ahab. “Do ye know the white whale then, Tash?”`,
		"seen—Moby Dick—Moby Dick!”",
	}
}
