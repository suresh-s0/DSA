package search

func BinarySearch(arr []int, val int) bool {

	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] == val {
			return true
		} else if arr[mid] < val {
			low = mid + 1

		} else {
			high = mid - 1
		}

	}

	return false
}
