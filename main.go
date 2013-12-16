package main

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	BOARD_WIDTH     = 9
	BOARD_HEIGHT    = 9
	NUMBER_OF_CELLS = BOARD_WIDTH * BOARD_HEIGHT
	MATCHES_TO_WIN  = 5
)

func main() {
	fmt.Println("Hello, welcome to tic-tac-toe!")
	var input int
	var board *Board

	for board = NewBoard(); board.IsGameOver() != true; {
		fmt.Println()
		fmt.Println(board)
		fmt.Printf("Enter your move (1-%d)", NUMBER_OF_CELLS)
		fmt.Scanf("%d", &input)
		if input > NUMBER_OF_CELLS || input < 1 {
			fmt.Println("Out of bounds! Try again")
			continue
		}
		fmt.Printf("Your move was %d", input)

		board.PlacePiece("X", input-1)
	}

	if len(board.Winner) != 0 {
		fmt.Printf("\n%s won the game!", board.Winner)
	} else {
		fmt.Println("It was a tie")
	}

	fmt.Println()
	fmt.Println(board)
}

type Cell struct {
	Value string
}

type Board struct {
	Cells  [NUMBER_OF_CELLS]*Cell
	Width  int
	Height int
	Winner string
}

func (b *Board) String() string {
	var buffer bytes.Buffer
	var rowDiv string = strings.Repeat("-", (BOARD_WIDTH*2)+1)

	buffer.WriteString(rowDiv)
	for i := range b.Cells {

		if i%BOARD_WIDTH == 0 {
			buffer.WriteString("\n")
			buffer.WriteString("|")
		}

		if b.Cells[i].IsOccupied() {
			buffer.WriteString(b.Cells[i].Value)
		} else {
			buffer.WriteString(" ")
		}

		buffer.WriteString("|")
	}
	buffer.WriteString("\n")
	buffer.WriteString(rowDiv)

	return buffer.String()
}

func (b *Board) PlacePiece(piece string, index int) {
	b.Cells[index].Value = piece
}

func (c Cell) IsOccupiedBy(occ string) bool {
	return c.Value == occ
}

func (c Cell) IsOccupied() bool {
	return len(c.Value) != 0
}

func NewBoard() *Board {
	board := new(Board)
	for i := range board.Cells {
		board.Cells[i] = new(Cell)
	}
	board.Height = BOARD_HEIGHT
	board.Width = BOARD_WIDTH
	return board
}

func (b *Board) IsGameOver() bool {
	// First check the rows
	var i, j, k int
	matches := 0
	piece := "X"

	for i = 0; i < b.Height; i += 1 {
		for j = 0; j < b.Width; j += 1 {
			// K is the current cell we're checking
			k = j + (i * b.Width)
			if b.Cells[k].IsOccupiedBy(piece) {
				matches += 1
			} else {
				matches = 0
			}

			if matches >= MATCHES_TO_WIN {
				b.Winner = piece
				return true
			}
		}
		matches = 0
	}

	// Next, check the columns
	for j = 0; j < b.Width; j += 1 {
		for i = 0; i < b.Height; i += 1 {
			k = j + (i * b.Width)
			if b.Cells[k].IsOccupiedBy(piece) {
				matches += 1
			} else {
				matches = 0
			}

			if matches >= MATCHES_TO_WIN {
				b.Winner = piece
				return true
			}
		}
		matches = 0
	}

	for i = range b.Cells {
		if !b.Cells[i].IsOccupied() {
			return false
		}
	}
	// No luck? We don't support diagonals at this time
	return true
}
