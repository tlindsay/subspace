package main

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/fastly/compute-sdk-go/fsthttp"
	easyjson "github.com/mailru/easyjson"
	"github.com/tlindsay/subspace"
)

// The entry point for your application.
//
// Use this function to define your main request handling logic. It could be
// used to route based on the request properties (such as method or path), send
// the request to a backend, make completely new requests, and/or generate
// synthetic responses.

type JsonMeta struct {
	Paragraphs int `json:"numParagraphs"`
	Lines      int `json:"linesPerParagraph"`
}
type JsonResponse struct {
	Character string   `json:"character"`
	Text      []string `json:"text"`
	Meta      JsonMeta `json:"meta"`
}

func Error(w fsthttp.ResponseWriter, msg string, code int) {
	w.WriteHeader(code)
	fmt.Fprintln(w, msg)
}

func main() {
	fsthttp.ServeFunc(func(_ctx context.Context, w fsthttp.ResponseWriter, r *fsthttp.Request) {
		// Filter requests that have unexpected methods.
		if r.Method != "HEAD" && r.Method != "GET" {
			w.WriteHeader(fsthttp.StatusMethodNotAllowed)
			fmt.Fprintf(w, "This method is not allowed\n")
			return
		}

		// if r.URL.Path == "/characters" {
		// 	j, err := easyjson.Marshal(subspace.ListAllCharacters())
		// 	if err != nil {
		// 		Error(w, "Unknown error", 500)
		// 		return
		// 	}
		// 	w.Header().Add("Content-Type", "application/json")
		// 	w.Write(j)
		// 	return
		// }

		q := r.URL.Query()

		numP, err := strconv.Atoi(q.Get("paragraphs"))
		if err != nil {
			Error(w, "Bad value for param \"paragraphs\"", 400)
			return
		} else if numP < 1 {
			numP = 1
		}

		numL, err := strconv.Atoi(q.Get("lines"))
		if err != nil {
			Error(w, "Bad value for param \"lines\"", 400)
			return
		} else if numL < 1 {
			numL = 1
		}

		char := r.URL.Query().Get("character")
		if char == "" {
			char = "picard"
		}
		if _, exist := subspace.CHARACTERS[char]; !exist {
			Error(w, "Character not supported. See valid characters at /characters", 400)
			return
		}

		output, err := subspace.MakeItSo(numP, numL, char)

		if err != nil {
			fmt.Printf("Unknown error occurred: %s\n", err)
			Error(w, "Unknown error", 500)
			return
		}

		if r.Header.Get("Accept") == "application/json" {

			j, err := easyjson.Marshal(JsonResponse{Character: char, Text: output, Meta: JsonMeta{Lines: numL, Paragraphs: numP}})
			if err != nil {
				fmt.Printf("Unknown error occurred: %s\n", err)
				Error(w, "Unknown error", 500)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.Write(j)
			return
		}

		w.Write([]byte(strings.Join(output, "\n")))
	})
}
