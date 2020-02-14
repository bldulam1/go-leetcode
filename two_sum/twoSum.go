package two_sum

func TwoSum(nums []int, target int) []int {
	recorded := make(map[int]int)
	pair := make([]int, 2)
	var num_pair int
	for ind, num := range nums {
		num_pair = target - num

		if recorded[num_pair] > 0 {
			pair = []int{
				recorded[num_pair] - 1,
				ind,
			}
		}

		recorded[num] = ind + 1
	}
	return pair
}
