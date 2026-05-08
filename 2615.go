package main

func distance(nums []int) []int64 {
    indexMap := make(map[int][]int, len(nums))
    for i, num := range nums {
        indexMap[num] = append(indexMap[num], i)
    }

    res := make([]int64, len(nums))
    for _, indices := range indexMap {
        var total int64
        for _, idx := range indices {
            total += int64(idx)
        }

        var sumLeft int64
        for j, idx := range indices {
            idx64 := int64(idx)
            left := int64(j)*idx64 - sumLeft
            right := (total - sumLeft - idx64) - int64(len(indices)-j-1)*idx64
            res[idx] = left + right
            sumLeft += idx64
        }
    }

    return res
}
