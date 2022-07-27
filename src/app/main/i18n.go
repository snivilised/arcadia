package main

import "golang.org/x/text/language"

// the active file should be in the same directory at the item that is
// loading the bundle
//
// Create the "active.en.json" file from internal/i18n:
// cd internal/i18n
// goi18n extract -format json

// do merge
// goi18n merge -outdir out -format json active.en.json translate.en-US.json
//

// LanguageInfo indicates information relating to current language. See members for
// details.
//
type LanguageInfo struct {
	// Default language reflects the base language. If all else fails, messages will
	// be in this language. It is fixed at BritishEnglish reflecting the language this
	// package is written in.
	//
	Default language.Tag

	// Detected is the language that is automatically detected of the host machine. Assuming
	// the ost machine is configured of the user's preference, there should be no other
	// reason to divert from this language.
	//
	Detected language.Tag

	// Territory reflects the region as automatically detected.
	//
	Territory string

	// Current reflects the language currently in force. Will by default be the detected
	// language. Client can change this with the UseTag function.
	//
	Current language.Tag

	// Supported indicates the list of languages for which translations are available.
	//
	Supported []language.Tag
}
