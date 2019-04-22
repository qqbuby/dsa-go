package sort

// 冒泡排序：重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来，直到再没有交换
//	 比较操作 O(n*2)
func BubbleSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := 0; j < len(a)-i; j++ {
			if a[j] > a[j+1] { // compare
				a[j], a[j+1] = a[j+1], a[j] // swap
			}
		}
	}
}

// 选择排序: 在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾
//	 比较操作 n(n-1)/2   = O(n*2)
//	 交换操作 0 ~ n-1
//	 赋值操作 0 ~ 3(n-1)
func SelectSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		min := i
		for j := i + 1; j < len(a); j++ {
			if a[j] < a[min] { // compare
				min = j
			}
		}
		a[min], a[i] = a[i], a[min] // swap
	}
}

// 插入排序: 通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
// 	比较操作: n-1 ~ n(n-1)/2
func InsertSort(a []int) {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0; j-- {
			if a[j-1] > a[j] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
}

//分治法：
//  分解: 原问题为若干子问题，这些子问题是原有问题的规模较小的实例
//  解决: 这些子问题，递归地求解各子问题。然而，若子问题的规模足够小，则直接求解
//  合并: 这些子问题的解成原问题的解
//
//归并排序:
//    分解: 分解待排序的n个元素的序列各具 n/2 个元素的两个子序列
//    解决: 使用递归排序递归地排序两个子序列
//    合并: 合并两个子序列以产生已排序的答案
func MergeSort(a []int, low, high int) {
	if high-low > 1 {
		mid := (low + high) / 2
		MergeSort(a, low, mid)
		MergeSort(a, mid, high)
		merge(a, low, mid, high)
	}
}

func merge(a []int, low, mid, high int) {
	left := make([]int, mid-low)
	copy(left, a[low:mid])
	right := make([]int, high-mid)
	copy(right, a[mid:high])

	l := 0
	r := 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			a[low] = left[l]
			l++
		} else {
			a[low] = right[r]
			r++
		}
		low++
	}

	if l == len(left) {
		for r < len(right) {
			a[low] = right[r]
			low++
			r++
		}
	} else {
		for l < len(left) {
			a[low] = left[l]
			low++
			l++
		}
	}
}

// 快速排序：
//   分解: 选择主元 P(viot), 将待排序序列分割成两个区域(Partition)，左边的分区的元素都小于或等于P，右边的元素大于P
//   解决: 对左右两个分区进行递归的快速排序
//   合并: 由于序列是原址排序，分区的操作即为排序的操作，无需合并
func QuickSort(a []int, low, high int) {
	if high-low > 1 {
		mid := partition(a, low, high)
		QuickSort(a, low, mid)
		QuickSort(a, mid, high)
	}
}

func partition(a []int, low, high int) int {
	pivot := a[high-1]

	i := low - 1
	for j := low; j < high-1; j++ {
		if a[j] <= pivot {
			i++
			a[j], a[i] = a[i], a[j]
		}
	}

	a[i+1], a[high-1] = a[high-1], a[i+1]
	return i + 1
}
