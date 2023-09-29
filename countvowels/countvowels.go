package countvowels

func CountVowels(s string) int {
	vowels := []rune{'A', 'E', 'I', 'O', 'U', 'Y', 'a', 'e', 'i', 'o', 'u', 'y'}
	count := 0

	for _, char := range s {
		for _, vowels := range vowels {
			if char == vowels {
				count++
				break
			}
		}
	}
	return count
}
