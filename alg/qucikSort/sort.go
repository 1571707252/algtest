package qucikSort

// python
// 对数列进行排序，所有比基准值小的元素都放在基准值前面，
//所有比基准值大的都放在后面。
// def quick_sort(values):
// if len(values) <= 1:
// return values
// priv = values[0]
// less = [i for i int values[1:] if i < priv]
// greater = [i for i in values[1:] if i > priv]
// return quick_sor(less) + [priv] + quick_sort(greater)

func QuickSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	mid := values[0] // 基准值
	head, tail := 0, len(values)-1

	// 遍历切片
	for i := 1; i <= tail; {
		if values[i] > mid {
			// 小的数值放基准值左侧
			values[i], values[tail] = values[tail], values[i]
			i--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	QuickSort(values[:head])   //对数组左边区的元素进行递归排序
	QuickSort(values[head+1:]) // 对数组右边区的元素进行递归排序
	return values
}
