package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		res := nb + 1
		for IsPrime(res) == false {
			res++
		}
		return res
	}
}
