package i18n

import (
	xi18n "github.com/snivilised/extendio/i18n"
)

var (
	// DefaultLanguage, update this with the required default
	DefaultLanguage = xi18n.DefaultLanguage
)

// Text is the function to use to obtain a string created from
// registered Localizers. The data parameter must be a go template
// defining the input parameters and the translatable message content.
func Text(data xi18n.Localisable) string {
	return xi18n.Text(data)
}

// Use, must be called by the client before any string data
// can be translated. If the client requests the default
// language, then only the language Tag needs to be provided.
// If the requested language is not the default and therefore
// requires translation from the translation file(s), then
// the client must provide the App and Path properties indicating
// how the l18n bundle is created.
func Use(options ...xi18n.UseOptionFn) error {
	return xi18n.Use(options...)
}
