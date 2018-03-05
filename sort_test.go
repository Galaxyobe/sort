package sort2_test

import (
	"testing"
	"fmt"
	"sort2"
)

func TestBubbleSort(t *testing.T) {
	data := []int{1, 0, 35, 9, 7, 8, 5, 4, 3, 33}
	fmt.Println("sort before", data)
	sort2.BubbleSort(data)
	fmt.Println("sort after", data)
}

func TestInsertSort(t *testing.T) {
	data := []int{1, 0, 35, 9, 7, 8, 5, 4, 3, 33}
	fmt.Println("sort before", data)
	sort2.InsertSort(data)
	fmt.Println("sort after", data)
}

func TestQuickSort(t *testing.T) {
	data := []int{0, 2, 4, 6, 8, 10, 9, 7, 5, 3, 1,8}
	fmt.Println("sort before", data)
	sort2.QuickSort(data)
	fmt.Println("sort after", data)
}
