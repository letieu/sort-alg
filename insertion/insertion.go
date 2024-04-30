package insertion

func Sort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
        j := i - 1
        key := arr[i]

        for j >= 0 && key < arr[j] {
            arr[j + 1] = arr[j]
            j--
        }

        arr[j + 1] = key
	}

    return arr
}
