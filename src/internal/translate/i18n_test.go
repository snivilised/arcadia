package translate_test

import (
	"fmt"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snivilised/arcadia/src/app/command"
	"github.com/snivilised/arcadia/src/internal/l10n"
	"github.com/snivilised/arcadia/src/internal/translate"
	"golang.org/x/text/language"
)

var _ = Describe("i18n", Ordered, func() {

	var es, us, expected string

	BeforeAll(func() {
		es = language.Spanish.String()
		us = language.AmericanEnglish.String()
		expected = fmt.Sprintf("language '%v' not supported", es)
		translate.Initialise(func(o *translate.LanguageInitOptions) {
			o.App = command.ApplicationName
			o.Path = "../l10n/out/"
		})
	})

	Context("UseTag", func() {
		When("given: tag is supported", func() {
			It("🧪 should: not return error", func() {
				us := language.AmericanEnglish
				Expect(translate.UseTag(us)).Error().To(BeNil())
				Expect(translate.GetLanguageInfo().Current).To(Equal(us))
			})

			It("🧪 should: localise in requested non default language", func() {
				_ = translate.UseTag(language.AmericanEnglish)
				data := l10n.LanguageNotSupportedTemplData{
					Language: us,
				}

				_, tag, _ := translate.GetLocaliser().LocalizeWithTag(&i18n.LocalizeConfig{
					DefaultMessage: data.Message(),
					TemplateData:   data,
				})
				Expect(tag.String()).To(Equal(language.AmericanEnglish.String()))
			})
		})

		When("given: tag is NOT supported", func() {
			It("🧪 should: return error", func() {
				Expect(translate.UseTag(language.Spanish)).Error().ToNot(BeNil())
			})
		})
	})

	Context("go-i18n", func() {
		When("using map of any", func() {
			It("🧪 should: translate", func() {
				notSupportedMsg := &i18n.Message{
					ID:    "language-not-supported.arcadia-lib",
					Other: "language '{{.Language}}' not supported",
				}

				localised := translate.GetLocaliser().MustLocalize(&i18n.LocalizeConfig{
					DefaultMessage: notSupportedMsg,
					TemplateData:   map[string]any{"Language": es},
				})

				Expect(localised).To(Equal(expected))
			})
		})

		When("using template", func() {
			It("🧪 should: translate", func() {
				localised := translate.GetLanguageNotSupportedErrorMessage(language.Spanish)

				Expect(localised).To(Equal(expected))
			})
		})
	})
})
