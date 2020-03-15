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
		return createFirstLine("A", 1)
	}

	if alphabet == "B" {
		return createFirstLine("A", 3) + createLine("B", 3, 2) + createFirstLine("A", 3)
	}

	if alphabet == "C" {
		return createFirstLine("A", 5) + createLine("B", 5, 2) + createLine("C", 5, 3) + createLine("B", 5, 2) + createFirstLine("A", 5)
	}

	return alphabet
}

func createFirstLine(alphabet string, width int) string {
	var spaceCount = width / 2
	return strings.Repeat(" ", spaceCount) + alphabet + strings.Repeat(" ", spaceCount) + "\n"
}

func createLine(alphabet string, width int, lineNo int) string {
	if width < 3 || width%2 == 0 {
		return "" //error
	}
	var innerSpaceCount = (lineNo-1)*2 - 1
	var outerSpaceCount = (width - (lineNo*2 - 1)) / 2
	return strings.Repeat(" ", outerSpaceCount) + alphabet + strings.Repeat(" ", innerSpaceCount) + alphabet + strings.Repeat(" ", outerSpaceCount) + "\n"

}
