package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func tobits(name string) int {
	fh, err := os.Open(name)
	must(err)
	pic, _, err := image.Decode(fh)
	must(err)
	r := 0
	for i := 0; i < 6; i++ {
		col := pic.At(45, 22+10*i)
		_, _, _, a := col.RGBA()
		r = r << 1
		if a == 0 {
			r++
		}
	}
	return r
}

func main() {
	fmt.Printf("package iching\n\n")
	var table = make([]rune, 64)
	fmt.Printf("var invtable = []uint64{\n")
	for ch := rune(0x4dc0); ch <= rune(0x4dff); ch++ {
		val := tobits(fmt.Sprintf("%x.png", ch))
		table[val] = ch
		fmt.Printf("\t%d,\n", val)
	}
	fmt.Printf("}\n\nvar table = []rune{\n")
	for i := range table {
		fmt.Printf("\t0x%x,\n", table[i])
	}
	fmt.Printf("}\n")
}
