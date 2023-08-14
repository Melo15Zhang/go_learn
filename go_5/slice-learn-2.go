package main

import (
	"fmt"
	"strings"
)

func main() {

	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// The players take turns.
	board[0][0] = "S"
	board[0][2] = "S"
	board[1][0] = "0"
	board[1][2] = "O"
	board[2][0] = "S"
	board[2][2] = "S"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
