package command_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/snivilised/arcadia/src/app/command"
	"github.com/snivilised/arcadia/src/internal/helpers"
	"github.com/snivilised/extendio/xfs/utils"

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
		repo = helpers.Repo("../..")
		l10nPath = helpers.Path(repo, "test/data/l10n")
		Expect(utils.FolderExists(l10nPath)).To(BeTrue())
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
