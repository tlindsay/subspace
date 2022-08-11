package subspace

import (
	"embed"
	"encoding/json"
	"math/rand"
	"strings"
	"time"
)

//go:embed assets/*.json
var fs embed.FS

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
	"tasha":   "yar.json",
}

type Paragraph struct {
	Lines []Line
}
type Line struct {
	Text      string `json:"text"`
	Episode   string `json:"episode"`
	WordCount int16  `json:"word_count"`
}

func MakeItSo(numParagraphs int, numLines int, character string) ([]string, error) {
	err := loadLines(character)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().Unix())

	var paragraphs []Paragraph
	for i := 0; i < numParagraphs; i++ {
		p, err := getParagraph(numLines)
		if err != nil {
			return nil, err
		}

		paragraphs = append(paragraphs, p)
	}

	var output []string
	for _, p := range paragraphs {
		output = append(output, p.String())
	}

	return output, nil
}

func ListAllCharacters() []string {
	var chars []string
	for name := range CHARACTERS {
		chars = append(chars, strings.Title(name))
	}
	return chars
}

func (l Line) String() string {
	return l.Text
}

func (p Paragraph) String() string {
	var lines []string
	for _, l := range p.Lines {
		lines = append(lines, l.String())
	}
	return strings.Join(lines, " ")
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

func getRandomLines(numLines int) ([]Line, error) {
	var randomLines []Line

	for i := 0; i < numLines; i++ {
		randomLines = append(randomLines, ALL_LINES[rand.Intn(len(ALL_LINES))])
	}
	return randomLines, nil
}

func loadLines(char string) error {
	f, err := fs.ReadFile("assets/" + char + ".json")
	if err != nil {
		return err
	}

	if err := json.Unmarshal(f, &ALL_LINES); err != nil {
		return err
	}

	return nil
}
