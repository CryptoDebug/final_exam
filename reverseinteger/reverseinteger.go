package reverseinteger

func ReverseInteger(n int) int {
	reversed := 0
	if n < 0 {
		n = -n
		reversed = -reversed
	}
	for n > 0 {
		reversed = reversed*10 + n%10
		n = n / 10
	}
	return reversed
}
