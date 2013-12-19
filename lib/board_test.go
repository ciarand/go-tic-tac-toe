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

	It("should know when the game is a tie", func() {
		b := lib.NewBoard(3, 3, 3)
		Expect(b.IsGameOver()).To(Equal(false))

		for i := 0; i < 9; i += 1 {
			b.PlacePiece("X", i)
		}

		Expect(b.IsGameOver()).To(Equal(true))
	})

	It("should know when someone has won horizontally", func() {
		b := lib.NewBoard(3, 3, 3)

		for i := 0; i < 3; i += 1 {
			b.PlacePiece("X", i)
		}

		Expect(b.IsGameOver()).To(Equal(true))
	})

	It("should know when someone has won vertically", func() {
		b := lib.NewBoard(3, 3, 3)

		for i := 0; i < 3; i += 1 {
			b.PlacePiece("X", i*3)
		}

		Expect(b.IsGameOver()).To(Equal(true))
	})

	It("should know when someone has won NW to SE diagonally", func() {
		b := lib.NewBoard(3, 3, 3)

		for _, i := range [3]int{0, 4, 8} {
			b.PlacePiece("X", i)
		}

		Expect(b.IsGameOver()).To(Equal(true), "%s", b)
	})

	It("should know when someone has won NE to SW diagonally", func() {
		b := lib.NewBoard(3, 3, 3)

		for _, i := range [3]int{2, 4, 6} {
			b.PlacePiece("X", i)
		}

		Expect(b.IsGameOver()).To(Equal(true), "%s", b)
	})
})
