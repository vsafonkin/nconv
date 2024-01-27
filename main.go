package main

import (
	"fmt"
	"os"
)

func main() {
	var num int
	fmt.Printf("> ")
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("scan error:", err)
		os.Exit(1)
	}

	fmt.Printf("%d %#[1]x %#08[1]b\n", num)
}
