package piscine

import "github.com/01-edu/z01"

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for w := 0; w <= y-1; w++ {
			for l := 0; l <= x-1; l++ {
				if l == 0 || l == x-1 || w == 0 || w == y-1 {
					if l == 0 && w == 0 { // mapping each index for the x
						z01.PrintRune('/')
					} else if w == 0 && l == x-1 {
						z01.PrintRune('\\')
					} else if l == 0 && w == y-1 {
						z01.PrintRune('\\')
					} else if l == x-1 && w == y-1 {
						z01.PrintRune('/')
					} else if w == 0 || w == y-1 {
						z01.PrintRune('*')
					} else if l == 0 || l == x-1 {
						z01.PrintRune('*')
					} else {
						// z01.PrintRune("")
					}
				} else {
					z01.PrintRune(' ')
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		z01.PrintRune(' ')
	}
}
