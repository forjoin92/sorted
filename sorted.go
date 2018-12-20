package sorted

func BubbleSort(array []int) {
	l := len(array)

	for i := 0; i < l - 1; i++ {
		for j := 0; j < l - 1 - i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
