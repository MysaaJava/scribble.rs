package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/scribble-rs/scribble.rs/internal/game"
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

// ParseLanguage checks whether the given value is part of the
// game.SupportedLanguages array. The input is trimmed and lowercased.
func ParseLanguage(value string) (string, error) {
	toLower := strings.ToLower(strings.TrimSpace(value))
	for languageKey := range game.SupportedLanguages {
		if toLower == languageKey {
			return languageKey, nil
		}
	}

	return "", errors.New("the given language doesn't match any supported language")
}

// ParseDrawingTime checks whether the given value is an integer between
// the lower and upper bound of drawing time. All other invalid
// input, including empty strings, will return an error.
func ParseDrawingTime(value string) (int, error) {
	return parseIntValue(value, game.LobbySettingBounds.MinDrawingTime,
		game.LobbySettingBounds.MaxDrawingTime, "drawing time")
}

// ParseRounds checks whether the given value is an integer between
// the lower and upper bound of rounds played. All other invalid
// input, including empty strings, will return an error.
func ParseRounds(value string) (int, error) {
	return parseIntValue(value, game.LobbySettingBounds.MinRounds,
		game.LobbySettingBounds.MaxRounds, "rounds")
}

// ParseMaxPlayers checks whether the given value is an integer between
// the lower and upper bound of maximum players per lobby. All other invalid
// input, including empty strings, will return an error.
func ParseMaxPlayers(value string) (int, error) {
	return parseIntValue(value, game.LobbySettingBounds.MinMaxPlayers,
		game.LobbySettingBounds.MaxMaxPlayers, "max players amount")
}

// ParseCustomWords checks whether the given value is a string containing comma
// separated values (or a single word). Empty strings will return an empty
// (nil) array and no error. An error is only returned if there are empty words.
// For example these wouldn't parse:
//
//	wordone,,wordtwo
//	,
//	wordone,
func ParseWordGroups(values []string) ([]int, error) {
	if (len(values)==0) {
		return nil, errors.New("You must select at least one word group")
	}
	result := make([]int,len(values))
	for index, item := range values {
		parsed, errResult := strconv.ParseInt(item, 10, 32)
		if (errResult != nil) {
			return nil,errResult
		}
		result[index] = int(parsed)
	}

	return result, nil
}

// ParseClientsPerIPLimit checks whether the given value is an integer between
// the lower and upper bound of maximum clients per IP. All other invalid
// input, including empty strings, will return an error.
func ParseClientsPerIPLimit(value string) (int, error) {
	return parseIntValue(value, game.LobbySettingBounds.MinClientsPerIPLimit,
		game.LobbySettingBounds.MaxClientsPerIPLimit, "clients per IP limit")
}

// ParseCustomWordsPerTurn checks whether the given value is an integer between
// 0 and 100. All other invalid input, including empty strings, will return an
// error.
func ParseCustomWordsPerTurn(value string) (int, error) {
	return parseIntValue(value, 1, 3, "custom words per turn")
}

func newIntOutOfBounds(value, valueName string, lower, upper int) error {
	return fmt.Errorf("the value '%s' must be an integer between %d and %d, but was: '%s'", valueName, lower, upper, value)
}

func parseIntValue(toParse string, lower, upper int, valueName string) (int, error) {
	var value int
	if parsed, err := strconv.ParseInt(toParse, 10, 64); err != nil {
		return 0, newIntOutOfBounds(toParse, valueName, lower, upper)
	} else {
		value = int(parsed)
	}

	if value < lower || value > upper {
		return 0, newIntOutOfBounds(toParse, valueName, lower, upper)
	}

	return value, nil
}

// ParseBoolean checks whether the given value is either "true" or "false".
// The checks are case-insensitive. If an empty string is supplied, false
// is returned. All other invalid input will return an error.
func ParseBoolean(valueName string, value string) (bool, error) {
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
