package game

import (
	"embed"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/text/cases"
	"github.com/scribble-rs/scribble.rs/internal/config"
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

	//go:embed words/*
	wordFS embed.FS
)

func getLanguageIdentifier(language string) string {
	return languageIdentifiers[language]
}

// readWordListInternal exists for testing purposes, it allows passing a custom
// wordListSupplier, in order to avoid having to write tests aggainst the
// default language lists.
func readWordListInternal(
	lowercaser cases.Caser, chosenLanguage string,
	wordlistSupplier func(string) (string, error),
) ([]string, error) {
	languageIdentifier := getLanguageIdentifier(chosenLanguage)
	words, available := wordListCache[languageIdentifier]
	if !available {
		wordListFile, err := wordlistSupplier(languageIdentifier)
		if err != nil {
			return nil, fmt.Errorf("error invoking wordlistSupplier: %w", err)
		}

		// Wordlists are guaranteed not to contain any carriage returns (\r).
		words = strings.Split(lowercaser.String(wordListFile), "\n")
		wordListCache[languageIdentifier] = words
	}

	// We don't shuffle the wordList directory, as the cache isn't threadsafe.
	shuffledWords := make([]string, len(words))
	copy(shuffledWords, words)
	shuffleWordList(shuffledWords)
	return shuffledWords, nil
}

// readDefaultWordList reads the wordlist for the given language from the filesystem.
// If found, the list is cached and will be read from the cache upon next
// request. The returned slice is a safe copy and can be mutated. If the
// specified has no corresponding wordlist, an error is returned. This has been
// a panic before, however, this could enable a user to forcefully crash the
// whole application.
func readDefaultWordList(lowercaser cases.Caser, chosenLanguage string) ([]string, error) {
	return readWordListInternal(lowercaser, chosenLanguage, func(key string) (string, error) {
		wordBytes, err := wordFS.ReadFile("words/" + key)
		if err != nil {
			return "", fmt.Errorf("error reading wordfile: %w", err)
		}

		return string(wordBytes), nil
	})
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
