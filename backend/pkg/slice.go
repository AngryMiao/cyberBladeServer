package pkg

func FindInt(s []int, item int) bool {
	for _, i := range s {
		if i == item {
			return true
		}
	}
	return false
}

func FindString(s []string, item string) bool {
	for _, i := range s {
		if i == item {
			return true
		}
	}
	return false
}

func AllStringsEqual(s []string, item string) bool {
	for _, i := range s {
		if i != item {
			return false
		}
	}
	return true
}

func StringsSub(s []string, item []string) bool {
	for _, i := range s {
		if !FindString(item, i) {
			return false
		}
	}
	return true
}

func RemoveDuplicateInt(values []int) []int {
	result := make([]int, 0, len(values))
	temp := map[int]struct{}{}
	for _, item := range values {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func findAndDeleteInt(s []int, itemToDelete int) []int {
	newSlice := make([]int, 0)
	index := 0
	for _, i := range s {
		if i != itemToDelete {
			newSlice[index] = i
			index++
		}
	}
	return newSlice[:index]
}

func FindAndDeleteMultiInt(s []int, d []int) []int {
	newSlice := make([]int, 0)

	for _, i := range s {
		isSave := true
		for _, j := range d {
			if i == j {
				isSave = false
				break
			}
		}
		if isSave {
			newSlice = append(newSlice, i)
		}
	}
	return newSlice
}
