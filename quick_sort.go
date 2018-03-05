// 快排序
//
package sort2

// 分区
// 将比pivot小的左移，比pivot大的右移。
// 返回pivot的位置
//
func partition(data []int, left, right int) int {
	// 交换数据的位置
	store := left
	for j := left; j < right; j++ {
		// pivot为right位置的值
		if data[j] <= data[right] {
			// 交换数据
			if store != j {
				data[store], data[j] = data[j], data[store]
			}
			store++
		}
	}
	// 将基准元素放置到最后的正确位置上
	data[store], data[right] = data[right], data[store]
	return store
}

// 递归调用
//
func quickSort(data []int, left, right int) {
	if left < right {
		idx := partition(data, left, right)
		quickSort(data, left, idx-1)
		quickSort(data, idx+1, right)
	}
}

// 对data进行正向排序
// 1、在数据集之中，选择中间元素作为基准（pivot）。
// 2、所有小于基准的元素，都移到基准的左边；所有大于基准的元素，都移到基准的右边。
// 这个操作称为分区 (partition) 操作，分区操作结束后，基准元素所处的位置就是最终排序后它的位置。
// 3、对基准左边和右边的两个子集，不断重复第1步和第2步，直到所有子集只剩下一个元素为止。
// 难度系数 O(n^2)
// 不稳定排序
//
func QuickSort(data []int) {
	quickSort(data, 0, len(data)-1)
}
