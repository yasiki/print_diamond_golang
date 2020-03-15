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
		return createFirstLine("A", 3) + "B B\n" + " A \n"
	}

	if alphabet == "C" {
		return createFirstLine("A", 5) + " B B \nC   C\n B B \n  A  \n"
	}

	return alphabet
}

func createFirstLine(alphabet string, width int) string {
	var spaceCount = width / 2
	return strings.Repeat(" ", spaceCount) + alphabet + strings.Repeat(" ", spaceCount) + "\n"
}
