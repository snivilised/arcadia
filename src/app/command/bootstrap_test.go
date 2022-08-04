package command_test

import (
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"golang.org/x/text/language"

	"github.com/snivilised/arcadia/src/app/command"
	"github.com/snivilised/arcadia/src/internal/translate"
)

type USFake struct {
}

func (j *USFake) Scan() language.Tag {
	return language.AmericanEnglish
}

var _ = Describe("Bootstrap", func() {

	Context("widget command", func() {
		It("🧪 should: invoke without error", func() {
			directory, _ := filepath.Abs("../../internal/l10n/out")

			bootstrap := command.Bootstrap{
				Detector: &USFake{},
			}
			bootstrap.Execute(func(detector command.LocaleDetector) []string {
				args := []string{"widget", "-p", "P?<date>", "-t", "30"}
				translate.Initialise(func(o *translate.LanguageInitOptions) {
					o.Detected = detector.Scan()
					o.App = "arcadia"
					o.Path = directory
				})

				return args
			})
			Expect(true)
		})
	})
})
