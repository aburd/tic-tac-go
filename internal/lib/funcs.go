package lib

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// BOARD FUNCS

func DrawBoard(board [][]string) {
	fmt.Println("THE BOARD:")
	for i := 0; i < len(board); i++ {
		fmt.Println(strings.Join(board[i], "|"))
	}
}

func ClearBoard() [][]string {
	return [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}
}

// BOARD CHECKS

func IsCatGame(board *[][]string) bool {
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
		playAgain(board)
	}
	return catGame
}

func CheckWin(board [][]string, player int) {
	for i := 0; i < len(board); i++ {
		line := board[i]
		checkRow(line, board)
		checkColumn(i, board)
	}
}

func checkRowOrColumn(plays int, rcLen int, winMsg string, board *[][]string) {
	if plays >= rcLen {
		DrawBoard(*board)
		fmt.Println(winMsg)
		playAgain(board)
	}
}

func checkRow(line []string, board [][]string) {
	var xs int
	var os int
	lineLen := len(line)
	for i := 0; i < len(line); i++ {
		space := line[i]
		if space == "X" {
			xs += 1
		}
		if space == "O" {
			os += 1
		}
	}
	checkRowOrColumn(xs, lineLen, "Xs win!", &board)
	checkRowOrColumn(os, lineLen, "Os win!", &board)
}

func checkColumn(columnIndex int, board [][]string) {
	var xs int
	var os int
	columnLen := len(board)
	for i := 0; i < len(board); i++ {
		space := board[i][columnIndex]
		if space == "X" {
			xs += 1
		}
		if space == "O" {
			os += 1
		}
	}
	checkRowOrColumn(xs, columnLen, "Xs win!", &board)
	checkRowOrColumn(os, columnLen, "Os win!", &board)
}

// PLAYER INTERACTION

func PlaySpace(board [][]string, row int, column int, play string) [][]string {
	board[row][column] = play
	return board
}

func GetPlayFromPlayer(player int) (int, int, string) {
	plays := []string{"X", "O"}
	row := getRow()
	column := getColumn()
	return row, column, plays[player]
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

func playAgain(board *[][]string) bool {
	fmt.Println("")
	playAgain := askPlayAgain()
	if playAgain {
		fmt.Println("Clearing board...")
		fmt.Println("")
		*board = ClearBoard()
		return false
	}
	return false
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
