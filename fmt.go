// Steve Phillips / elimisteve
// 2013.12.28

package jam

import "strconv"

// FmtInts returns a string of the form "%d %d %d ...", one %d for
// each number in nums
func FmtInts(nums []int) string {
	if len(nums) == 0 {
		return ""
	}
	s := strconv.Itoa(nums[0])
	for _, n := range nums[1:] {
		s += " " + strconv.Itoa(n)
	}
	return s
}
