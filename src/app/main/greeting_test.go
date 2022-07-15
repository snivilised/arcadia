package main_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/snivilised/arcadia/src/app/main"
)

var _ = Describe("Greeting", func() {
	It("should: get greeting", func() {
		greeting := main.Greeting("fido")
		Expect(greeting).To(Equal("hello fido"))
	})
})
