package lib_test

import (
	lib "github.com/ciarand/go-tic-tac-toe/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Board", func() {
	It("should construct with the provided params", func() {
		height, width, matches := 10, 10, 10
		b := lib.NewBoard(height, width, matches)

		Expect(b.Width).To(Equal(width))
		Expect(b.Height).To(Equal(height))
		Expect(b.MatchesToWin).To(Equal(matches))
	})

	It("should print an empty board correctly", func() {
		b := lib.NewBoard(3, 3, 3)
		str := "-------\n" +
			"| | | |\n" +
			"| | | |\n" +
			"| | | |\n" +
			"-------"

		Expect(b.String()).To(Equal(str))
	})
})
