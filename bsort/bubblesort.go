package bsort

func BubbleSort(values []int) (sorted []int) {
	var maxKey int = len(values)
	for i := 0; i < maxKey-1; i++ {
		for j := 0; j < maxKey-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	return
}
