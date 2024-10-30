package search

func LinearSearch(arr []int, val int) bool {

	for i := range arr {
		if arr[i] == val {
			return true
		}

	}

	return false
}
