package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
				flowerbed[i] = 1
				n--
			}
		}
	}
	return n <= 0
}

func main() {
	println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1)) // true
	println(canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2)) // false
	println(canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1)) // true
}