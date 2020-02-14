package medianOfTwoSortedArrays

func MedianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {
	n1Len, n2Len := len(nums1), len(nums2)
	sum := n1Len + n2Len
	medianInd := sum/2 + 1

	ind, i1, i2 := 0, 0, 0
	newList := make([]int, 0)

	for ind < medianInd {
		var smaller int
		if i1 >= n1Len {
			smaller = nums2[i2]
			i2++
		} else if i2 >= n2Len {
			smaller = nums1[i1]
			i1++
		} else {
			if nums1[i1] <= nums2[i2] {
				smaller = nums1[i1]
				i1++
			} else if i2 < n2Len {
				smaller = nums2[i2]
				i2++
			}
		}

		newList = append(newList, smaller)

		ind = i1 + i2
	}

	if sum%2 == 0 {
		return float64(newList[ind-1]+newList[ind-2]) / 2
	}
	return float64(newList[ind-1])
}
