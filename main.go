package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var interactive *bool = flag.Bool("i", false, "interactive")

func main() {
	flag.Parse()

	for *interactive {
		cycleConv()
	}

	if len(os.Args) > 1 {
		for i := 1; i < len(os.Args); i++ {
			num, err := strconv.Atoi(os.Args[i])
			if err != nil {
				fmt.Println("scan error:", err)
				continue
			}
			output(num)
		}
		os.Exit(0)
	}
	usage()
}

func cycleConv() {
	var num int
	fmt.Printf("> ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("scan error:", err)
	}
	if num == -1 {
		os.Exit(0)
	}
	output(num)
}

func output(num int) {
	fmt.Printf("%d %#[1]x %#08[1]b\n", num)
}

func usage() {
	fmt.Println(`Usage of nconv:
	nconv <number>
or interactive:
	nconv -i`)
}
