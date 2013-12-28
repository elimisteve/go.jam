// Steve Phillips / elimisteve
// 2013.12.27

package jam

import "fmt"

func GetInt() int {
	var num int
	fmt.Scanf("%d", &num)
	return num
}

func GetInts(howMany int) []int {
	if howMany <= 0 {
		return nil
	}
	nums := make([]int, howMany)
	for i := 0; i < howMany; i++ {
		fmt.Scan(&nums[i])
	}
	return nums
}
