package translate

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/cubiest/jibberjabber"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/samber/lo"
	"github.com/snivilised/arcadia/src/internal/l10n"
	"golang.org/x/text/language"
)

var languages *LanguageInfo
var localiser *i18n.Localizer

type LanguageInitOptions struct {
	// App name
	//
	App string

	// Path denoting install location
	//
	Path string
}

// ValidatorContainerOptionFn definition of a client defined function to
// set ValidatorContainer options.
//
type LanguageInitOptionFn func(*LanguageInitOptions)

func Initialise(options ...LanguageInitOptionFn) {

	o := LanguageInitOptions{}
	for _, fo := range options {
		fo(&o)
	}

	languages = createInitialLanguageInfo(o)
	localiser = createLocaliser(languages)
}

// the active file should be in the same directory at the item that is
// loading the bundle
//
// Create the "active.en.json" file from internal/l18n:
// cd internal/l10n
// goi18n extract -format json

// do merge
// goi18n merge -outdir out -format json active.en.json translate.en-US.json
//
// rename out/active.en-US.json => out/arcadia.active.en-US.json
//
// nb: translate.en-US.json is an output file that does not need to be stored
// in source control
//

// LanguageInfo indicates information relating to current language. See members for
// details.
//
type LanguageInfo struct {
	// App name which forms part of the language filename
	//
	App string

	// Path denoting where to load language file from, defaults to exe location
	//
	Path string

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
// other than the one automatically detected.
//
func UseTag(tag language.Tag) error {
	_, found := lo.Find(languages.Supported, func(t language.Tag) bool {
		return t == tag
	})

	if found {
		languages = createIncrementalLanguageInfo(tag, languages)
		localiser = createLocaliser(languages)
	} else {
		return fmt.Errorf(GetLanguageNotSupportedErrorMessage(tag))
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

func detect() *detectInfo {
	detectedLang, _ := jibberjabber.DetectLanguage()
	territory, _ := jibberjabber.DetectTerritory()

	detectedLangTag, _ := language.Parse(fmt.Sprintf("%v-%v", detectedLang, territory))

	return &detectInfo{
		tag:       detectedLangTag,
		territory: territory,
	}
}

func createInitialLanguageInfo(options LanguageInitOptions) *LanguageInfo {
	dInfo := detect()

	return &LanguageInfo{
		App:       options.App,
		Path:      options.Path,
		Default:   language.BritishEnglish,
		Detected:  dInfo.tag,
		Territory: dInfo.territory,
		Current:   dInfo.tag,
		Supported: []language.Tag{language.BritishEnglish, language.AmericanEnglish},
	}
}

func createIncrementalLanguageInfo(requested language.Tag, existing *LanguageInfo) *LanguageInfo {

	return &LanguageInfo{
		App:       existing.App,
		Path:      existing.Path,
		Default:   language.BritishEnglish,
		Detected:  existing.Detected,
		Territory: existing.Territory,
		Current:   requested,
		Supported: []language.Tag{language.BritishEnglish, language.AmericanEnglish},
	}
}

func createLocaliser(li *LanguageInfo) *i18n.Localizer {
	bundle := i18n.NewBundle(li.Current)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	if li.Current != li.Default {
		filename := fmt.Sprintf("%v.active.%v.json", li.App, li.Current)

		exe, _ := os.Executable()
		resolved, _ := filepath.Abs(li.Path)
		directory := lo.Ternary(li.Path != "", resolved, filepath.Dir(exe))
		path := filepath.Join(directory, filename)

		_, err := bundle.LoadMessageFile(path)

		if err != nil {
			panic(fmt.Errorf(GetCouldNotLoadTranslationFileErrorMessage(li.Current, path)))
		}
	}

	supported := lo.Map(li.Supported, func(t language.Tag, _ int) string {
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
