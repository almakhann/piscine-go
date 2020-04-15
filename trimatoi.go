package piscine

func TrimAtoi(s string) int {
	count := 0
	count_len := 0
	var signs int

	var q string
	for _, letter := range s {
		if ('0' <= letter && '9' >= letter) || letter == '-' {
			count++
			if letter == '-' {
				signs = count
			} else {
				count_len++
				q = q + string(letter)
			}
		}
	}

	last := 0

	for _, nbr := range q {
		counter := 0
		for i := '0'; i <= '9'; i++ {
			if i == nbr {
				break
			} else {
				counter++
			}
		}
		last = last + counter*Power10(count_len-1)
		count_len--
	}

	if signs == 1 {
		last = last * (-1)
	}

	return last
}
