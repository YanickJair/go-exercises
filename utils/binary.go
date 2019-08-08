package utils

// BinarySearch - impliment the binary search algorithm
func BinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1
	}

	mid := int((lowIndex + highIndex) / 2)
	if array[mid] > target {
		return BinarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return BinarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}
