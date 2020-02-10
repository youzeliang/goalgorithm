package _605_can_place_flowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	flowerbed = append([]int{0}, append(flowerbed, 0)...)
	l := len(flowerbed)
	for i := 1; i < l-1; i++ {
		if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n--
		}
	}
	return n <= 0
}