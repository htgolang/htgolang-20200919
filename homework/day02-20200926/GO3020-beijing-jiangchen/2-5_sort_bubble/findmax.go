package main

//FindMax ...
func FindMax(slice []int) (index int, result int) {
	index, result = 0, -1
	for idx, v := range slice {
		if result < v {
			index, result = idx, v
		} else {
			continue
		}
	}
	return index, result
}
