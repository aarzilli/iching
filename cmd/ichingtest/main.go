package main

import (
	"fmt"
	"github.com/aarzilli/iching"
	"os"
	"strconv"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: ichingtest <hex num>|<oct num>|<dec num>|<iching num>\n")
	os.Exit(1)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Wrong number of arguments.\n")
		usage()
	}

	s := os.Args[1]

	var n uint64
	var err error

	ch := ([]rune(s))[0]
	if ch >= 0x4dc0 && ch <= 0x4dff {
		n, err = iching.Ichingtoi(s)
	} else {
		n, err = strconv.ParseUint(s, 0, 64)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Wrong argument: %v\n", err)
		usage()
	}

	fmt.Printf("decimal = %d\n", n)

	oct := fmt.Sprintf("%o", n)

	fmt.Printf("octal   =")
	start := 0
	if len(oct)%2 != 0 {
		fmt.Printf(" 0%c", oct[0])
		start = 1
	}
	for i := start; i < len(oct); i += 2 {
		fmt.Printf(" %c%c", oct[i], oct[i+1])
	}
	fmt.Printf("\n")

	ich := iching.Itoiching(n)

	fmt.Printf("iching  =")
	for _, ch := range []rune(ich) {
		fmt.Printf("  %c", ch)
	}
	fmt.Printf("\n")
}
