package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func drawBoard(board [][]string) {
	fmt.Println("THE BOARD:")
	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], "|"))
	}
}

func askPlayAgain() bool {
	input := ""
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Play again? ")
		scanner.Scan()
		input = scanner.Text()
		if input == "y" || input == "n" {
			break
		}
	}
	return input == "y"
}

func clearBoard() [][]string {
	return [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
}

func isCatGame(board *[][]string) bool {
	catGame := true
	boardCopy := *board
	for i := 0; i < len(boardCopy) && catGame; i++ {
		line := boardCopy[i]
		for j := 0; j < len(line) && catGame; j++ {
			space := line[j]
			catGame = space != "-"
		}
	}
	if catGame {
		fmt.Println("Cat game!")
		fmt.Println("")
		playAgain := askPlayAgain()
		if playAgain {
			fmt.Println("Clearing board...")
			fmt.Println("")
			*board = clearBoard()
			return false
		}
	}
	return catGame
}

// func checkWin(board) {

// }

// func checkSpace(board [][]string, row int, column int) E

func playSpace(board [][]string, row int, column int, play string) [][]string {
	board[row][column] = play
	return board
}

func getRow() int {
	row := 0
	for row == 0 || row > 3 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("What row?")
		fmt.Println("Enter 1, 2, 3")
		scanner.Scan()
		row, _ = strconv.Atoi(scanner.Text())
	}
	// Decrease by one to have correct index in array
	return row - 1
}

func getColumn() int {
	column := 0
	for column == 0 || column > 3 {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("What column?")
		fmt.Println("Enter 1, 2, 3")
		scanner.Scan()
		column, _ = strconv.Atoi(scanner.Text())
	}
	// Decrease by one to have correct index in array
	return column - 1
}

func getPlayFromPlayer(player int) (int, int, string) {
	plays := []string{"X", "O"}
	row := getRow()
	column := getColumn()
	return row, column, plays[player]
}

func main() {
	board := clearBoard()
	player := 0

	fmt.Println("Tic-Tac-Go")
	fmt.Println("---------------------")
	fmt.Println("")

	for !isCatGame(&board) {
		fmt.Printf("It's player %v's turn.", player+1)
		drawBoard(board)
		fmt.Println("")
		catGame, column, play := getPlayFromPlayer(player)
		playSpace(board, catGame, column, play)
		if player == 0 {
			player = 1
		} else {
			player = 0
		}
		fmt.Print("\n")
	}
}
