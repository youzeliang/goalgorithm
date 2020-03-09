package _717_1_bit_and_2_bit_characters

func isOneBitCharacter(bits []int) bool {
	n := len(bits)
	i := 0
	for i < n-1 {
		if bits[i] == 1 {
			// 遇见 1 说明遇见了 twoBits
			i += 2
		} else {
			i++
		}
	}
	return i == n-1
}
