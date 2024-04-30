package quick

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[i], arr[high] = arr[high], arr[i]
	return i
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}

	p := partition(arr, low, high)
	quickSort(arr, low, p-1)
	quickSort(arr, p+1, high)
}

func Sort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)

    return arr
}
