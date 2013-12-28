// Steve Phillips / elimisteve
// 2013.12.28

package jam

func QuicksortInts(nums []int, ascend bool) []int {
	if len(nums) < 2 {
		return nums
	}
	left, right := split(nums[0], nums[1:], ascend)
	sorted := append(QuicksortInts(left, ascend), nums[0])
	return append(sorted, QuicksortInts(right, ascend)...)
}

func split(head int, rest []int, ascend bool) (left, right []int) {
	for i := 0; i < len(rest); i++ {
		if ascend {
			// Small to big
			if rest[i] <= head {
				left = append(left, rest[i])
			} else {
				right = append(right, rest[i])
			}
		} else {
			// Big to small
			if rest[i] >= head {
				left = append([]int{rest[i]}, left...)
			} else {
				right = append([]int{rest[i]}, right...)
			}
		}
	}
	return
}
