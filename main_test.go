package main

import "testing"

func TestInputA(t *testing.T) {
	actual := PrintDiamond("A")
	if actual != "A\n" {
		t.Fatal("Error return string")
	}
}

func TestInputB(t *testing.T) {
	actual := PrintDiamond("B")
	if actual != " A \nB B\n A \n" {
		t.Fatal("Error return string")
	}
}

func TestInputC(t *testing.T) {
	actual := PrintDiamond("C")
	if actual != "  A  \n B B \nC   C\n B B \n  A  \n" {
		t.Fatal("Error return string")
	}
}
