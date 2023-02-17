package l10n

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Localisable interface {
	Message() *i18n.Message
}

// language not supported.
type LanguageNotSupportedTemplData struct {
	Language string
}

func (td LanguageNotSupportedTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "language-not-supported.arcadia",
		Description: "The language specified is not supported; no translations for this language.",
		Other:       "language '{{.Language}}' not supported",
	}
}

// could not load translation file.
type CouldNotLoadTranslationFileTemplData struct {
	Language string
	FullPath string
}

func (td CouldNotLoadTranslationFileTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "could-not-load-translation-file.arcadia",
		Description: "Could not load translation file for this language.",
		Other:       "could not load translation file for language '{{.Language}}' ({{.FullPath}})",
	}
}
