package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := '0'; k <= '9'; k++ {
				for q := '0'; q <= '9'; q++ {
					if i <= k && j < q {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')

						z01.PrintRune(k)
						z01.PrintRune(q)

						if !(i == '9' && j == '8' && k == '9' && q == '9') {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					} else if i < k && j >= q {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')

						z01.PrintRune(k)
						z01.PrintRune(q)
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune(10)
}
