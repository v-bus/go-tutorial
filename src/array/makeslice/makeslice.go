package array

// MakeSlice returns []int with length and capacity
func MakeSlice(length, capacity byte) (output []int) {
	output = make([]int, length, capacity)
	return
}
