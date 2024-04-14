package game

import (
	"math/rand"
	
	"github.com/scribble-rs/scribble.rs/internal/config"
	
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	wordListCache       = make(map[string][]string)
	languageIdentifiers = map[string]string{
		"english_gb": "en_gb",
		"english":    "en_us",
		"italian":    "it",
		"german":     "de",
		"french":     "fr",
		"dutch":      "nl",
	}

)

type LanguageData struct {
	Lowercaser   func() cases.Caser
	LanguageCode string
}

var (
	WordlistData       = map[string]LanguageData{
		"english_gb": {
			LanguageCode: "en_gb",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.BritishEnglish) },
		},
		"english": {
			LanguageCode: "en_us",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.AmericanEnglish) },
		},
		"italian": {
			LanguageCode: "it",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Italian) },
		},
		"german": {
			LanguageCode: "de",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.German) },
		},
		"french": {
			LanguageCode: "fr",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.French) },
		},
		"dutch": {
			LanguageCode: "nl",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Dutch) },
		},
		"ukrainian": {
			LanguageCode: "ua",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Ukrainian) },
		},
	}
)

func getLanguageIdentifier(language string) string {
	return languageIdentifiers[language]
}

// GetRandomWords gets a custom amount of random words for the passed Lobby.
// The words will be chosen from the custom words and the default
// dictionary, depending on the settings specified by the lobbies creator.
func GetRandomWords(wordCount int, lobby *Lobby) []string {
	allwords := config.GetAllWords(lobby.WordGroups)
	words := config.RandNUniqueOfSlice(wordCount,allwords)
	
	return words
}

func shuffleWordList(wordlist []string) {
	rand.Shuffle(len(wordlist), func(a, b int) {
		wordlist[a], wordlist[b] = wordlist[b], wordlist[a]
	})
}
