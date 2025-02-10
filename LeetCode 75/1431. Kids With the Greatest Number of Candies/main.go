package main

func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	for _, v := range candies {
		if v > max {
			max = v
		}
	}
	result := make([]bool, len(candies))
	for i, v := range candies {
		if v + extraCandies >= max {
			result[i] = true
		} else {
			result[i] = false
		}
	}
	return result
} 
func main() {
	kidsWithCandies([]int{2, 3, 5, 1, 3}, 3) // [true,true,true,false,true]
	kidsWithCandies([]int{4, 2, 1, 1, 2}, 1) // [true,false,false,false,false]
	kidsWithCandies([]int{12, 1, 12}, 10) // [true,false,true]
}