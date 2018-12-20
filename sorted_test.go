package sorted

import "testing"

func TestBubbleSort(t *testing.T) {
	var array = []int{8, 10, 4, 3, 1, 9, 6, 2, 7, 5}

	BubbleSort(array)
	t.Log(array)
}
