// 冒泡排序
//
package sort2

// 对data进行正向排序
// 每一趟都把最大的元素移到最后。
// 1、比较相邻的2个元素，如果前一个比后一个大，就交换2个元素；
// 2、对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对；
// 3、针对所有的元素重复以上的步骤，除了最后一个；
// 4、持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
// 设置swap交换标志位可提供效率。
// 难度系数 O(n^2)
//
func BubbleSort(data []int) {
	// 交换标志
	swap := false
	for i := 0; i < len(data)-1; i++ {
		swap = false
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
				swap = true
			}
		}
		if !swap {
			break
		}
	}
}
