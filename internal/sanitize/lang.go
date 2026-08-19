package sanitize

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type LanguageData struct {
	Lowercaser   func() cases.Caser
	AlwaysVisibleCharacters string
	Transliterations map[rune]string
	LanguageCode string
	IsRtl        bool
}

var (
	AllLanguageData       = map[string]LanguageData{
		"custom": {
			LanguageCode: "en_gb",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.BritishEnglish) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"english_gb": {
			LanguageCode: "en_gb",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.BritishEnglish) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"english": {
			LanguageCode: "en_us",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.AmericanEnglish) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"italian": {
			LanguageCode: "it",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Italian) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"german": {
			LanguageCode: "de",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.German) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"french": {
			LanguageCode: "fr",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.French) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters + "'.",
		},
		"dutch": {
			LanguageCode: "nl",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Dutch) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"ukrainian": {
			LanguageCode: "ua",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Ukrainian) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"russian": {
			LanguageCode: "ru",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Russian) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"polish": {
			LanguageCode: "pl",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Polish) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"arabic": {
			IsRtl:        true,
			LanguageCode: "ar",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Arabic) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"hebrew": {
			IsRtl:        true,
			LanguageCode: "he",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Hebrew) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
		"persian": {
			IsRtl:        true,
			LanguageCode: "fa",
			Lowercaser:   func() cases.Caser { return cases.Lower(language.Persian) },
			Transliterations: DefaultTransliterations,
			AlwaysVisibleCharacters: DefaultAlwaysVisibleCharacters,
		},
	}
)