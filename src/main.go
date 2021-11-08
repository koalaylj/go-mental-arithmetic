package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

var ops = [...]string{"+", "-"}

const max = 10
const min = 1
const ARG_PAGE_SIZE = 60
const ARG_COL = 4
const ARG_PAGES = 3

func main() {

	page := [][]string{}

	for row_i := 0; row_i < ARG_PAGE_SIZE/ARG_COL; row_i++ {

		row := []string{}

		for col_i := 0; col_i < ARG_COL; col_i++ {
			no := row_i + col_i*ARG_PAGE_SIZE/ARG_COL + 1

			random := getRandom()

			op := randomOp(random)

			item := ""

			switch op {
			case "+":
				item = randomAdd(random, min, max)
			case "-":
				item = randomSub(random, min, max)
			}

			cell := fmt.Sprintf("%2d) %s\n", no, item)

			row = append(row, cell)
		}

		page = append(page, row)
	}

	// testPDF()
	toPDF(page)
}

func testPDF() {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.Letter)
	m.SetBorder(false)

	m.Row(40, func() {
		m.Col(3, func() {
			m.Text("Any Text1")
		})
		m.Col(3, func() {
			m.Text("Any Text2")
		})
		m.Col(3, func() {
			m.Text("Any Text3")
		})
		m.Col(3, func() {
			m.Text("Any Text3")
		})
	})

	err := m.OutputFileAndClose("./textgrid.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func toPDF(page [][]string) {
	begin := time.Now()
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetBorder(false)

	newPage(m, page)

	// for i := 0; i < 10; i++ {
	// 	// newPage(m,[])

	// 	if i < 9 {
	// 		m.AddPage()
	// 	}
	// }

	err := m.OutputFileAndClose("./fuck.pdf")
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}

	end := time.Now()
	fmt.Println(end.Sub(begin))
}

func newPage(m pdf.Maroto, cells [][]string) {

	for i := 0; i < ARG_PAGE_SIZE/ARG_COL; i++ {
		// for j := 0; j < ARG_COL; j++ {
		// fmt.Printf("a[%d][%d] = %s\n", i, j, page[i][j])
		m.Row(15, func() {
			m.Col(3, func() {
				m.Text(cells[i][0], props.Text{
					Style: consts.Normal,
					Align: consts.Left,
					Size:  14,
				})
			})
			m.Col(3, func() {
				m.Text(cells[i][1], props.Text{
					Style: consts.Normal,
					Align: consts.Left,
					Size:  14,
				})
			})
			m.Col(3, func() {
				m.Text(cells[i][2], props.Text{
					Style: consts.Normal,
					Align: consts.Left,
					Size:  14,
				})
			})
			m.Col(3, func() {
				m.Text(cells[i][3], props.Text{
					Style: consts.Normal,
					Align: consts.Left,
					Size:  14,
				})
			})
		})

		// }
	}

	// m.Row(40, func() {
	// 	m.Col(3, func() {
	// 		m.Text("Any Text1")
	// 	})
	// 	m.Col(3, func() {
	// 		m.Text("Any Text2")
	// 	})
	// 	m.Col(3, func() {
	// 		m.Text("Any Text3")
	// 	})
	// 	m.Col(3, func() {
	// 		m.Text("Any Text3")
	// 	})
	// })

}

func randomAdd(random *rand.Rand, min int, max int) string {
	a := random.Intn(max) + min
	b := random.Intn(max) + min

	//a + b
	for {
		if a+b <= max {
			break
		}

		if b > a {
			b = random.Intn(max) + min
		} else if b < a {
			a = random.Intn(max) + min
		} else {
			a = random.Intn(max) + min
			b = random.Intn(max) + min
		}
	}

	return render(a, b, "+")
}

func randomSub(random *rand.Rand, min int, max int) string {
	a := random.Intn(max) + min
	b := random.Intn(max) + min

	//a - b
	for {
		if a > b {
			break
		} else if a < b {
			temp := b
			b = a
			a = temp
		} else {
			a = random.Intn(max) + min
			b = random.Intn(max) + min
		}
	}

	return render(a, b, "-")
}

func render(a int, b int, op string) string {
	return fmt.Sprintf("%2d %s %d =", a, op, b)
}

func randomOp(random *rand.Rand) string {

	index := random.Intn(len(ops))

	return ops[index]
}

func getRandom() *rand.Rand {
	now := time.Now().UnixNano()
	seed := rand.NewSource(now)
	random := rand.New(seed)
	// fmt.Print(now, ":")

	return random
}
