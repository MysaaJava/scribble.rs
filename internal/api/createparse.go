package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/scribble-rs/scribble.rs/internal/config"
	"github.com/scribble-rs/scribble.rs/internal/game"
	"github.com/scribble-rs/scribble.rs/internal/sanitize"
	"golang.org/x/text/cases"
)

// ParsePlayerName checks if the given value is a valid playername. Currently
// this only includes checkin whether the value is empty or only consists of
// whitespace character.
func ParsePlayerName(value string) (string, error) {
	trimmed := strings.TrimSpace(value)
	if trimmed == "" {
		return trimmed, errors.New("the player name must not be empty")
	}

	return trimmed, nil
}

// GetWordlist returns the Wordlist corresponding to the specified
// name from the input list of wordlists
// A name can contain slash, in which case, the wordlist will be searched
// in the children of the wordlist (like a system path)
func GetWordList(lists []*game.WordList, name string) (*game.WordList, error) {
	if (name == "") {
		return nil, errors.New("The empty string is not a valid wordlist name")
	}
	part0, rest, multipart := strings.Cut(name, "/")
	i := -1
	for index, wl := range lists {
		if wl.Name == part0 {
			i = index
		}
	}
	if i == -1 {
		return nil, errors.New("Wordlist name not found")
	}
	wl0 := lists[i]
	if multipart {
		return GetWordList(wl0.Children, rest)
	}
	return wl0, nil
}

// ParseWordLists checks whether the given value is a string containing comma
// separated wordlists names. It will also try to check if every specified wordlist is known to the system.
// double underscore are replaced with slashes.
// If the input string is empty and language is not empty, this will return the default list for the 
// language (i.e. list named `scribblers/$lang`)
func ParseWordLists(cfg *config.Config, languageKey string, values []string) ([]*game.WordList, error) {
	if (len(values)==0) {
		if languageKey == "" {
			return nil, errors.New("You must select at least one word group or select a valid language")
		}
		values = []string{"scribblers/" + languageKey}
	}

	allLists := cfg.AllWordLists()

	count := len(values)
	result := make([]*game.WordList,count)
	for index, item := range values {
		trimmedItem := strings.TrimSpace(item)
		replacedItem := strings.Replace(trimmedItem, "__", "/", -1)
		wl,error := GetWordList(allLists, replacedItem)
		if error != nil {
			return nil, fmt.Errorf("Could not find word group %s", replacedItem)
		}
		result[index] = wl
	}
	return result, nil
}

// ParseLanguage checks whether the given value is part of the
// game.SupportedLanguages array. The input is trimmed and lowercased.
func ParseLanguage(value string) (*sanitize.LanguageData, string, error) {
	toLower := strings.ToLower(strings.TrimSpace(value))
	for languageKey, data := range sanitize.AllLanguageData {
		if toLower == languageKey {
			return &data, languageKey, nil
		}
	}

	return nil, "", errors.New("the given language doesn't match any supported language")
}

func ParseScoreCalculation(value string) (game.ScoreCalculation, error) {
	toLower := strings.ToLower(strings.TrimSpace(value))
	switch toLower {
	case "", "chill":
		return game.ChillScoring, nil
	case "competitive":
		return game.CompetitiveScoring, nil
	}

	return nil, errors.New("the given score calculation doesn't match any supported algorithm")
}

// ParseDrawingTime checks whether the given value is an integer between
// the lower and upper bound of drawing time. All other invalid
// input, including empty strings, will return an error.
func ParseDrawingTime(cfg *config.Config, value string) (int, error) {
	return parseIntValue(value, cfg.LobbySettingBounds.MinDrawingTime,
		cfg.LobbySettingBounds.MaxDrawingTime, "drawing time")
}

// ParseRounds checks whether the given value is an integer between
// the lower and upper bound of rounds played. All other invalid
// input, including empty strings, will return an error.
func ParseRounds(cfg *config.Config, value string) (int, error) {
	return parseIntValue(value, cfg.LobbySettingBounds.MinRounds,
		cfg.LobbySettingBounds.MaxRounds, "rounds")
}

// ParseMaxPlayers checks whether the given value is an integer between
// the lower and upper bound of maximum players per lobby. All other invalid
// input, including empty strings, will return an error.
func ParseMaxPlayers(cfg *config.Config, value string) (int, error) {
	return parseIntValue(value, cfg.LobbySettingBounds.MinMaxPlayers,
		cfg.LobbySettingBounds.MaxMaxPlayers, "max players amount")
}

// ParseCustomWords checks whether the given value is a string containing comma
// separated values (or a single word). Empty strings will return an empty
// (nil) array and no error. An error is only returned if there are empty words.
// For example these wouldn't parse:
//
//	wordone,,wordtwo
//	,
//	wordone,
func ParseCustomWords(lowercaser cases.Caser, value string) ([]string, error) {
	trimmedValue := strings.TrimSpace(value)
	if trimmedValue == "" {
		return nil, nil
	}

	result := strings.Split(trimmedValue, ",")
	for index, item := range result {
		trimmedItem := lowercaser.String(strings.TrimSpace(item))
		if trimmedItem == "" {
			return nil, errors.New("custom words must not be empty")
		}
		result[index] = trimmedItem
	}

	return result, nil
}

// ParseClientsPerIPLimit checks whether the given value is an integer between
// the lower and upper bound of maximum clients per IP. All other invalid
// input, including empty strings, will return an error.
func ParseClientsPerIPLimit(cfg *config.Config, value string) (int, error) {
	return parseIntValue(value, cfg.LobbySettingBounds.MinClientsPerIPLimit,
		cfg.LobbySettingBounds.MaxClientsPerIPLimit, "clients per IP limit")
}

// ParseCustomWordsPerTurn checks whether the given value is an integer between
// 0 and 100. All other invalid input, including empty strings, will return an
// error.
func ParseCustomWordsPerTurn(cfg *config.Config, value string) (int, error) {
	return parseIntValue(value, 1, cfg.LobbySettingBounds.MaxWordsPerTurn, "custom words per turn")
}

func ParseWordsPerTurn(cfg *config.Config, value string) (int, error) {
	return parseIntValue(value, 1, cfg.LobbySettingBounds.MaxWordsPerTurn, "words per turn")
}

func newIntOutOfBounds(value, valueName string, lower, upper int) error {
	if upper != -1 {
		return fmt.Errorf("the value '%s' must be an integer between %d and %d, but was: '%s'", valueName, lower, upper, value)
	}
	return fmt.Errorf("the value '%s' must be an integer larger than %d, but was: '%s'", valueName, lower, value)
}

func parseIntValue(toParse string, lower, upper int, valueName string) (int, error) {
	var value int
	if parsed, err := strconv.ParseInt(toParse, 10, 64); err != nil {
		return 0, newIntOutOfBounds(toParse, valueName, lower, upper)
	} else {
		value = int(parsed)
	}

	if value < lower || (upper > -1 && value > upper) {
		return 0, newIntOutOfBounds(toParse, valueName, lower, upper)
	}

	return value, nil
}

// ParseBoolean checks whether the given value is either "true" or "false".
// The checks are case-insensitive. If an empty string is supplied, false
// is returned. All other invalid input will return an error.
func ParseBoolean(valueName, value string) (bool, error) {
	if value == "" {
		return false, nil
	}

	if strings.EqualFold(value, "true") {
		return true, nil
	}

	if strings.EqualFold(value, "false") {
		return false, nil
	}

	return false, fmt.Errorf("the %s value must be a boolean value ('true' or 'false)", valueName)
}
