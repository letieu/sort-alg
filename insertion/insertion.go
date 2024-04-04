package insertion

func Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		// Move elements of arr[0...i-1], that are greater than key, to one position ahead of their current position.
		// This loop will keep shifting elements to the right until it finds the correct position for the key.
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		// Place the key in its correct location in the sorted sequence.
		arr[j+1] = key
	}

	return arr
}
