package pdf

import (
	"fmt"
	"os"
	"time"

	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/pdf"
	"github.com/johnfercher/maroto/pkg/props"
)

func New() pdf.Maroto {
	m := pdf.NewMaroto(consts.Portrait, consts.A4)
	m.SetPageMargins(10, 10, 10)
	m.SetBorder(false)
	m.AddUTF8Font("CustomArial", consts.Normal, "res/WenQuan.ttf")
	m.SetFirstPageNb(1)
	return m
}

func SetHeader(m pdf.Maroto, row int, col int) {

	m.RegisterHeader(func() {
		m.Row(20, func() {
			m.Col(12, func() {
				m.Text("数 学 运 算 练 习 v2", props.Text{
					Family: "CustomArial",
					Style:  consts.Normal,
					Align:  consts.Center,
					Size:   20,
				})
			})
		})
		m.Row(16, func() {
			m.Col(2, func() {
				m.Text(fmt.Sprintf("题目数量:%3d", row*col), props.Text{
					Family: "CustomArial",
					Style:  consts.Normal,
					Align:  consts.Left,
					Size:   12,
				})
			})
			m.Col(4, func() {
				m.Text("日期:_______年____月____日", props.Text{
					Family: "CustomArial",
					Style:  consts.Normal,
					Align:  consts.Left,
					Size:   12,
				})
			})
			m.Col(3, func() {
				m.Text("用时:________", props.Text{
					Family: "CustomArial",
					Style:  consts.Normal,
					Align:  consts.Left,
					Size:   12,
				})
			})
			m.Col(3, func() {
				m.Text("得分:________", props.Text{
					Family: "CustomArial",
					Style:  consts.Normal,
					Align:  consts.Left,
					Size:   12,
				})
			})
		})
	})
}

func SetFooter(m pdf.Maroto) {
	m.RegisterFooter(func() {
		m.Row(5, func() {
			m.Col(12, func() {
				m.Text(fmt.Sprintf("%d", m.GetCurrentPage()), props.Text{
					Style: consts.Normal,
					Align: consts.Center,
					Size:  14,
				})
			})
		})
	})
}

func BuildPage(m pdf.Maroto, cells [][]string) {

	row_count := len(cells)
	// col_count := len(cells[0])

	// _, m_top, _, m_bottom := m.GetPageMargins()

	// row_height := (m.GetCurrentOffset() - m_bottom - m_top) / float64(row_count)

	// fmt.Println(row_height)

	for i := 0; i < row_count; i++ {
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
	}
}

func Export(m pdf.Maroto, path string) {
	os.Mkdir(path, os.ModePerm)
	file := fmt.Sprintf("%s/%s.pdf", path, time.Now().Format("2006-01-02"))

	err := m.OutputFileAndClose(file)
	if err != nil {
		fmt.Println("Could not save PDF:", err)
		os.Exit(1)
	}
}
