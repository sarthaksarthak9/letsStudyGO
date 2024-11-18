package newginkgo_test

import (
	adt "example.com/letsStudyGO/testing/03ginkgo/code"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var s *adt.Set

var _ = Describe("Set", func() {

	BeforeEach(func() {
		s = adt.NewSet()
	})

	Describe("Emptiness", func() {
		Context("When the set does not contain items", func() {
			It("Should be empty", func() {

				Expect(s.IsEmpty()).To(BeTrue())
			})

		})

		Context("When set contain items", func() {
			It("Should add elements", func() {

				s.ADD("blue")
				Expect(s.IsEmpty()).To(BeFalse())
			})
		})

	})

	Describe("Size", func() {
		Context("As items are added", func() {
			It("Should return increasing size", func() {

				By("Empty set size being 0", func() {
					Expect(s.Size()).To(BeZero())
				})

				By("Adding First item", func() {
					s.ADD("red")
					Expect(s.Size()).To(Equal(1))
				})
			})
		})
	})

	Describe("Contains", func() {
		Context("When Red has not been added", func() {
			It("Should return false", func() {
				Expect(s.Contains("red")).To(BeFalse())
			})
		})

		Context("When Red has been added", func() {
			It("Should return true", func() {

				s.ADD("red")
				Expect(s.Contains("red")).To(BeTrue())
			})
		})
	})
})
