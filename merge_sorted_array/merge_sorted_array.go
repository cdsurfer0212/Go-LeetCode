package mergeSortedArray

func merge(nums1 []int, m int, nums2 []int, n int) {
	current := len(nums1) - 1
	i := m - 1
	j := n - 1
	for i >= 0 || j >= 0 {
		if j < 0 || (i >= 0 && nums1[i] >= nums2[j]) {
			nums1[current] = nums1[i]
			i--
		} else {
			nums1[current] = nums2[j]
			j--
		}
		current--
	}
}
