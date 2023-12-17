package main

import "fmt"

func QuadA(x, y int) {
	if x > 0 && y > 0 {
		for w := 0; w <= y-1; w++ {
			if w == 0 {
				for l := 0; l <= x-1; l++ {
					if l == 0 || l == x-1 {
						fmt.Print("o")
					} else if l > 0 || l < x-1 {
						fmt.Print("-")
					} else {
						fmt.Print("\n")
					}
				}
				fmt.Print("\n")
			} else if w > 0 && w < y-1 {
				for l := 0; l <= x-1; l++ {
					if l == 0 || l == x-1 {
						fmt.Print("|")
					} else if l > 0 && l < x-1 {
						fmt.Print(" ")
					} else {
						fmt.Print("\n")
					}
				}
				fmt.Print("\n")
			} else {
				for l := 0; l <= x-1; l++ {
					if l == 0 || l == x-1 {
						fmt.Print("o")
					} else if l > 0 || l < x-1 {
						fmt.Print("-")
					} else {
						fmt.Print("\n")
					}
				}
				fmt.Print("\n")
			}
		}
	} else {
		fmt.Print("")
	}
}

func main() {
	QuadA(5, 3)
	QuadA(5, 1)
}
