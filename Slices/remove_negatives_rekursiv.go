package slices

// RemoveNegatives entfernt alle negativen Zahlen aus dem Slice,
// indem es das Slice direkt über den übergebenen Pointer verändert.
func RemoveNegativesRekursiv(nums []int) []int {
	//TODO
	temp := make([]int, 0)
	count := remove(nums, 0, &temp)
	return temp[:count]
}

func remove(nums []int, count int, temp *[]int) int {
	//TODO
	if len(nums) == 0 {
		return count
	}
	if nums[0] >= 0 {
		*temp = append(*temp, nums[0])
		count++
	}
	return remove(nums[1:], count, temp)
}
