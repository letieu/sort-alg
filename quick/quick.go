package quick

func Sort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	// work with pivot
	pivot := len(arr) - 1
	j := -1

	for i := 0; i < len(arr); i++ {
		if arr[i] > arr[pivot] {
			continue
		}

		j++
		arr[i], arr[j] = arr[j], arr[i]
	}

	// split
	left := Sort(arr[:j])
	right := Sort(arr[j+1:])

	left = append(left, arr[j])
	left = append(left, right...)

	return left
}
