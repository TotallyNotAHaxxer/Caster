package CastHunter

import (
	"fmt"
	"strings"
)

type TableSeperatorCol struct {
	HeaderCrossOriginL       string // This will be the far left cross origin header frame meet
	HeaderCrossOriginL_Color string // This will be the far left cross origin header frame meet
	HeaderCrossOriginR       string // This will be the far right cross origin header frame meet
	HeaderCrossOriginR_Color string // This will be the far right cross origin header frame meet
	CrossLineY               string // This will be the far y cross line meet
	CrossLineY_Color         string // This will be the far y cross line meet color
	CrossLineX               string // This will be the far x cross line meet
	CrossLineX_Color         string // This will be the far x cross line meet color
	ColumnTitleColor         string // This will be the column title color
	RowDataColor             string // This will be the row data color
}

var (
	T TableSeperatorCol
)

func DrawTableSepColBased(rows [][]string, cols []string) {
	fmt.Print("\n\n")
	T.HeaderCrossOriginL_Color = "\033[38;5;57m"
	T.HeaderCrossOriginL = "┣"
	T.HeaderCrossOriginR = "┫"
	T.HeaderCrossOriginR_Color = "\033[38;5;57m"
	T.CrossLineY = "┃"
	T.CrossLineY_Color = "\033[38;5;57m"
	T.CrossLineX_Color = "\033[38;5;57m"
	T.ColumnTitleColor = "\033[38;5;255m"
	T.CrossLineX = "━"
	// Getting column width based on the len of the columns
	colwidth := make([]int, len(cols))
	for o, col := range cols {
		colwidth[o] = len(col)
		for _, rowdata := range rows {
			if len(rowdata[o]) > colwidth[o] {
				colwidth[o] = len(rowdata[o])
			}
		}
	}
	// Generate and calculate header
	headsep := T.HeaderCrossOriginL_Color + T.HeaderCrossOriginL
	for _, w := range colwidth {
		headsep += strings.Repeat(T.CrossLineX_Color+T.CrossLineX, w+2) + T.HeaderCrossOriginR_Color + T.HeaderCrossOriginR
	}
	head := T.CrossLineY_Color + T.CrossLineY
	for i, col1 := range cols {
		head += " " + T.ColumnTitleColor + col1 + strings.Repeat(" ", colwidth[i]-len(col1)) + " " + T.CrossLineY_Color + T.CrossLineY
	}
	// Generate and calculate row data
	Rowdata := make([]string, len(rows))
	for k, row := range rows {
		RowT := T.CrossLineY_Color + T.CrossLineY
		for l, col := range row {
			RowT += " " + T.ColumnTitleColor + col + strings.Repeat(" ", colwidth[l]-len(col)) + " " + T.CrossLineY_Color + T.CrossLineY
		}
		Rowdata[k] = RowT
	}
	fmt.Println(headsep)
	fmt.Println(head)
	fmt.Println(headsep)
	for _, rt := range Rowdata {
		fmt.Println(rt)
	}
	fmt.Println(headsep)
}
