package main

func BinarySearch(arr []int, target int) int {
	l, r := 0, len(arr)-1
	for l <= r {
		mid := l + (r-l)/2
		if arr[mid] == target {
			return mid
		}
		if arr[mid] > target {
			r = mid - 1
		}
		if arr[mid] < target {
			l = mid + 1
		}
	}
	return -1
}
