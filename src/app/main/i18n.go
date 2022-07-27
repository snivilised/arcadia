package main

import (
	"encoding/json"
	"fmt"

	"github.com/cubiest/jibberjabber"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/samber/lo"
	"github.com/snivilised/arcadia/src/internal/l10n"
	"golang.org/x/text/language"
)

// the active file should be in the same directory at the item that is
// loading the bundle
//
// Create the "active.en.json" file from internal/l18n:
// cd internal/l18n
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

// UseTag allows the client to change the language currently in use to a language
// othr than the one automatically detected.
//
func UseTag(tag language.Tag) error {
	_, found := lo.Find(languages.Supported, func(t language.Tag) bool {
		return t == tag
	})

	if found {
		languages = createIncrementalLanguageInfo(tag, languages)
		localiser = createLocaliser(languages)
	} else {
		return fmt.Errorf(getLanguageNotSupportedErrorMessage(tag))
	}

	return nil
}

// GetLanguageInfo gets LanguageInfo
//
func GetLanguageInfo() *LanguageInfo {
	return languages
}

// GetLocaliser gets the current go-i18n localizer instance
//
func GetLocaliser() *i18n.Localizer {
	return localiser
}

type detectInfo struct {
	tag       language.Tag
	territory string
}

var languages *LanguageInfo
var localiser *i18n.Localizer

func init() {
	languages = createInitialLanguageInfo()
	localiser = createLocaliser(languages)
}

func detect() *detectInfo {
	detectedLang, _ := jibberjabber.DetectLanguage()
	territory, _ := jibberjabber.DetectTerritory()

	detectedLangTag, _ := language.Parse(fmt.Sprintf("%v-%v", detectedLang, territory))

	return &detectInfo{
		tag:       detectedLangTag,
		territory: territory,
	}
}

func createInitialLanguageInfo() *LanguageInfo {
	dInfo := detect()

	return &LanguageInfo{
		Default:   language.BritishEnglish,
		Detected:  dInfo.tag,
		Territory: dInfo.territory,
		Current:   dInfo.tag,
		Supported: []language.Tag{language.BritishEnglish, language.AmericanEnglish},
	}
}

func createIncrementalLanguageInfo(requested language.Tag, existing *LanguageInfo) *LanguageInfo {

	return &LanguageInfo{
		Default:   language.BritishEnglish,
		Detected:  existing.Detected,
		Territory: existing.Territory,
		Current:   requested,
		Supported: []language.Tag{language.BritishEnglish, language.AmericanEnglish},
	}
}

func createLocaliser(li *LanguageInfo) *i18n.Localizer {
	bundle := i18n.NewBundle(languages.Current)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	bundle.MustLoadMessageFile("../../internal/l10n/out/translate.en-US.json")

	supported := lo.Map(languages.Supported, func(t language.Tag, _ int) string {
		return t.String()
	})

	return i18n.NewLocalizer(bundle, supported...)
}

func localise(data l10n.Localisable) string {
	return localiser.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: data.Message(),
		TemplateData:   data,
	})
}

func getLanguageNotSupportedErrorMessage(tag language.Tag) string {
	data := l10n.LanguageNotSupportedTemplData{
		Language: tag.String(),
	}
	return localise(data)
}
