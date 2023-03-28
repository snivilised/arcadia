package command_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/arcadia/src/app/command"
)

var _ = Describe("RootCmd", func() {
	It("🧪 should: execute", func() {
		Expect(command.Execute()).Error().To(BeNil())
	})
})
