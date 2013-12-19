package lib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type boundryTest struct {
	n        int
	expected bool
}

var _ = Describe("cell boundry tests", func() {
	var (
		b *Board
	)

	BeforeEach(func() {
		// Create a new board
		b = NewBoard(20, 20, 20)
	})

	It("should accurately report when cell is on top row", func() {
		dataSet := []boundryTest{
			{0, true}, {1, true}, {9, true}, {19, true},
			{20, false}, {22, false}, {399, false},
		}

		for _, data := range dataSet {
			Expect(b.isCellOnTopRow(data.n)).To(Equal(data.expected))
		}
	})

	It("should accurately report when cell is on bottom row", func() {
		dataSet := []boundryTest{
			{380, true}, {388, true}, {399, true},
			{0, false}, {1, false}, {9, false}, {19, false},
		}

		for _, data := range dataSet {
			Expect(b.isCellOnBottomRow(data.n)).To(Equal(data.expected))
		}
	})

	It("should accurately report when cell is on left col", func() {
		dataSet := []boundryTest{
			{0, true}, {20, true}, {380, true},
			{1, false}, {18, false}, {19, false}, {379, false},
		}

		for _, data := range dataSet {
			Expect(b.isCellOnLeftCol(data.n)).To(Equal(data.expected))
		}

	})

	It("should accurately report when cell is on right col", func() {
		dataSet := []boundryTest{
			{399, true}, {19, true}, {39, true},
			{0, false}, {1, false}, {377, false}, {398, false},
		}

		for _, data := range dataSet {
			Expect(b.isCellOnRightCol(data.n)).To(Equal(data.expected))
		}
	})
})

var _ = Describe("win conditions", func() {
	Context("should accurately report a horizontal win condition", func() {
		piece := "X"
		b := NewBoard(20, 20, 5)

		// Sanity check, empty board should not have a win condition
		It("should not report a win condition for an empty board", func() {
			Expect(b.hasHorizontalWin(piece)).To(Equal(false), "%s", b)
		})

		It("should not report a win condition for an insufficient number of matches", func() {
			for i := 0; i < 4; i += 1 {
				b.PlacePiece(piece, i)
			}
			Expect(b.hasHorizontalWin(piece)).To(Equal(false), "%s", b)
		})

		It("should report a win condition for a sufficient number of matches", func() {
			b.PlacePiece(piece, 4)
			Expect(b.hasHorizontalWin(piece)).To(Equal(true), "%s", b)
		})

		It("should continue to report a win condition for an overly sufficient number of matches", func() {
			b.PlacePiece(piece, 5)
			b.PlacePiece(piece, 6)
			b.PlacePiece(piece, 7)
			Expect(b.hasHorizontalWin(piece)).To(Equal(true), "%s", b)
		})
	})

	Context("should accurately report a vertical win condition", func() {
		piece := "X"
		b := NewBoard(20, 20, 5)

		It("should not report a win condition for an empty board", func() {
			Expect(b.hasVerticalWin(piece)).To(Equal(false), "%s", b)
		})

		It("should not report a win condition for an insufficient number of matches", func() {
			for i := 0; i < 4; i += 1 {
				b.PlacePiece(piece, i*b.Width)
			}
			Expect(b.hasVerticalWin(piece)).To(Equal(false), "%s", b)
		})

		It("should report a win condition for a sufficient number of matches", func() {
			b.PlacePiece(piece, 4*b.Width)
			Expect(b.hasVerticalWin(piece)).To(Equal(true), "%s", b)
		})

		It("should continue to report a win condition for an overly sufficient number of matches", func() {
			b.PlacePiece(piece, 5*b.Width)
			b.PlacePiece(piece, 6*b.Width)
			b.PlacePiece(piece, 7*b.Width)
			Expect(b.hasVerticalWin(piece)).To(Equal(true), "%s", b)
		})
	})

	Context("should accurately report a top left to bottom right win condition", func() {
		piece := "X"
		b := NewBoard(20, 20, 5)

		It("should not report a win condition for an empty board", func() {
			Expect(b.hasDiagonalWin(piece)).To(Equal(false), "%s", b)
		})

		It("should not report a win condition for an insufficient number of matches", func() {
			for i := 0; i < 4; i += 1 {
				b.PlacePiece(piece, (i*b.Width)+i)
			}
			Expect(b.hasDiagonalWin(piece)).To(Equal(false), "%s", b)
		})

		It("should report a win condition for a sufficient number of matches", func() {
			i := 4
			b.PlacePiece(piece, (i*b.Width)+i)
			Expect(b.hasDiagonalWin(piece)).To(Equal(true), "%s", b)
		})

		It("should continue to report a win condition for an overly sufficient number of matches", func() {
			for i := 5; i < 8; i += 1 {
				b.PlacePiece(piece, (i*b.Width)+i)
			}
			Expect(b.hasDiagonalWin(piece)).To(Equal(true), "%s", b)
		})
	})

	Context("should accurately report a top right to bottom left win condition", func() {
		piece := "X"
		b := NewBoard(20, 20, 5)

		It("should not report a win condition for an empty board", func() {
			Expect(b.hasDiagonalWin(piece)).To(Equal(false), "%s", b)
		})

		It("should not report a win condition for an insufficient number of matches", func() {
			for i := 0; i < 4; i += 1 {
				b.PlacePiece(piece, ((i+1)*b.Width)-(i+1))
			}
			Expect(b.hasDiagonalWin(piece)).To(Equal(false), "%s", b)
		})

		It("should report a win condition for a sufficient number of matches", func() {
			i := 4
			b.PlacePiece(piece, ((i+1)*b.Width)-(i+1))
			Expect(b.hasDiagonalWin(piece)).To(Equal(true), "%s", b)
		})

		It("should continue to report a win condition for an overly sufficient number of matches", func() {
			for i := 5; i < 8; i += 1 {
				b.PlacePiece(piece, ((i+1)*b.Width)-(i+1))
			}
			Expect(b.hasDiagonalWin(piece)).To(Equal(true), "%s", b)
		})
	})

	It("should know when a board is empty and full", func() {
		b := NewBoard(3, 3, 3)
		Expect(b.isFull()).To(Equal(false), "%s", b)
		for i := 0; i < b.Height*b.Width; i += 1 {
			b.PlacePiece("X", i)
		}
		Expect(b.isFull()).To(Equal(true), "%s", b)
	})
})
