package removeduplicates

func RemoveDuplicates(nums []int) []int {
	uniqueElements := make([]int, 0, len(nums))

	for _, num := range nums {
		found := false
		for _, uniqueNum := range uniqueElements {
			if uniqueNum == num {
				found = true
			}
		}

		if !found {
			uniqueElements = append(uniqueElements, num)
		}
	}
	return uniqueElements
}
