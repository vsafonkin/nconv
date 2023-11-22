package main

import "fmt"

func main() {
	var num int
	for num != -1 {
		fmt.Printf("> ")
		_, err := fmt.Scan(&num)
		if err != nil {
			fmt.Println("scan error:", err)
			continue
		}
		if num == -1 {
			continue
		}
		fmt.Printf("%d %#[1]x %#08[1]b\n", num)
	}
}
