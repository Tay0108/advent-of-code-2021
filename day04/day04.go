package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const BOARD_WIDTH = 5
const BOARD_HEIGHT = 5

type BingoCell struct {
	value    int
	isMarked bool
}

type BingoBoard struct {
	victorious bool
	cells      [BOARD_HEIGHT][BOARD_WIDTH]BingoCell
}

func convertStringLineToInts(line string, separator string) []int {
	var ints []int
	strings := strings.Split(line, separator)
	for _, value := range strings {
		parsedNumber, _ := strconv.ParseInt(value, 10, 64)
		ints = append(ints, int(parsedNumber))
	}
	return ints
}

func createBoardRow(row string) [BOARD_WIDTH]BingoCell {
	var bingoBoardRow [BOARD_WIDTH]BingoCell
	values := strings.Fields(row)
	for index, stringNum := range values {
		num, _ := strconv.ParseInt(stringNum, 10, 64)
		bingoCell := BingoCell{value: int(num), isMarked: false}
		bingoBoardRow[index] = bingoCell
	}
	return bingoBoardRow
}

func loadInputFile(path string) ([]int, []BingoBoard) {
	file, error := os.Open(path)
	if error != nil {
		log.Fatal(error)
	}
	scanner := bufio.NewScanner(file)

	var bingoInputSequence []int
	var bingoBoards []BingoBoard

	readingFirstLine := true

	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		if readingFirstLine {
			bingoInputString := scanner.Text()
			bingoInputSequence = convertStringLineToInts(bingoInputString, ",")
			readingFirstLine = false
			continue
		}

		boardRowIndex := 0
		var currentBoard BingoBoard
		for scanner.Text() != "" { // read board
			textRow := scanner.Text()
			boardRow := createBoardRow(textRow)
			currentBoard.cells[boardRowIndex] = boardRow
			currentBoard.victorious = false
			scanner.Scan()
			boardRowIndex++
		}
		bingoBoards = append(bingoBoards, currentBoard)
	}

	file.Close()

	return bingoInputSequence, bingoBoards
}

func markValueInBoard(value int, board *BingoBoard) bool { // returns if board won
	for i := 0; i < len(board.cells); i++ {
		row := &board.cells[i]
		for j := 0; j < len(row); j++ {
			cell := &row[j]
			if cell.value == value {
				cell.isMarked = true
				if boardWon(board, i, j) {
					sumOfUnmarkedNumbers := calculateSumOfUnmarkedNumbers(board)
					fmt.Println("Board won on value:", value)
					fmt.Println("Sum of unmarked numbers:", sumOfUnmarkedNumbers)
					fmt.Println("Final score:", sumOfUnmarkedNumbers*value)
					board.victorious = true
					return true
				}
			}
		}
	}
	return false
}

func playBingo(inputSequence *[]int, bingoBoards *[]BingoBoard) {
	for i := 0; i < len(*inputSequence); i++ {
		for j := 0; j < len(*bingoBoards); j++ {
			if (*bingoBoards)[j].victorious {
				continue
			}
			markValueInBoard((*inputSequence)[i], &(*bingoBoards)[j])
		}
	}
}

func checkRowVictory(board *BingoBoard, markedCellRow int) bool {
	for i := 0; i < BOARD_WIDTH; i++ {
		if !board.cells[markedCellRow][i].isMarked {
			return false
		}
	}
	return true
}

func checkColumnVictory(board *BingoBoard, markedCellColumn int) bool {
	for i := 0; i < BOARD_HEIGHT; i++ {
		if !board.cells[i][markedCellColumn].isMarked {
			return false
		}
	}
	return true
}

func boardWon(board *BingoBoard, markedCellRow int, markedCellColumn int) bool {
	return checkRowVictory(board, markedCellRow) || checkColumnVictory(board, markedCellColumn)
}

func calculateSumOfUnmarkedNumbers(board *BingoBoard) int {
	sum := 0
	for i := 0; i < BOARD_HEIGHT; i++ {
		for j := 0; j < BOARD_WIDTH; j++ {
			if !board.cells[i][j].isMarked {
				sum += board.cells[i][j].value
			}
		}
	}
	return sum
}

func main() {
	bingoInputSequence, bingoBoards := loadInputFile("./input.txt")
	playBingo(&bingoInputSequence, &bingoBoards)
}

/*
	Algorytm:
	1. Bierzemy kolejny numer z sekwencji
	2. Sprawdzamy kazdą z plansz:
	3. szukamy podanego numeru w planszy iterując po niej.
	4. Jeśli liczba znajduje się na planszy zamalowujemy ją (jakaś mapa?), a następnie sprawdzamy kolumnę i wiersz do której nalezy, czy jest juz zwyciestwo. Jesli jest zwyciestwo to zwroc numer planszy (i z petli + 1)
	5. Sprawdzamy kolejna plansze jesli nie ma zwyciestwa
*/
