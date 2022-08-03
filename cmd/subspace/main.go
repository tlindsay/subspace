package main

import (
	"flag"
	"fmt"
	"github.com/tlindsay/subspace"
)

func main() {
	var numParagraphs int
	var numLines int
	var shouldPrintCharacters bool
	var shouldStartServer bool
	var character string

	flag.IntVar(&numParagraphs, "paragraphs", 1, "the number of paragraphs to print")
	flag.IntVar(&numParagraphs, "p", 1, "the number of paragraphs to print")
	flag.IntVar(&numLines, "lines", 3, "the number of lines to print")
	flag.IntVar(&numLines, "l", 3, "the number of lines to print")
	flag.StringVar(&character, "c", "picard", "the character whose dialog you want")
	flag.StringVar(&character, "character", "picard", "the character whose dialog you want")
	flag.BoolVar(&shouldPrintCharacters, "lc", false, "list the available characters")
	flag.BoolVar(&shouldPrintCharacters, "list-chars", false, "list the available characters")
	flag.BoolVar(&shouldStartServer, "s", false, "start JSON server")
	flag.BoolVar(&shouldStartServer, "serve", false, "start JSON server")
	flag.Parse()

	if shouldPrintCharacters {
		subspace.ListAllCharacters()
		return
	}

	if shouldStartServer {

	}

	text := subspace.MakeItSo(numParagraphs, numLines, character)
	fmt.Println(text)
}
