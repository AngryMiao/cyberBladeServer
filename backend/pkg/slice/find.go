package slice

func FindInt(s []int, item int) bool {
	for _, i := range s {
		if i == item {
			return true
		}
	}
	return false
}
