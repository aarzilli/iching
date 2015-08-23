package main

import "fmt"

func main() {
	for ch := rune(0x4dc0); ch <= rune(0x4dff); ch++ {
		fmt.Printf("wget -c http://www.fileformat.info/info/unicode/char/%x/%x.png\n", ch, ch)
	}
}
