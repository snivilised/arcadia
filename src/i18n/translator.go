package i18n

import (
	"github.com/samber/lo"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/snivilised/extendio/xfs/utils"
	"golang.org/x/text/language"
)

var (
	// DefaultLanguage, update this with the required default
	DefaultLanguage = utils.NewRoProp(language.BritishEnglish)
	tx              xi18n.Translator
	TxRef           utils.RoProp[xi18n.Translator] = utils.NewRoProp(tx)
)

// containsLanguage should be exported by extendio, so it can be reused
// instead of re-defining it here
func containsLanguage(languages xi18n.SupportedLanguages, tag language.Tag) bool {
	return lo.ContainsBy(languages, func(t language.Tag) bool {
		return t == tag
	})
}

// Text is the function to use to obtain a string created from
// registered Localizers. The data parameter must be a go template
// defining the input parameters and the translatable message content.
func Text(data xi18n.Localisable) string {
	return tx.Localise(data)
}

// Use, must be called by the client before any string data
// can be translated. If the client requests the default
// language, then only the language Tag needs to be provided.
// If the requested language is not the default and therefore
// requires translation from the translation file(s), then
// the client must provide the App and Path properties indicating
// how the l18n bundle is created.
func Use(options ...xi18n.UseOptionFn) error {

	o := &xi18n.UseOptions{}
	for _, fo := range options {
		fo(o)
	}

	// we let extendio manage translations for itself ...
	//
	err := xi18n.Use(func(uo *xi18n.UseOptions) {
		uo.Tag = o.Tag
		uo.From = o.From
	})

	if err != nil {
		return err
	}

	// The root command has specified the translation requirements
	// so we just use what has been requested.
	//
	lang := xi18n.NewLanguageInfo(o)
	if !containsLanguage(lang.Supported, o.Tag) {
		return xi18n.NewLanguageNotAvailableNativeError(o.Tag)
	}

	factory := lo.TernaryF(len(o.From.Sources) > 1,
		func() xi18n.TranslatorFactory {
			return &xi18n.MultiTranslatorFactory{}
		},
		func() xi18n.TranslatorFactory {
			return &xi18n.SingularTranslatorFactory{}
		},
	)

	tx = factory.New(lang)
	TxRef = utils.NewRoProp(tx)

	if TxRef.IsNone() {
		return xi18n.NewFailedToCreateTranslatorNativeError(o.Tag)
	}

	return nil
}

// UseTx, do not use, required for unit testing only and is not considered
// part of the public api and may be removed without corresponding version
// number change.
func UseTx(with xi18n.Translator, setters ...xi18n.UseOptionFn) error {
	o := &xi18n.UseOptions{}
	for _, fo := range setters {
		fo(o)
	}

	tx = with
	TxRef = utils.NewRoProp(tx)

	if TxRef.IsNone() {
		return xi18n.NewFailedToCreateTranslatorNativeError(o.Tag)
	}

	return nil
}