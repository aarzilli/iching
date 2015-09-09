package main

import (
	"bufio"
	"fmt"
	"github.com/aarzilli/iching"
	"io"
	"os"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: ichingdump <path to file>\n")
	os.Exit(1)
}

func printLine(addr int, buf []byte, sz int, wr io.Writer, end int) {
	_, err := io.WriteString(wr, iching.SpacePad(iching.Itoiching(uint64(addr)), sz))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on output: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(wr, " ")
	for i := 0; i < len(buf); i += 2 {
		n := uint64(buf[i])<<8 + uint64(buf[i+1])
		if i == 8 {
			io.WriteString(wr, " ")
		}

		if i >= end {
			fmt.Fprintf(wr, "    ")
		} else {
			ich := iching.QianPad(iching.Itoiching(n), 3)
			fmt.Fprintf(wr, " %s", ich)
		}
	}
	io.WriteString(wr, "  |")
	for i := 0; i < len(buf); i++ {
		if i >= end {
			fmt.Fprintf(wr, " ")
		} else {
			if buf[i] >= 0x20 && buf[i] <= 0x7e {
				fmt.Fprintf(wr, "%c", buf[i])
			} else {
				io.WriteString(wr, ".")
			}
		}
	}
	io.WriteString(wr, "|\n")
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Wrong number of arguments\n")
		usage()
	}
	fh, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not open %s: %v\n", os.Args[1], err)
		usage()
	}

	fi, err := fh.Stat()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not stat %s: %v\n", os.Args[1], err)
		usage()
	}

	sz := len([]rune(iching.Itoiching(uint64(fi.Size()))))

	rd := bufio.NewReader(fh)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	buf := make([]byte, 16)
	i := 0
	for {
		n, err := io.ReadFull(rd, buf)

		for i := n; i < len(buf); i++ {
			buf[i] = 0x00
		}

		if n > 0 {
			printLine(i, buf, sz, wr, n)
			i += n
		}

		if err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
				os.Exit(1)
			}
			break
		}
	}
}
