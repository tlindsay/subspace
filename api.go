package subspace

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/tlindsay/subspace/response"
)

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
			j, err := json.Marshal(ListAllCharacters())
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

		numP, err := strconv.Atoi(q.Get("paragraphs"))
		if err != nil {
			http.Error(w, "Bad value for param \"paragraphs\"", 400)
			logError.Printf("%s\n", err)
			return
		} else if numP < 1 {
			numP = 1
		}

		numL, err := strconv.Atoi(q.Get("lines"))
		if err != nil {
			http.Error(w, "Bad value for param \"lines\"", 400)
			logError.Printf("%s\n", err)
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
			logError.Printf("Character %s not supported\n", char)
			return
		}

		output, err := MakeItSo(numP, numL, char)

		if err != nil {
			logError.Printf("Unknown error occurred: %s\n", err)
			http.Error(w, "Unknown error", 500)
			return
		}

		if r.Header.Get("Accept") == "application/json" {
			j, err := json.Marshal(response.JsonResponse{Character: char, Text: output, Meta: response.JsonMeta{Lines: numL, Paragraphs: numP}})
			if err != nil {
				logError.Printf("Unknown error occurred: %s\n", err)
				http.Error(w, "Unknown error", 500)
				return
			}
			w.Header().Add("Content-Type", "application/json")
			w.Write(j)
			return
		}

		w.Write([]byte(strings.Join(output, "\n")))
	}
}
