package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ahmedkhaeld/translator-api/translation"
)

type Resp struct {
	Language    string `json:"language"`
	Translation string `json:"translation"`
}

// TranslateHandler translates a word from one language to another
func TranslateHandler(w http.ResponseWriter, r *http.Request) {
	// create a new JSON encoder
	enc := json.NewEncoder(w)
	// set the content type to JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// get the language from the URL query parameter
	language := r.URL.Query().Get("language")
	// if the language is not specified, set it to English
	if language == "" {
		language = "english"
	}
	// get the word from the URL path
	word := strings.ReplaceAll(r.URL.Path, "/", "")
	// translate the word to the specified language
	translation := translation.Translate(word, language)
	// if the translation is empty, set the language to empty and return a 404 status code
	if translation == "" {
		language = ""
		w.WriteHeader(404)
		return
	}
	// create a response struct with the translated word and language
	resp := Resp{
		Language:    language,
		Translation: translation,
	}
	// encode the response struct as JSON and write it to the response writer
	if err := enc.Encode(resp); err != nil {
		panic("unable to encode response")
	}
}
