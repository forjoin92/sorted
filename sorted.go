package sorted

// BubbleSort bubble sort
func BubbleSort(array []int) {
	l := len(array)

	if l <= 1 {
		return
	}

	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

// SelectionSort selection sort
func SelectionSort(array []int) {
	l := len(array)

	if l <= 1 {
		return
	}

	for i := 0; i < l; i++ {
		temp := i

		for j := i; j < l; j++ {
			if array[j] < array[temp] {
				temp = j
			}
		}
		array[i], array[temp] = array[temp], array[i]
	}
}
