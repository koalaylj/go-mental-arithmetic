package main

import (
	"fmt"
	"time"
)

var ops = [...]string{"+", "-"}

const ARG_MAX = 10
const ARG_MIN = 1
const ARG_COL = 4
const ARG_ROW = 15
const ARG_PAGES = 10

func main() {

	begin := time.Now()
	m := createPDF()

	setHeader(m)
	setFooter(m)

	for i := 0; i < ARG_PAGES; i++ {
		cells := innerText(ARG_ROW, ARG_COL)
		buildPage(m, cells)
		if i < ARG_PAGES-1 {
			m.AddPage()
		}
	}

	save(m)

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func innerText(arg_row int, arg_col int) [][]string {
	cells := [][]string{}

	for row_i := 0; row_i < arg_row; row_i++ {

		row := []string{}

		for col_i := 0; col_i < arg_col; col_i++ {
			no := row_i + col_i*arg_row + 1

			random := getRandom()

			op := randomOp(random)

			item := ""

			switch op {
			case "+":
				item = randomAdd(random)
			case "-":
				item = randomSub(random)
			}

			cell := fmt.Sprintf("%2d) %s\n", no, item)

			row = append(row, cell)
		}

		cells = append(cells, row)
	}

	return cells
}
