package sorting

func Bubble(a []int) []int {

	for i := 0; i < len(a)-1; i++ {

		for j := 0; j < len(a)-1-i; j++ {

			if a[j] > a[j+1] {
				temp := a[j+1]
				a[j+1] = a[j]
				a[j] = temp

			}

		}

	}

	return a
}
