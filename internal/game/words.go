package game

import (
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"strings"
	"unicode/utf8"
	"slices"
	"bufio"
	"os"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type LanguageData struct {
	Lowercaser   func() cases.Caser
	LanguageCode string
	IsRtl        bool
}

var (
	ErrUnknownWordList = errors.New("wordlist unknown")
	WordlistData       = map[string]LanguageData{
		"custom": {
			LanguageCode: "en_gb",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.BritishEnglish) },
		},
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
		"russian": {
			LanguageCode: "ru",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Russian) },
		},
		"polish": {
			LanguageCode: "pl",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Polish) },
		},
		"arabic": {
			IsRtl:        true,
			LanguageCode: "ar",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Arabic) },
		},
		"hebrew": {
			IsRtl:        true,
			LanguageCode: "he",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Hebrew) },
		},
		"persian": {
			IsRtl:        true,
			LanguageCode: "fa",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Persian) },
		},
	}
)

func getLanguageIdentifier(language string) string {
	return WordlistData[language].LanguageCode
}

// readWordListInternal exists for testing purposes, it allows passing a custom
// wordListSupplier, in order to avoid having to write tests aggainst the
// default language lists.
func readWordListInternal(
	lowercaser cases.Caser, chosenLanguage string,
	wordlistSupplier func(string) (string, error),
) ([]string, error) {
	languageIdentifier := getLanguageIdentifier(chosenLanguage)
	if languageIdentifier == "" {
		return nil, ErrUnknownWordList
	}

	wordListFile, err := wordlistSupplier(languageIdentifier)
	if err != nil {
		return nil, fmt.Errorf("error invoking wordlistSupplier: %w", err)
	}

	// Wordlists are guaranteed not to contain any carriage returns (\r).
	words := strings.Split(lowercaser.String(wordListFile), "\n")
	shuffleWordList(words)
	return words, nil
}

// readFileWordLists reads the input path and appends all read words to the input array pointer.
func readFileWordList(path string, lines *[]string) error {
	file, err := os.Open(path)
    if err != nil {
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
    }
    return scanner.Err()
}

// readWordList reads a wordlist (and its children recursively) and add all read words to the input array pointer
func readWordList(wl *WordList, data *[]string) error {
	if (wl.Path != "") {
		return readFileWordList(wl.Path, data)
	} else {
		for _, ch := range wl.Children {
			err := readWordList(ch, data)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

// reloadlobbyWords outputs the list of words for the given lobby, reading required wordlists files
func reloadLobbyWords(lobby *Lobby) ([]string, error) {
	wls := lobby.WordLists
	log.Printf("Reading %d word lists", len(wls))
	var words []string = make([]string, 0)
	wordsp := &words
	for _, wl := range wls {
		err := readWordList(wl, wordsp)
		if err != nil {
			return nil, err
		}
	}
	log.Printf("Read a total of %d words", len(*wordsp))
	return *wordsp, nil
}

// Helper function that takes out `count` random items out af an array `arr`
func RandNUniqueOfSlice(count int, arr []string) []string {
	out := make([]string,count)
	for i := 0; i < count; i++ {
		var candidate string
		for k := 0; k < count * count * 100 && (slices.Contains(out,candidate)); k++ {
			candidate = arr[rand.IntN(len(arr))]
		}
		out[i] = candidate
	}
	return out
}

// GetRandomWords gets a custom amount of random words for the passed Lobby.
// The words will be chosen from the custom words and the default
// dictionary, depending on the settings specified by the lobbies creator.
func GetRandomWords(wordCount int, lobby *Lobby) []string {
	return getRandomWords(wordCount, lobby, reloadLobbyWords)
}

// getRandomWords exists for test purposes, allowing to define a custom
// reloader, allowing us to specify custom wordlists in the tests without
// running into a panic on reload.
func getRandomWords(wordCount int, lobby *Lobby, reloadWords func(lobby *Lobby) ([]string, error)) []string {
	words := make([]string, wordCount)

	// If we have custom words only, we don't want to pop them off the stack.
	// We want to keep going in circles, worstcase returning the same word 3 times.
	if len(lobby.WordLists) == 0 && len(lobby.CustomWords) > 0 {
		for i := range wordCount {
			if lobby.customWordIndex >= len(lobby.CustomWords) {
				lobby.customWordIndex = 0
			}
			words[i] = lobby.CustomWords[lobby.customWordIndex]
			lobby.customWordIndex++
		}
		return words
	}

	for customWordsLeft, i := lobby.CustomWordsPerTurn, 0; i < wordCount; i++ {
		if customWordsLeft > 0 && len(lobby.CustomWords) > 0 {
			customWordsLeft--
			words[i] = popCustomWord(lobby)
		} else {
			words[i] = popWordpackWord(lobby, reloadWords)
		}
	}

	return words
}

func popCustomWord(lobby *Lobby) string {
	lastIndex := len(lobby.CustomWords) - 1
	lastWord := lobby.CustomWords[lastIndex]
	lobby.CustomWords = lobby.CustomWords[:lastIndex]
	return lastWord
}

// popWordpackWord gets X words from the wordpack. The major difference to
// popCustomWords is, that the wordlist gets reset and reshuffeled once every
// item has been popped.
func popWordpackWord(lobby *Lobby, reloadWords func(lobby *Lobby) ([]string, error)) string {
	if len(lobby.words) == 0 {
		var err error
		lobby.words, err = reloadWords(lobby)
		if err != nil {
			// Since this list should've been successfully read once before, we
			// can "safely" panic if this happens, assuming that there's a
			// deeper problem.
			panic(err)
		}
	}
	lastIndex := len(lobby.words) - 1
	lastWord := lobby.words[lastIndex]
	lobby.words = lobby.words[:lastIndex]
	return lastWord
}

func shuffleWordList(wordlist []string) {
	rand.Shuffle(len(wordlist), func(a, b int) {
		wordlist[a], wordlist[b] = wordlist[b], wordlist[a]
	})
}

const (
	EqualGuess   = 0
	CloseGuess   = 1
	DistantGuess = 2
)

// CheckGuess compares the strings with eachother. Possible results:
//   - EqualGuess (0)
//   - CloseGuess (1)
//   - DistantGuess (2)
//
// This works mostly like levensthein distance, but doesn't check further than
// to a distance of 2 and also handles transpositions where the runes are
// directly next to eachother.
func CheckGuess(a, b string) int {
	// We only want to indicate a close guess if:
	//   * 1 additional character is found (abc ~ abcd)
	//   * 1 character is missing (abc ~ ab)
	//   * 1 character is wrong (abc ~ adc)
	//   * 2 characters are swapped (abc ~ acb)

	// If the longer string can't be on both sides, the follow-up logic can
	// be simpler, so we switch them here.
	if len(a) < len(b) {
		a, b = b, a
	}

	// A maximum of 4 bytes is used for one unicode code point. So if there is a definitive character
	// count difference of 2 or more characters, we can clearly indicate a distant guess.\
	// This prevents having to count all runes in b.
	if len(a)-len(b) >= 8 {
		return DistantGuess
	}

	if a == b {
		return EqualGuess
	}

	var distance int
	aBytes := []byte(a)
	bBytes := []byte(b)
	for {
		aRune, aSize := utf8.DecodeRune(aBytes)
		// If a eaches the end, then so does b, as we make sure a is longer at
		// the top, therefore we can be sure no additional conflict diff occurs.
		if aRune == utf8.RuneError {
			// If a is longer in terms of bytes, but contains for example an emoji that takes up 4 bytes, this CAN happen.
			distance += utf8.RuneCount(bBytes)
			if distance >= 2 {
				return DistantGuess
			}
			return distance
		}
		bRune, bSize := utf8.DecodeRune(bBytes)

		// Either different runes, or b is empty, returning RuneError (65533).
		if aRune != bRune {
			// Check for transposition (abc ~ acb)
			nextARune, nextASize := utf8.DecodeRune(aBytes[aSize:])
			if nextARune == bRune {
				if nextARune != utf8.RuneError {
					nextBRune, nextBSize := utf8.DecodeRune(bBytes[bSize:])
					if nextBRune == aRune {
						distance++
						aBytes = aBytes[aSize+nextASize:]
						bBytes = bBytes[bSize+nextBSize:]
						continue
					}
				}

				// Make sure to not pop from b, so we can compare the rest, in
				// case we are only missing one character for cases such as:
				//   abc ~ bc
				//   abcde ~ abde
				bSize = 0
			} else if distance == 1 {
				// We'd reach a diff of 2 now. Needs to happen after transposition
				// though, as transposition could still prove us wrong.
				return DistantGuess
			}

			distance++
		}

		aBytes = aBytes[aSize:]
		bBytes = bBytes[bSize:]
	}
}
