package lib_test

import (
	lib "github.com/ciarand/go-tic-tac-toe/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Cell", func() {
	It("should initialize as unoccupied", func() {
		c := lib.Cell{}
		Expect(c.IsOccupied()).To(Equal(false))
	})

	It("should be occupied when a piece is on it", func() {
		c := lib.Cell{}
		c.SetValue("X")
		Expect(c.IsOccupied()).To(Equal(true))
	})

	It("should only return true for the same piece", func() {
		c := lib.Cell{}
		c.SetValue("X")
		Expect(c.IsOccupiedBy("Y")).To(Equal(false))
		Expect(c.IsOccupiedBy("X")).To(Equal(true))
	})
})
