package qsort

func Qsort(arr_Array []int) {
	var i_ArrLen = len(arr_Array)
	if i_ArrLen <= 1 {
		return
	}
	var i_Mid, i_Index int = arr_Array[0], 1 //定义数组第一个元素为中间值    中间值位置键值
	var i_Start, i_End int = 0, i_ArrLen - 1 //数组起始键值    数组结束键值
	for i_Start < i_End {
		//循环ing
		if arr_Array[i_Index] > i_Mid {
			//值比中间值大就放到最后一位
			arr_Array[i_Index], arr_Array[i_End] = arr_Array[i_End], arr_Array[i_Index]
			i_End-- //排除最后一位
		} else {
			//比中间值小就将两个对调位置
			arr_Array[i_Index], arr_Array[i_Start] = arr_Array[i_Start], arr_Array[i_Index]
			i_Start++ //排除第一位
			i_Index++ //中间键的键值对应发生改变+=1
		}
	}
	Qsort(arr_Array[:i_Start])   //<=中间值的数组前部份。
	Qsort(arr_Array[i_Start+1:]) //>中间值的数组后部分。
}
