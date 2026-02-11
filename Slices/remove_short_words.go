package slices

// RemoveShortWords entfernt alle Wörter aus dem Slice,
// die kürzer sind als die angegebene Mindestlänge minLen.
// Das Slice wird direkt über einen Pointer verändert.
func RemoveShortWords(words []string, minLen int) []string {
	//TODO
	temp := words[:0]
	for _, word := range words {
		if len(word) >= minLen {
			temp = append(temp, word)
		}
	}
	return temp
}
