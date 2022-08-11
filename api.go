package subspace

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type handler struct{}

func StartServer(port int) {
	h := newHandler()
	fmt.Printf("Opening hailing frequencies on port %d...", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), h))
}

func newHandler() http.Handler {
	return handler{}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/characters" {
		j, err := json.Marshal(ListAllCharacters())
		if err != nil {
			http.Error(w, "Unknown error", 500)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
		return
	}
	q := r.URL.Query()
	numP, err := strconv.Atoi(q.Get("paragraphs"))
	if err != nil {
		http.Error(w, "Bad value for param \"paragraphs\"", 400)
		return
	} else if numP < 1 {
		numP = 1
	}
	numL, err := strconv.Atoi(q.Get("lines"))
	if err != nil {
		http.Error(w, "Bad value for param \"lines\"", 400)
		return
	} else if numL < 1 {
		numL = 1
	}
	char := r.URL.Query().Get("character")
	if char == "" {
		char = "picard"
	}

	if _, exist := CHARACTERS[char]; !exist {
		http.Error(w, "Character not supported. See valid characters at /characters", 400)
		return
	}

	output, err := MakeItSo(numP, numL, char)

	if err != nil {
		log.Printf("Unknown error occurred: %s\n", err)
		http.Error(w, "Unknown error", 500)
		return
	}

	if r.Header.Get("Content-Type") == "application/json" {
		j, err := json.Marshal(output)
		if err != nil {
			log.Printf("Unknown error occurred: %s\n", err)
			http.Error(w, "Unknown error", 500)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.Write(j)
		return
	}

	w.Write([]byte(strings.Join(output, "\n")))
}
