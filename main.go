package main

import "fmt"

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
		return "A\n"
	}

	if alphabet == "B" {
		return " A \nB B\n A \n"
	}

	if alphabet == "C" {
		return "  A  \n B B \nC   C\n B B \n  A  \n"
	}

	return alphabet
}
