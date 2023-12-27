package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

type Args struct {
	bytesCount bool
	linesCount bool
	wordsCount bool
	charsCount bool
}

func parseCliArgs() *Args {
	var args Args
	flag.BoolVar(&args.bytesCount, "c", false, "Prints the number of bytes")
	flag.BoolVar(&args.linesCount, "l", false, "Prints the number of lines")
	flag.BoolVar(&args.wordsCount, "w", false, "Prints the number of words")
	flag.BoolVar(&args.charsCount, "m", false, "Prints the number of characters")
	flag.Parse()
	return &args
}

func openFile(filename string) (*os.File, error) {
	if filename == "" {
		return os.Stdin, nil
	}
	return os.Open(filename)
}

type Counter struct {
	bytes int
	lines int
	words int
	chars int
}

func countAll(file io.Reader) *Counter {
	var counter Counter
	inWord := false
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		scanned := scanner.Text()

		counter.chars++
		counter.bytes += len(scanned)

		if scanned == "\n" {
			counter.lines++
			inWord = false
		} else if scanned == " " || scanned == "\t" {
			inWord = false
		} else {
			if !inWord {
				counter.words++
				inWord = true
			}
		}
	}

	return &counter
}

func printCount(filename string, args *Args, counter *Counter) {
	switch {
	case args.bytesCount:
		fmt.Printf("%d %s\n", counter.bytes, filename)
	case args.linesCount:
		fmt.Printf("%d %s\n", counter.lines, filename)
	case args.wordsCount:
		fmt.Printf("%d %s\n", counter.words, filename)
	case args.charsCount:
		fmt.Printf("%d %s\n", counter.chars, filename)
	default:
		fmt.Printf("%d %d %d %s\n", counter.lines, counter.words, counter.bytes, filename)
	}
}

func main() {
	args := parseCliArgs()

	filename := flag.Arg(0)
	file, err := openFile(filename)
	if err != nil {
		fmt.Printf("can't open file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	printCount(filename, args, countAll(file))
}
