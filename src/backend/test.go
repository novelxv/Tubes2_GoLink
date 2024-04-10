package main

import (
	"github.com/angiekierra/Tubes2_GoLink/ids"
	"github.com/angiekierra/Tubes2_GoLink/tree"
)

func main() {
	// Panggil idsfunc untuk membuat pohon
	var test *tree.Tree
	test = ids.Idsfunc("Joko Widodo", "Jusuf Kalla")

	// Cetak struktur pohon menggunakan metode PrintTreeIds
	test.PrintTreeIds()
}
