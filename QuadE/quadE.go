package piscine

import "github.com/01-edu/z01"

func QuadE(x, y int) {
	if y > x {
		// y > x et x==1
		if x == 1 {
			z01.PrintRune('A')
			z01.PrintRune('\n')
			for j := 0; j < y-2; j++ {
				z01.PrintRune('B')
				z01.PrintRune('\n')
			}
			z01.PrintRune('C')
			z01.PrintRune('\n')
		} else {
			// y > x et x !=1
			for i := 0; i < x; i++ {
				if i == 0 {
					z01.PrintRune('A')
				} else if i == x-1 {
					z01.PrintRune('C')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
			for j := 0; j < y-2; j++ {
				z01.PrintRune('B')
				for k := 0; k < x-2; k++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune('B')
				z01.PrintRune('\n')
			}
			// affichage sur une deuxième ligne
			if y > 1 {
				for i := 0; i < x; i++ {
					if i == 0 {
						z01.PrintRune('C')
					} else if i == x-1 {
						z01.PrintRune('A')
					} else {
						z01.PrintRune('B')
					}
				}
				z01.PrintRune('\n')
			}
		}
		// x > y
	} else {
		for i := 0; i < x; i++ {
			if i == 0 {
				z01.PrintRune('A')
			} else if i == x-1 {
				z01.PrintRune('C')
			} else {
				z01.PrintRune('B')
			}
		}
		z01.PrintRune('\n')
		for j := 0; j < y-2; j++ {
			z01.PrintRune('B')
			for k := 0; k < x-2; k++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('B')
			z01.PrintRune('\n')
		} // AFFICHAGE DE LE DEUXIÉME LIGNE
		if y > 1 {
			for i := 0; i < x; i++ {
				if i == 0 {
					z01.PrintRune('C')
				} else if i == x-1 {
					z01.PrintRune('A')
				} else {
					z01.PrintRune('B')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
