package api

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tlindsay/subspace/response"
	"github.com/tlindsay/subspace/subspace"
)

var DEFAULT_PARAGRAPHS = 3
var DEFAULT_LINES = 5

func StartServer(port int) {
	fmt.Printf("Opening hailing frequencies on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), Handler()))
}

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prefix := fmt.Sprintf("%s | %s | ", os.Getenv("FASTLY_SERVICE_VERSION"), r.RemoteAddr)
		logInfo := log.New(os.Stdout, prefix, log.LstdFlags|log.LUTC)
		logError := log.New(os.Stderr, prefix, log.LstdFlags|log.LUTC)

		logInfo.Printf("[Info] Received %s", r.URL)
		if r.URL.Path == "/characters" {
			j, err := json.Marshal(subspace.ListAllCharacters())
			if err != nil {
				http.Error(w, "Unknown error", 500)
				logError.Printf("%s\n", err)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.Write(j)
			return
		}

		q := r.URL.Query()

		var numP int
		p := q.Get("paragraphs")
		if p == "" {
			numP = DEFAULT_PARAGRAPHS
		} else {
			var err error
			numP, err = strconv.Atoi(p)

			if err != nil {
				http.Error(w, "Bad value for param \"paragraphs\"", 400)
				logError.Printf("%s\n", err)
				return
			} else if numP < 1 {
				numP = DEFAULT_PARAGRAPHS
			}
		}

		var numL int
		l := q.Get("lines")
		if l == "" {
			numL = DEFAULT_LINES
		} else {
			var err error
			numL, err = strconv.Atoi(l)

			if err != nil {
				http.Error(w, "Bad value for param \"lines\"", 400)
				logError.Printf("%s\n", err)
				return
			} else if numL < 1 {
				numL = DEFAULT_LINES
			}
		}

		char := r.URL.Query().Get("character")
		if char == "" {
			logInfo.Println("Selecting random character...")
			rand.Seed(time.Now().Unix())
			charKeys := []string{}
			for key, _ := range subspace.CHARACTERS {
				charKeys = append(charKeys, key)
			}
			idx := rand.Intn(len(charKeys))
			char = charKeys[idx]
		}
		if _, exist := subspace.CHARACTERS[char]; !exist {
			http.Error(w, "Character not supported. See valid characters at /characters", 400)
			logError.Printf("Character %s not supported\n", char)
			return
		}

		logInfo.Printf("Getting lines for character: %s\n", char)
		output, err := subspace.MakeItSo(numP, numL, char)

		if err != nil {
			logError.Printf("Unknown error occurred: %s\n", err)
			http.Error(w, "Unknown error", 500)
			return
		}

		logInfo.Printf("Returning data: %s\n", output[0])
		logInfo.Printf("Paragraphs: %d, Lines: %d\n", numP, numL)

		if r.Header.Get("Accept") == "application/json" {
			j, err := json.Marshal(response.JsonResponse{Character: char, Text: output, Meta: response.JsonMeta{Lines: numL, Paragraphs: numP}})
			if err != nil {
				logError.Printf("Unknown error occurred: %s\n", err)
				http.Error(w, "Unknown error", 500)
				return
			}

			w.Header().Add("Content-Type", "application/json")
			w.Write(j)
		} else {
			w.Write([]byte(strings.Join(output, "\n")))
		}
	}
}
