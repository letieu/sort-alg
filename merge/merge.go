package merge

func Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// split
	mid := len(arr) / 2

	left := Sort(arr[:mid])
	right := Sort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int

	// while remain 2 side
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	// if left remain
	if len(left) > 0 {
		result = append(result, left...)
	}

	// if right remain
	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
