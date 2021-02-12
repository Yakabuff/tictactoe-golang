package main

import "fmt"

var b *board

type player struct {
	symbol string
	name   string
}

type board struct {
	board [3][3]string
	p1    *player
	p2    *player
	curr  *player
}

func makeMove(x int, y int) {

	if b.board[x][y] == "." {
		b.board[x][y] = b.curr.symbol
		checkWin(x, y)
		if b.curr == b.p1 {
			b.curr = b.p2
			fmt.Println("switching to p2")
		} else {
			b.curr = b.p1
			fmt.Println("switching to p1")
		}
	} else {
		fmt.Println("Slot already occupied")
	}

}

//816
//357
//492
func checkWin(x int, y int) {
	col, row, diag, rdiag := 0, 0, 0, 0
	for i := 0; i <= 2; i++ {
		if b.board[x][i] == b.curr.symbol {
			col++

		}
		if b.board[i][y] == b.curr.symbol {
			row++

		}
		if b.board[i][i] == b.curr.symbol {
			diag++

		}
		if b.board[i][3-i-1] == b.curr.symbol {
			rdiag++
		}

	}
	if row == 3 || diag == 3 || rdiag == 3 || col == 3 {
		win = true
		winner = b.curr.name
	}
}

func newBoard(p1 *player, p2 *player) *board {
	b := board{board: [3][3]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}, p1: p1, p2: p2, curr: p1}
	return &b
}

var win bool = false
var winner string = ""

func main() {
	var inputX int
	var inputY int
	var p1Name string
	var p2Name string

	fmt.Println("Enter p1 name")
	fmt.Scanln(&p1Name)
	p1 := player{symbol: "o", name: p1Name}

	fmt.Println("Enter p2 name")
	fmt.Scanln(&p2Name)
	p2 := player{symbol: "x", name: p2Name}

	b = newBoard(&p1, &p2)

	fmt.Println(b.p1.name)
	fmt.Println((b.p2.name))
	for win == false {

		fmt.Println("make move: " + b.curr.name)
		fmt.Scanln(&inputX)
		if inputX > 2 || inputX < 0 {
			continue
		}
		fmt.Scanln(&inputY)
		if inputY > 2 || inputY < 0 {
			continue
		}
		fmt.Printf("X value: %d, Y value: %d", inputX, inputY)
		makeMove(inputX, inputY)
		fmt.Println(b.curr.name)
	}
	fmt.Println("winner is " + winner + " !!!!!!!!!!!!!!!!!!!")

}
