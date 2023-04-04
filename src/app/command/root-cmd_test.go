package command_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/arcadia/src/app/command"
	"github.com/snivilised/arcadia/src/internal/helpers"
	xi18n "github.com/snivilised/extendio/i18n"
	"github.com/snivilised/extendio/xfs/utils"
)

var _ = Describe("RootCmd", Ordered, func() {
	var (
		repo     string
		l10nPath string
		from     *xi18n.LoadFrom
	)

	BeforeAll(func() {
		repo = helpers.Repo("../..")
		l10nPath = helpers.Path(repo, "test/data/l10n")
		Expect(utils.FolderExists(l10nPath)).To(BeTrue())
	})

	BeforeEach(func() {
		from = &xi18n.LoadFrom{
			Path: l10nPath,
			Sources: xi18n.TranslationFiles{
				xi18n.SOURCE_ID: xi18n.TranslationSource{
					Name: fmt.Sprintf("test.%v", command.ApplicationName),
				},
			},
		}
	})

	It("🧪 should: execute", func() {
		Expect(command.Execute(func(o *command.ExecutionOptions) {
			o.Detector = &DetectorStub{}
			o.From = from
		})).Error().To(BeNil())
	})
})
