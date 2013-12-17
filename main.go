package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, welcome to tic-tac-toe!")
	var input int
	var board *Board
	var width, height, matches_to_win, number_of_cells int

	fmt.Printf("Enter the width of the board: ")
	fmt.Scanf("%d", &width)

	fmt.Printf("Enter the height of the board: ")
	fmt.Scanf("%d", &height)

	fmt.Printf("Finally, enter the number of matches required to win: ")
	fmt.Scanf("%d", &matches_to_win)

	number_of_cells = width * height

	for board = NewBoard(height, width, matches_to_win); board.IsGameOver() != true; {
		fmt.Println()
		fmt.Println(board)
		fmt.Printf("Enter your move (1-%d)", number_of_cells)
		fmt.Scanf("%d", &input)
		if input > number_of_cells || input < 1 {
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
	Cells        []*Cell
	Width        int
	Height       int
	MatchesToWin int
	Winner       string
}

func (b *Board) String() string {
	var buffer bytes.Buffer
	var rowDiv string = strings.Repeat("-", (b.Width*2)+1)

	buffer.WriteString(rowDiv)
	for i := range b.Cells {

		if i%b.Width == 0 {
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

func NewBoard(height, width, matches_to_win int) *Board {
	board := new(Board)
	board.Cells = make([]*Cell, height*width)
	for i := range board.Cells {
		board.Cells[i] = new(Cell)
	}
	board.Height = height
	board.Width = width
	board.MatchesToWin = matches_to_win

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

			if matches >= b.MatchesToWin {
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

			if matches >= b.MatchesToWin {
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
