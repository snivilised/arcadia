package l10n

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Localisable interface {
	Message() *i18n.Message
}

// language not supported
//
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
