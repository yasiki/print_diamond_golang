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

	var alphabetList = createAlphabetList()
	var index = searchIndex(alphabetList, alphabet)
	var width = getWidth(index)

	var outputLines []string

	for i := 0; i < index+1; i++ {
		outputLines = append(outputLines, createLine(alphabetList[i], width, i+1))
	}

	for i := len(outputLines) - 2; i >= 0; i-- {
		outputLines = append(outputLines, outputLines[i])
	}

	return strings.Join(outputLines, "")

}

//リスト内のindexを調べる
func searchIndex(list []string, searchValue string) int {
	for index, value := range list {
		if value == searchValue {
			return index
		}
	}
	return -1
}

//diamondの横幅を求める
func getWidth(index int) int {
	return (index+1)*2 - 1
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

//AからZまでの配列を作成する
func createAlphabetList() []string {
	var list []string

	const AsciiA = 65
	const AlphabetCount = 26

	for i := AsciiA; i < AsciiA+AlphabetCount; i++ {
		list = append(list, string(i))
	}

	return list
}
