package main

/**
 * 二分查找算法: 在指定切片(有序)中查找等于给定值的元素索引
 * 递归实现
 * @Param: point []float32 指定的切片
 * @Param: v float32 给定的值
 * @Return: index int 查找符合条件的元素在切片中的索引
 */
func BinarySearchEqualByRecursive(point []float32, v float32) int {
	size := len(point)
	if size == 0 {
		// 切片长为空时返回 -1
		return -1
	}
	return bsEqualByRecursive(point, v, 0, size-1)
}

/**
 * 二分查找算法的递归部分: 在指定切片(有序)中查找等于给定值的元素索引
 * @Param: point []float32 指定的切片
 * @Param: v float32 给定的值
 * @Param: low, high int 查找的指定上下标
 * @Return: index int 查找符合条件的元素在切片中的索引
 */
func bsEqualByRecursive(point []float32, v float32, low, high int) int {
	if low > high {
		// 下标大于上标
		return -1
	}
	mid := (low + high) / 2
	if point[mid] == v {
		return mid
	} else if point[mid] < v {
		return bsEqualByRecursive(point, v, mid+1, high)
	} else {
		return bsEqualByRecursive(point, v, low, mid-1)
	}
}

// BinarySearch 二分查找递归解法
func BinarySearch(array []int, target int, l, r int) int {
	if l > r {
		// 出界了，找不到
		return -1
	}

	// 从中间开始找
	mid := (l + r) / 2
	middleNum := array[mid]

	if middleNum == target {
		return mid // 找到了
	} else if middleNum > target {
		// 中间的数比目标还大，从左边找
		return BinarySearch(array, target, 1, mid-1)
	} else {
		// 中间的数比目标还小，从右边找
		return BinarySearch(array, target, mid+1, r)
	}

}

// BinarySearch2 二分查找非递归解法
func BinarySearch2(array []int, target int, l, r int) int {
	ltemp := l
	rtemp := r

	for {
		if ltemp > rtemp {
			// 出界了，找不到
			return -1
		}

		// 从中间开始找
		mid := (ltemp + rtemp) / 2
		middleNum := array[mid]

		if middleNum == target {
			return mid // 找到了
		} else if middleNum > target {
			// 中间的数比目标还大，从左边找
			rtemp = mid - 1
		} else {
			// 中间的数比目标还小，从右边找
			ltemp = mid + 1
		}
	}
}
