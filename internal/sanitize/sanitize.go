// Package sanitize is used for cleaning up text.
package sanitize

import (
	"unicode/utf8"
)

var DefaultAlwaysVisibleCharacters = " _-,:?!/#+&()=<>\""

// FIXME Improve transliteration set or document why the current state
// is acceptable. These transliterations originally come from
// github.com/kennygrant/sanitize.
var DefaultTransliterations = map[rune]string{
	'À': "A",
	'Á': "A",
	'Â': "A",
	'Ã': "A",
	'Ä': "A",
	'Å': "AA",
	'Æ': "AE",
	'Ç': "C",
	'È': "E",
	'É': "E",
	'Ê': "E",
	'Ë': "E",
	'Ì': "I",
	'Í': "I",
	'Î': "I",
	'Ï': "I",
	'Ð': "D",
	'Ł': "L",
	'Ñ': "N",
	'Ò': "O",
	'Ó': "O",
	'Ô': "O",
	'Õ': "O",
	'Ö': "OE",
	'Ø': "OE",
	'Œ': "OE",
	'Ù': "U",
	'Ú': "U",
	'Ü': "UE",
	'Û': "U",
	'Ý': "Y",
	'Þ': "TH",
	'ẞ': "SS",
	'à': "a",
	'á': "a",
	'â': "a",
	'ã': "a",
	'ă': "a",
	'ä': "ae",
	'å': "aa",
	'æ': "ae",
	'ć': "c",
	'ç': "c",
	'è': "e",
	'é': "e",
	'ê': "e",
	'ë': "e",
	'ę': "e",
	'ì': "i",
	'í': "i",
	'î': "i",
	'ï': "i",
	'ð': "d",
	'ł': "l",
	'ñ': "n",
	'ń': "n",
	'ò': "o",
	'ó': "o",
	'ô': "o",
	'õ': "o",
	'ō': "o",
	'ö': "oe",
	'ø': "oe",
	'œ': "oe",
	'ś': "s",
	'š': "s",
	'ù': "u",
	'ú': "u",
	'û': "u",
	'ū': "u",
	'ü': "ue",
	'ý': "y",
	'ÿ': "y",
	'ż': "z",
	'ž': "z",
	'þ': "th",
	'ß': "ss",
	'²': "2",
	'¹': "1",
	'³': "3",
	'⁴': "4",
	'⁵': "5",
	'⁶': "6",
	'⁷': "7",
	'⁸': "8",
	'⁹': "9",
	'’': "'",
	'—': "-",
	'–': "-",
	' ': " ",
	'′': "'",
}

// CleanText removes all kinds of characters that could disturb the algorithm
// checking words for similarity.
func (lang LanguageData) CleanText(str string) string {

	str = lang.Lowercaser().String(str)

	var buffer []byte

	// We try to stack allocate, but also make
	// space for the worst case scenario.
	if len(str) <= 32 {
		buffer = make([]byte, 0, 64)
	} else {
		buffer = make([]byte, 0, len(str)*2)
	}

	var changed bool
	for _, character := range str {
		if character < utf8.RuneSelf {
			if lang.IsAlwaysVisibleCharacter(character) {
				changed = true
			} else {
				buffer = append(buffer, byte(character))
			}
			continue
		}

		if val, contains := lang.Transliterations[character]; contains {
			buffer = append(buffer, val...)
			changed = true
		} else {
			buffer = utf8.AppendRune(buffer, character)
		}
	}

	if !changed {
		return str
	}
	return string(buffer)
}

func (lang LanguageData) IsAlwaysVisibleCharacter(c rune) bool {
	if val, contains := lang.Transliterations[c]; contains {
		for _, cc := range val {
			if (! lang.IsAlwaysVisibleCharacter(cc)) {
				return false
			}
		}
		return true
	}
	for _, x := range lang.AlwaysVisibleCharacters {
		if x == c {
			return true
		}
	}
	return false
}
