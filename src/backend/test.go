package main

import (
	"github.com/angiekierra/Tubes2_GoLink/ids"
	"github.com/angiekierra/Tubes2_GoLink/tree"
)

func main() {
	var test *tree.Tree
	test = ids.Idsfunc("Joko Widodo", "Jusuf Kalla")

	test.PrintTreeIds()
}
