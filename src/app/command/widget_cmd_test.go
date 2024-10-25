package command_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2" //nolint:revive // ok for testing
	. "github.com/onsi/gomega"    //nolint:revive // ok for testing
	"github.com/spf13/cobra"

	"github.com/snivilised/arcadia/src/app/command"
	"github.com/snivilised/arcadia/src/internal/helpers"
	"github.com/snivilised/arcadia/src/locale"
	"github.com/snivilised/li18ngo"
	nef "github.com/snivilised/nefilim"

	ci18n "github.com/snivilised/cobrass/src/assistant/locale"
)

var _ = Describe("WidgetCmd", Ordered, func() {
	var (
		repo        string
		l10nPath    string
		bootstrap   command.Bootstrap
		rootCommand *cobra.Command
	)

	BeforeAll(func() {
		repo = helpers.Repo("")
		l10nPath = helpers.Path(repo, "test/data/l10n")
		fS := nef.NewUniversalABS()
		Expect(fS.DirectoryExists(l10nPath)).To(BeTrue(),
			fmt.Sprintf("💥 l10Path: '%v' does not exist", l10nPath),
		)
	})

	BeforeEach(func() {
		err := li18ngo.Use(func(uo *li18ngo.UseOptions) {
			uo.From = li18ngo.LoadFrom{
				Path: l10nPath,
				Sources: li18ngo.TranslationFiles{
					locale.ArcadiaSourceID: li18ngo.TranslationSource{
						Name: "pixa",
					},

					ci18n.CobrassSourceID: li18ngo.TranslationSource{
						Name: "cobrass",
					},
				},
			}
		})

		if err != nil {
			Fail(err.Error())
		}
		bootstrap = command.Bootstrap{}
		rootCommand = bootstrap.Root(func(co *command.ConfigureOptions) {
			co.Detector = &DetectorStub{}
			co.Config.Name = configName
			co.Config.ConfigPath = configPath
		})
	})

	When("specified flags are valid", func() {
		It("🧪 should: execute without error", func() {
			tester := helpers.CommandTester{
				Args: []string{"widget", "-p", "P?<date>", "-t", "42"},
				Root: rootCommand,
			}
			_, err := tester.Execute()
			Expect(err).Error().To(BeNil(),
				"should pass validation due to all flag being valid",
			)
		})
	})

	When("specified flags are valid", func() {
		It("🧪 should: return error due to option validation failure", func() {
			tester := helpers.CommandTester{
				Args: []string{"widget", "-p", "P?<date>", "-t", "99"},
				Root: rootCommand,
			}
			_, err := tester.Execute()
			Expect(err).Error().NotTo(BeNil(),
				"expected validation failure due to -t being within out of range",
			)
		})
	})
})
