
func MergeSort(array []int) []int {
    return mergeSort(array)
}

//归并排序
func mergeSort(srcArray []int) []int {
    //如果数组长度小于2则不需要排序，直接返回
    if len(srcArray) < 2 {
        return srcArray
    }
    
    //分割数组为两份
    split := len(srcArray) / 2
    arr1 := mergeSort(srcArray[:split])
    arr2 := mergeSort(srcArray[split:])
    
    //获取原始数组的总长
    totalLen := len(srcArray)
    return mergeParts(arr1, arr2, totalLen)
}

//归并两块数组
func mergeParts(arr1, arr2 []int, totalLen int) []int {
    left, right := 0, 0
    res := make([]int, totalLen)
    i := 0
    for i < totalLen {
        switch true {
        case left >= len(arr1):
            res[i] = arr2[right]
            right++
        case right >= len(arr2):
            res[i] = arr2[left]
            left++
        default:
            if arr1[left] > arr2[right] {
                
                res[i] = arr2[right]
                right++
                
            } else {
                
                res[i] = arr2[left]
                left++
                
            }
        }
        i++
    }
    return res
}
