package main_test

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/arcadia/src/app/main"
	"github.com/snivilised/arcadia/src/internal/l10n"
)

var _ = Describe("i18n", func() {

	// var languages *main.LanguageInfo

	Context("go-i18n", func() {
		When("using template", func() {
			It("🧪 should: translate", func() {
				notSupportedMsg := &i18n.Message{
					ID:    "language-not-supported.arcadia",
					Other: "language '{{.Language}}' not supported",
				}

				localised := main.GetLocaliser().MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: notSupportedMsg,
					TemplateData: l10n.LanguageNotSupportedTemplData{
						Language: "es",
					},
				})
				expected := "language 'es' not supported"
				Expect(localised).To(Equal(expected))
			})
		})
	})
})
