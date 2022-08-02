package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var MIN_WORD_COUNT = 3

type Line struct {
	Text      string `json:"text"`
	Episode   string `json:"episode"`
	WordCount int16  `json:"word_count"`
}

func main() {
	lines, err := getLines()
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().Unix())
	ipsum, err := getRandomLines(3, lines)
	printLines(ipsum)
}

func printLines(lines []Line) {
	var toPrint []string
		for _, l := range lines {
			toPrint = append(toPrint, l.Text)
	}
	fmt.Println(strings.Join(toPrint, " "))
}

func getRandomLines(numLines int, lines []Line) ([]Line, error) {
	var randomLines []Line
	for i := 0; i < numLines; i++{
		randomLines = append(randomLines, lines[rand.Intn(len(lines))])
	}
	return randomLines, nil
}

func getLines() ([]Line, error) {
f, err := os.Open("./picard.json")
	if err != nil {
		return nil, err
	}

	d := json.NewDecoder(f)
	var lines []Line

	if err := d.Decode(&lines); err != nil {
		return nil, err
	}

	return lines, nil
}
