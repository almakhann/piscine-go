package piscine

func ConcatParams(args []string) string {

	s := ""
	l := 0
	for range args {
		l++
	}
	l = l - 1
	for i := range args {
		s = s + args[i]
		if i < l {
			s = s + "\n"
		}
	}
	return s
}
