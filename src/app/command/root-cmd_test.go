package command_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive // ok for testing
	. "github.com/onsi/gomega"    //nolint:revive // ok for testing

	"github.com/snivilised/arcadia/src/app/command"
	"github.com/snivilised/arcadia/src/internal/helpers"
	nef "github.com/snivilised/nefilim"
)

var _ = Describe("RootCmd", Ordered, func() {
	var (
		repo     string
		l10nPath string
	)

	BeforeAll(func() {
		repo = helpers.Repo("")
		l10nPath = helpers.Path(repo, "test/data/l10n")
		fS := nef.NewUniversalABS()
		Expect(fS.DirectoryExists(l10nPath)).To(BeTrue())
	})

	It("🧪 should: execute", func() {
		bootstrap := command.Bootstrap{}
		tester := helpers.CommandTester{
			Args: []string{},
			Root: bootstrap.Root(func(co *command.ConfigureOptions) {
				co.Detector = &DetectorStub{}
				co.Config.Name = configName
				co.Config.ConfigPath = configPath
			}),
		}
		_, err := tester.Execute()
		Expect(err).Error().To(BeNil())
	})
})
