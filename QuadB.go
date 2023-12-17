package main

import "fmt"

func QuadB(width, height int) {
	if width > 0 && height > 0 {
		for h := 0; h <= height-1; h++ {
			for w := 0; w <= width-1; w++ {
				if w == 0 || w == width-1 || h == 0 || h == height-1 {
					if w == 0 && h == 0 { // mapping each index for the width
						fmt.Print("/")
					} else if w == 0 && h == height-1 {
						fmt.Print('\')
					} else if w == width-1 && h == height-1 {
						fmt.Print("/")
					} else if w == width-1 && h == 0 {
						fmt.Print('\')
					} else if h == 0 || h == height-1 {
						fmt.Print("*")
					} else if w == 0 || w == width-1 {
						fmt.Print("*")
					} else {
						// fmt.Print("")
					}
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Print("\n")
		}
	} else {
		fmt.Print("")
	}
}

func main(){
	Quad(5,3)
}