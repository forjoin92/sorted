package sorted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	var array = []int{8, 10, 4, 3, 1, 9, 6, 2, 7, 5}

	BubbleSort(array)
	assert.Equal(t, array, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestSelectionSort(t *testing.T) {
	var array = []int{8, 10, 4, 3, 1, 9, 6, 2, 7, 5}

	SelectionSort(array)
	assert.Equal(t, array, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}

func TestInsertionSort(t *testing.T) {
	var array = []int{8, 10, 4, 3, 1, 9, 6, 2, 7, 5}

	InsertionSort(array)
	assert.Equal(t, array, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
