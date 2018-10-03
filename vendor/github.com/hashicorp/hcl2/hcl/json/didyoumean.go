package json

import (
	"github.com/agext/levenshtein"
)

var keywords = []string{"false", "true", "null"}

// keywordSuggestion tries to find a valid JSON keyword that is close to the
// given string and returns it if found. If no keyword is close enough, returns
// the empty string.
func keywordSuggestion(given string) string {
	return nameSuggestion(given, keywords)
}

// nameSuggestion tries to find a name from the given slice of suggested names
// that is close to the given name and returns it if found. If no suggestion
// is close enough, returns the empty string.
//
// The suggestions are tried in order, so earlier suggestions take precedence
// if the given string is similar to two or more suggestions.
//
// This function is intended to be used with a relatively-small number of
// suggestions. It's not optimized for hundreds or thousands of them.
func nameSuggestion(given string, suggestions []string) string {
	for _, suggestion := range suggestions {
		dist := levenshtein.Distance(given, suggestion, nil)
		if dist < 3 { // threshold determined experimentally
			return suggestion
		}
	}
	return ""
}
