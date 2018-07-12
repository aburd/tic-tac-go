package main

import (
	"fmt"

	"github.com/aburd/tictacgo/internal/lib"
)

// func checkSpace(board [][]string, row int, column int) {

// }

func main() {
	board := lib.ClearBoard()
	player := 0

	fmt.Println("Tic-Tac-Go")
	fmt.Println("---------------------")
	fmt.Println("")

	for ; !lib.IsCatGame(&board); lib.CheckWin(board, player) {
		fmt.Printf("It's player %v's turn.\n", player+1)
		lib.DrawBoard(board)
		fmt.Println("")
		catGame, column, play := lib.GetPlayFromPlayer(player)
		lib.PlaySpace(board, catGame, column, play)
		if player == 0 {
			player = 1
		} else {
			player = 0
		}
		fmt.Print("\n")
	}
}
