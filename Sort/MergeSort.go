package sorting

func MergeSort(arr []int) []int {

	if len(arr) < 2 {
		return arr
	}
	mid := len(arr) / 2

	firstarray := make([]int, mid)
	secoundarray := make([]int, len(arr)-mid)

	for i := 0; i < mid; i++ {
		firstarray[i] = arr[i]
	}

	for i := mid; i < len(arr); i++ {
		secoundarray[i-mid] = arr[i]
	}

	MergeSort(firstarray)
	MergeSort(secoundarray)

	Merge(arr, firstarray, secoundarray)

	return arr
}

func Merge(arr, frstarr, scudarr []int) []int {
	i := 0
	j := 0
	k := 0

	for i < len(frstarr) && j < len(scudarr) {
		if frstarr[i] < scudarr[j] {
			arr[k] = frstarr[i]
			i++
		} else {
			arr[k] = scudarr[j]
			j++
		}
		k++
	}

	for i < len(frstarr) {
		arr[k] = frstarr[i]
		i++
		k++

	}
	for j < len(scudarr) {
		arr[k] = scudarr[j]
		j++
		k++

	}

	return arr
}
