package sorting

func QuickSort(arr []int, low, high int) []int {

	if low < high {
		pi := Partition(arr, low, high)

		QuickSort(arr, low, pi-1)
		QuickSort(arr, pi+1, high)
	}

	return arr
}

func Partition(arr []int, low, high int) int {

	pivot := arr[high]
	i := low
	j := high
	if i < j {
		for arr[i] <= pivot && i < j {
			i++
		}
		for arr[j] >= pivot && i < j {
			j--
		}

		Swap(arr, i, j)

	}
	Swap(arr, j, high)

	return j

}

func Swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

// func Partition(arr []int, low, high int) int {

// 	pivort := arr[high]
// 	j := low - 1
// 	for i := low; i < high; i++ {
// 		if arr[i] < pivort {
// 			j++
// 			arr[j], arr[i] = arr[i], arr[j]
// 		}

// 	}

// 	arr[j+1], arr[high] = arr[high], arr[j+1]
// 	return j + 1
// }
