package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	l := []rune{}

	len_s := 1

	div := n / 10
	mod := n % 10
	l = append(l, rune(mod))

	for div > 0 {
		mod = div % 10
		l = append(l, rune(mod))
		div = div / 10
		len_s = len_s + 1
	}

	for {
		if len_s == 1 {
			break
		}
		for i := 0; i < len_s-1; i++ {
			if l[i] > l[i+1] {
				l[i], l[i+1] = l[i+1], l[i]
			}
		}
		len_s = len_s - 1
	}

	for _, digit := range l {
		z01.PrintRune(digit + 48)

	}

}
