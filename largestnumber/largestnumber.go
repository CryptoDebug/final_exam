package largestnumber

func LargestNumber(nums []int) int {

	if len(nums) == 0 {
		return 0
	}

	largest := nums[0]

	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}
	return largest
}
