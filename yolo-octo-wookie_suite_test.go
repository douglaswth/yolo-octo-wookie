package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestYoloOctoWookie(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "YoloOctoWookie Suite")
}

var _ = Describe("Yolo Octo Wookie", func() {
	Context("The Universe", func() {
		It("should be okay", func() {
			Expect(true).To(Equal(true))
		})
	})
})
