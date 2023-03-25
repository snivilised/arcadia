package command_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snivilised/arcadia/src/app/command"
	"github.com/snivilised/arcadia/src/internal/helpers"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/snivilised/extendio/xfs/utils"

	"golang.org/x/text/language"
)

type USFake struct {
}

func (j *USFake) Scan() language.Tag {
	return language.AmericanEnglish
}

var _ = Describe("Bootstrap", Ordered, func() {

	var (
		repo     string
		l10nPath string
	)

	BeforeAll(func() {
		repo = helpers.Repo("../..")
		l10nPath = helpers.Path(repo, "test/data/l10n")
		Expect(utils.FolderExists(l10nPath)).To(BeTrue())
	})

	Context("widget command", func() {
		It("🧪 should: invoke without error", func() {
			bootstrap := command.Bootstrap{
				Detector: &USFake{},
			}
			bootstrap.Execute(func(detector command.LocaleDetector) []string {

				from := xi18n.LoadFrom{
					Path: l10nPath,
					Sources: xi18n.TranslationFiles{
						command.SOURCE_ID: xi18n.TranslationSource{
							Name: fmt.Sprintf("test.%v", command.ApplicationName)},
					},
				}

				err := xi18n.Use(func(uo *xi18n.UseOptions) {
					uo.Tag = detector.Scan()
					uo.From = from
				})
				if err != nil {
					Fail(err.Error())
				}

				args := []string{"widget", "-p", "P?<date>", "-t", "30"}
				return args
			})
			Expect(true)
		})
	})
})
