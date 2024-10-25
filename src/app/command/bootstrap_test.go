package command_test

import (
	. "github.com/onsi/ginkgo/v2" //nolint:revive // ok for testing
	. "github.com/onsi/gomega"    //nolint:revive // ok for testing
	"github.com/snivilised/arcadia/src/app/command"
	"github.com/snivilised/arcadia/src/internal/helpers"
	nef "github.com/snivilised/nefilim"

	"golang.org/x/text/language"
)

const (
	configName = "arcadia-test"
	configPath = "../../test/data/configuration"
)

type DetectorStub struct {
}

func (j *DetectorStub) Scan() language.Tag {
	return language.BritishEnglish
}

var _ = Describe("Bootstrap", Ordered, func() {

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

	Context("given: root defined with widget sub-command", func() {
		It("🧪 should: setup command without error", func() {
			bootstrap := command.Bootstrap{}
			rootCmd := bootstrap.Root(func(co *command.ConfigureOptions) {
				co.Detector = &DetectorStub{}
				co.Config.Name = configName
				co.Config.ConfigPath = configPath
			})
			Expect(rootCmd).NotTo(BeNil())
		})
	})
})
