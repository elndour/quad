package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	// largeur > longueur
	if y > x {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')
		for j := 0; j < y-2; j++ {
			z01.PrintRune('|')
			for k := 0; k < x-2; k++ {
				z01.PrintRune(' ')
			}
			// condition pour l'affichage du deuxième "|"
			if x != 1 {
				z01.PrintRune('|')
			}
			z01.PrintRune('\n')
		}
		if y > 1 {
			for i := 0; i < x; i++ {
				if i == 0 || i == x-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
			z01.PrintRune('\n')
		}
	} else {
		for i := 0; i < x; i++ {
			if i == 0 || i == x-1 {
				z01.PrintRune('o')
			} else {
				z01.PrintRune('-')
			}
		}
		z01.PrintRune('\n')
		for j := 0; j < y-2; j++ {
			z01.PrintRune('|')
			for k := 0; k < x-2; k++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('|')
			z01.PrintRune('\n')
		}
		if y > 1 {
			for i := 0; i < x; i++ {
				if i == 0 || i == x-1 {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			}
			z01.PrintRune('\n')
		}
	}
}
