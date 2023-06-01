package subspace

import (
	"embed"
	"math/rand"
	"strings"
	"time"

	easyjson "github.com/mailru/easyjson"
)

//go:embed assets/*.json
var fs embed.FS

//easyjson:json
type LineSlice []Line

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
	WordCount int    `json:"word_count"`
}

func MakeItSo(numParagraphs int, numLines int, character string) ([]string, error) {
	lines, err := loadLines(character)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().Unix())

	var paragraphs []Paragraph
	for i := 0; i < numParagraphs; i++ {
		p := getParagraph(lines, numLines)
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

func getParagraph(lines LineSlice, numLines int) Paragraph {
	var p Paragraph
	l := getRandomLines(lines, numLines)
	p.Lines = l
	return p
}

func getRandomLines(lines LineSlice, numLines int) []Line {
	var randomLines []Line

	for i := 0; i < numLines; i++ {
		randomLines = append(randomLines, lines[rand.Intn(len(lines))])
	}
	return randomLines
}

func loadLines(char string) (LineSlice, error) {
	f, err := fs.ReadFile("assets/" + char + ".json")
	if err != nil {
		return nil, err
	}

	var lines LineSlice
	if err := easyjson.Unmarshal(f, &lines); err != nil {
		return nil, err
	}

	return lines, nil
}
