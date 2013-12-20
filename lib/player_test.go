package lib_test

import (
	lib "github.com/ciarand/go-tic-tac-toe/lib"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Player", func() {
	It("should create a player from the provided string", func() {
		player := lib.NewPlayer("X")
		Expect(player.Piece()).To(Equal("X"))
	})

	It("should be able to change pieces", func() {
		player := lib.NewPlayer("X")
		player.SetPiece("O")
		Expect(player.Piece()).To(Equal("O"))
	})

	It("should initialize with a score of 0", func() {
		player := lib.NewPlayer("X")
		Expect(player.Score()).To(Equal(0))
	})

	It("should be able to increment the score", func() {
		player := lib.NewPlayer("X")
		player.IncrementScore()
		Expect(player.Score()).To(Equal(1))
		for i := 0; i < 3; i += 1 {
			player.IncrementScore()
		}
		Expect(player.Score()).To(Equal(4))
	})
})
