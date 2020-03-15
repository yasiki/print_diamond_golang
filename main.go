package main

import (
	"fmt"
	"strings"
)

func main() {
	var alphabet string
	_, err := fmt.Scan(&alphabet)
	if err != nil {
		fmt.Printf("入力でエラーが発生しました。処理を終了します。")
		return
	}
	fmt.Printf(PrintDiamond(alphabet))
}

func PrintDiamond(alphabet string) string {

	if alphabet == "A" {
		return createLine("A", 1, 1)
	}

	if alphabet == "B" {
		return createLine("A", 3, 1) + createLine("B", 3, 2) + createLine("A", 3, 1)
	}

	if alphabet == "C" {
		return createLine("A", 5, 1) + createLine("B", 5, 2) + createLine("C", 5, 3) + createLine("B", 5, 2) + createLine("A", 5, 1)
	}

	return alphabet
}

func createLine(alphabet string, width int, lineNo int) string {
	if width%2 == 0 {
		return "" //error
	}

	var outerSpaceCount = (width - (lineNo*2 - 1)) / 2

	if lineNo == 1 {
		return strings.Repeat(" ", outerSpaceCount) + alphabet + strings.Repeat(" ", outerSpaceCount) + "\n"
	}

	var innerSpaceCount = (lineNo-1)*2 - 1
	return strings.Repeat(" ", outerSpaceCount) + alphabet + strings.Repeat(" ", innerSpaceCount) + alphabet + strings.Repeat(" ", outerSpaceCount) + "\n"

}
