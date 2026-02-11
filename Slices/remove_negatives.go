package slices

// RemoveNegatives entfernt alle negativen Zahlen aus dem Slice,
// indem es das Slice direkt über den übergebenen Pointer verändert.
func RemoveNegatives(nums []int) []int {
	//TODO
	temp := nums[:0]
	for _, num := range nums {
		if num >= 0 {
			temp = append(temp, num)
		}
	}
	return temp
}
