package main

import (
	"bufio"
	"fmt"
	"gifhelper"
	"image/png"
	"os"
	"strconv"
	"strings"
)

// The Cell storing the information of each entry of GameBoard
type Cell struct {
	strategy  string //represents "C" or "D" corresponding to the type of prisoner in the cell
	score float64 //represents the score of the cell based on the prisoner's relationship with neighboring cells
}

// The GameBoard is a 2D slice of Cell objects
type GameBoard [][]Cell
var Board GameBoard
var newBoard GameBoard
var rows int
var cols int
var b float64
var steps int
var bestScore float64
var GameBoards []GameBoard

func printBoard(board GameBoard) {
	fmt.Println()
	for row := 0; row < rows; row ++ {
		for col := 0; col < cols; col ++ {
			two_precision := fmt.Sprintf("%.2f",board[row][col].score)
			fmt.Print(board[row][col].strategy,two_precision,"  ")
		}
		fmt.Printf("\n\n")
	}
	fmt.Println()
}

func printAllBoards() {
	for _, board := range GameBoards {
		printBoard(board)
	}
}

//copy newBoard to Board
func copyToBoard(newBoard GameBoard) {
	for row := 0; row < rows; row ++ {
		for col := 0; col < cols; col ++ {
			if newBoard[row][col].strategy != "_" {
				Board[row][col].strategy = newBoard[row][col].strategy
			}
			Board[row][col].score = newBoard[row][col].score
		}
	}
}

//Contruct a board with size of rows * cols
func ConstructBoard(rows int, cols int) GameBoard {
	var Board = make([][]Cell,rows)
	for row := 0; row < rows; row ++ {
		Board[row]=make([]Cell,cols)
		for col := 0; col < cols; col ++ {
			Board[row][col].strategy = "_"
			Board[row][col].score = 0.0
		}
	}
	return Board
}

func UpdateAllScores() {
	for row := 0; row < rows; row ++ {
		for col := 0; col < cols; col ++ {
			UpdateOne(row, col, true)
		}
	}
}

//update the strategies in newBoard
func UpdateAllStrategies() {
	for row := 0; row < rows; row ++ {
		for col := 0; col < cols; col ++ {
			//fmt.Println("UPDATE STRATEGY OF BOARD ",row, col)
			UpdateOne(row, col, false)
			//fmt.Println()
		}
	}
}
//row - 1 >= 0, row + 1 < rows, col - 1 >= 0, col + 1 < cols
//update strategy to newBoard
//update score only if updateScore is true, and update strategy only if updateScore is false
func UpdateOne(row int, col int, updateScore bool) {
	//Initialize the score of the cell so that it won't be affected by previous board
	//Board[row][col].score = 0 TAKE NOTES HERE ! CANNOT UPDATE SCORE BEFORE APPLYING RULES, BECAUSE IT WILL JUST GET THE ONE WITH SCORE > 0
	//upper left [row-1][col-1]
	bestScore = Board[row][col].score //Don't change anything (score & strategy) during updating!!!
	if (row-1) >= 0 && (col-1) >= 0 {
		if updateScore == true {
			ApplyScoreRules(row, col, row-1, col-1)
		} else {
			ApplyStrategyRules(row, col, row-1, col-1)
		}
	}
	//up [row-1]
	if (row-1) >= 0 {
		if updateScore == true {
			ApplyScoreRules(row, col, row-1, col)
		} else {
			ApplyStrategyRules(row, col, row-1, col)
		}
	}
	//upper right [row-1][col+1]
	if (row-1) >= 0 && (col+1) < cols {
		if updateScore == true {
			ApplyScoreRules(row, col, row-1, col+1)
		} else {
			ApplyStrategyRules(row, col, row-1, col+1)
		}
	}
	//left [row][col-1]
	if (col-1) >= 0 {
		if updateScore == true {
			ApplyScoreRules(row, col, row, col-1)
		} else {
			ApplyStrategyRules(row, col, row, col-1)
		}
	}
	//right [row][col+1]
	if (col+1) < cols {
		if updateScore == true {
			ApplyScoreRules(row, col, row, col+1)
		} else {
			ApplyStrategyRules(row, col, row, col+1)
		}
	}
	//lower left [row+1][col-1]
	if (row+1) < rows && (col-1) >= 0 {
		if updateScore == true {
			ApplyScoreRules(row, col, row+1, col-1)
		} else {
			ApplyStrategyRules(row, col, row+1, col-1)
		}
	}
	//down [row+1][col]
	if (row+1) < rows {
		if updateScore == true {
			ApplyScoreRules(row, col, row+1, col)
		} else {
			ApplyStrategyRules(row, col, row+1, col)
		}
	}
	//lower right [row+1][col+1]
	if (row+1) < rows && (col+1) < cols{
		if updateScore == true {
			ApplyScoreRules(row, col, row+1, col+1)
		} else {
			ApplyStrategyRules(row, col, row+1, col+1)
		}
	}
}
func ApplyScoreRules(row int, col int, neighbor_row int, neighbor_col int) {
	if Board[row][col].strategy == "C" {
		if Board[neighbor_row][neighbor_col].strategy == "C" {
			Board[row][col].score += 1 // Okay since we reset the score to 0 in copyToBoard
		}
	} else {
		if Board[neighbor_row][neighbor_col].strategy == "C" {
			Board[row][col].score += b
		}
	}
}
/*
Apply strategy rules to newBoard
If current best score of a cell is smaller than the score of its neighbor, then adopt its strategy
and update best score to the neighbor's score
If current best score is still the best, do nothing?
 */
func ApplyStrategyRules(row int, col int, neighbor_row int, neighbor_col int) {
	//fmt.Println("self: ",row, col, ": ", Board[row][col].score, "neighbor: ", neighbor_row, neighbor_col, Board[neighbor_row][neighbor_col].score)
	if bestScore < Board[neighbor_row][neighbor_col].score {
		//fmt.Println("Change ", row, col," ",Board[row][col].strategy, "to ",neighbor_row, neighbor_col, " ",Board[neighbor_row][neighbor_col].strategy)
		newBoard[row][col].strategy = Board[neighbor_row][neighbor_col].strategy
		bestScore = Board[neighbor_row][neighbor_col].score
	}
}

func BoardsToList() GameBoard {
	Board_ := ConstructBoard(rows, cols)
	for row := 0; row < rows; row ++ {
		for col := 0; col < cols; col ++ {
			Board_[row][col].strategy = Board[row][col].strategy
			Board_[row][col].score = Board[row][col].score
		}
	}
	return Board_
}

func main() {
	//fmt.Printf("Spatial Game Start!!\n")
	// Process the arguments
	filename := os.Args[1]
	b_string:= os.Args[2]
	b_float, err := strconv.ParseFloat(b_string, 8)
	b = b_float
	if err != nil {
		panic("Error: Problem converting parameter b to a float\n")
	}
	steps, err1 := strconv.Atoi(os.Args[3])
	if err1 != nil {
		panic("Error: Problem converting parameter steps to an integer.\n")
	}
	//fmt.Println("Got the parameters successfully!")
	// Open the file given
	file, err2 := os.Open(filename)
	if err2 != nil {
		panic("Error: Problem opening the file\n")
	}
	//print("Opened the file ",filename," successfully!\n")

	// Read the file
	scanner := bufio.NewScanner(file)
	var row int
	for scanner.Scan() {
		//get number of rows and cols from first line of the file
		if row == 0 {
			firstLineString := scanner.Text()
			firstLine := strings.Split(firstLineString," ") //assume the row and col # are integers
			rows, _ = strconv.Atoi(firstLine[0])
			cols, _ = strconv.Atoi(firstLine[1])
			// Construct an empty Board, each cell with Strategy = "_", score = 0.0
			Board = ConstructBoard(rows, cols)
		} else {
			// Initialize the empty Board using the information of the file
			for col := 0; col < cols; col++ {
				Board[row-1][col].strategy = string(scanner.Text()[col]) //Be careful! It's [row-1][col]
			}
		}
		row ++
	}
	if scanner.Err() != nil {
		//fmt.Printf("Error: There was a problem reading the file\n")
		os.Exit(1)
	}
	file.Close()
	UpdateAllScores()
	//println(".........First Board.......")
	//printBoard(Board)
	Board_ := BoardsToList()
	GameBoards = append(GameBoards, Board_)
	for n := 0; n < steps; n++ {
		//fmt.Println("\nSTEP ",n, "...............")
		newBoard = ConstructBoard(rows, cols)
		//fmt.Println("NEWBOARD CONSTRUCTED")
		UpdateAllStrategies()
		copyToBoard(newBoard)
		UpdateAllScores()
		Board_ := BoardsToList()
		//printBoard(Board_)
		//fmt.Println("BOARD UPDATE FINISHED")
		GameBoards = append(GameBoards, Board_)
	}
	imglist := DrawGameBoards(GameBoards, 20)
	//fmt.Println("Boards drawn to images! Now, convert to animated GIF.")
	// convert images to a GIF
	var outputFile string = "out"
	gifhelper.ImagesToGIF(imglist, outputFile)
	out_png := imglist[len(imglist)-1]
	file, _ = os.Create("Prisoners.png")
	png.Encode(file, out_png)
	file.Close()
	//fmt.Println("Success! GIF produced.")
	//fmt.Println("Spatial Game Finished!!")
}



