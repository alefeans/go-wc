package main

import (
	"strings"
	"testing"
)

func TestCountAll(t *testing.T) {
	tests := []struct {
		input string
		want  *Counter
	}{
		{input: "", want: &Counter{bytes: 0, lines: 0, words: 0, chars: 0}},
		{input: "word", want: &Counter{bytes: 4, lines: 0, words: 1, chars: 4}},
		{input: "two words", want: &Counter{bytes: 9, lines: 0, words: 2, chars: 9}},
		{input: "encõding test", want: &Counter{bytes: 14, lines: 0, words: 2, chars: 13}},
		{input: "one\nline", want: &Counter{bytes: 8, lines: 1, words: 2, chars: 8}},
		{input: "two\nline\ns", want: &Counter{bytes: 10, lines: 2, words: 3, chars: 10}},
		{input: "trailing\nli\nne\n", want: &Counter{bytes: 15, lines: 3, words: 3, chars: 15}},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		got := countAll(input)
		if *got != *test.want {
			t.Errorf("got %+v, want %+v", *got, *test.want)
		}
	}
}

func BenchmarkCountAllSimpleInput(b *testing.B) {
	b.StopTimer()
	input := strings.NewReader("Testing simple input")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		countAll(input)
	}
}

func BenchmarkCountAllTestFile(b *testing.B) {
	b.StopTimer()
	input, _ := openFile("test.txt")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		countAll(input)
	}
}
