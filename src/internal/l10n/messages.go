package l10n

import "github.com/nicksnyder/go-i18n/v2/i18n"

type Localisable interface {
	Message() *i18n.Message
}

type PsDummyTemplData struct {
	Name string
}

func (td PsDummyTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "dummy.arcadia",
		Description: "THIS IS A DUMMY",
		Other:       "Dummy ('{{.Name}}')",
	}
}
