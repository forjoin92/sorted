package sorted

import "fmt"

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

// InsertionSort insertion sort
func InsertionSort(array []int) {
	l := len(array)

	if l <= 1 {
		return
	}

	for i := 0; i < l; i++ {
		for j := i; j > 0; j-- {
			if array[j] >= array[j-1] {
				break
			}
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}

// QuickSort quick sort
func QuickSort(array []int) {
	l := len(array)

	if l <= 1 {
		return
	}

	pivot, i := 0, 1
	head, tail := 0, l-1

	for head < tail {
		fmt.Println(array, "head:", array[head], "i:", array[i], "tail:", array[tail])
		if array[i] > array[pivot] {
			array[i], array[tail] = array[tail], array[i]
			tail--
		} else {
			array[i], array[head] = array[head], array[i]
			head++
			i++
		}
	}
	//array[head] = pivot
	QuickSort(array[:head])
	QuickSort(array[head+1:])
}
