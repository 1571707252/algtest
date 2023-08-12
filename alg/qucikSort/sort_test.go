package qucikSort

import "testing"

// 测试普通数组排序
func TestQuickSort(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	QuickSort(values)
	// 判断排序结果
	if values[0] != 1 || values[1] != 2 || values[2] != 3 ||
		values[3] != 4 || values[4] != 5 {
		t.Error("QuickSort() failed . Got",
			values, "Excepted 1, 2, 3, 4,5")
	}
}

// 测试有重复元素的数组排序
func TestQuickSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	QuickSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 ||
		values[3] != 5 || values[4] != 5 {
		t.Error("BubbleSort() failed . Got",
			values, "Excepted 1 2 3 5 5")
	}
}

// 测试只有一个元素的corner case
func TestQuickSort3(t *testing.T) {
	values := []int{5}
	QuickSort(values)
	if values[0] != 5 {
		t.Error("BubbleSort() falied. Got", values,
			"Excepted5")
	}
}
