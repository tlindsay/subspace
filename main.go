package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var ALL_LINES []Line

var CHARACTERS = map[string]string{
	"beverly": "beverly.json",
	"data":    "data.json",
	"geordi":  "geordi.json",
	"guinan":  "guinan.json",
	"obrien":  "obrien.json",
	"picard":  "picard.json",
	"riker":   "riker.json",
	"troi":    "troi.json",
	"wesley":  "wesley.json",
	"worf":    "worf.json",
	"tasha":     "yar.json",
}

type Paragraph struct {
	Lines []Line
}
type Line struct {
	Text      string `json:"text"`
	Episode   string `json:"episode"`
	WordCount int16  `json:"word_count"`
}

func main() {
	var numParagraphs int
	var numLines int
	var shouldPrintCharacters bool
	var character string
	flag.IntVar(&numParagraphs, "paragraphs", 1, "the number of paragraphs to print")
	flag.IntVar(&numParagraphs, "p", 1, "the number of paragraphs to print")
	flag.IntVar(&numLines, "lines", 3, "the number of lines to print")
	flag.IntVar(&numLines, "l", 3, "the number of lines to print")
	flag.StringVar(&character, "c", "picard", "the character whose dialog you want")
	flag.StringVar(&character, "character", "picard", "the character whose dialog you want")
	flag.BoolVar(&shouldPrintCharacters, "lc", false, "list the available characters")
	flag.BoolVar(&shouldPrintCharacters, "list-chars", false, "list the available characters")
	flag.Parse()

	if shouldPrintCharacters {
		listAllCharacters()
		return
	}

	err := getLines(character)
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().Unix())

	var paragraphs []Paragraph
	for i := 0; i < numParagraphs; i++ {
		p, err := getParagraph(numLines)
		if err != nil {
			panic(err)
		}

		paragraphs = append(paragraphs, p)
	}

	for _, p := range paragraphs {
		printParagraph(p)
		fmt.Println()
	}
}

func listAllCharacters() {
	for name := range CHARACTERS {
		fmt.Println(strings.Title(name))
	}
}

func printParagraph(p Paragraph) {
	printLines(p.Lines)
}

func getParagraph(numLines int) (Paragraph, error) {
	var p Paragraph
	l, err := getRandomLines(numLines)
	if err != nil {
		return p, err
	}
	p.Lines = l
	return p, nil
}

func printLines(lines []Line) {
	var toPrint []string
	for _, l := range lines {
		toPrint = append(toPrint, l.Text)
	}
	fmt.Println(strings.Join(toPrint, " "))
}

func getRandomLines(numLines int) ([]Line, error) {
	var randomLines []Line

	for i := 0; i < numLines; i++ {
		randomLines = append(randomLines, ALL_LINES[rand.Intn(len(ALL_LINES))])
	}
	return randomLines, nil
}

func getLines(char string) error {
	f, err := os.Open("./lines/"+char+".json")
	if err != nil {
		return err
	}

	d := json.NewDecoder(f)

	if err := d.Decode(&ALL_LINES); err != nil {
		return err
	}

	return nil
}
