package main

import "fmt"

func rotate(nums []int, k int) []int {
	if k < 0 || k > 100 || len(nums) == 0 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		x := nums[i]
		if x < -1000 || x > 1000 {
			return nums
			//fmt.Println("os")
		}
	}
	//fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)

	r := len(nums) - k%len(nums)

	nums = append(nums[r:], nums[:r]...)

	fmt.Println(nums)
	return nums

}

func main() {
	nums := []int{3, 8, 9, 7, 6}

	//fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)

	nums = rotate(nums, 3)

	//fmt.Printf("nums %p array %p len %d cap %d slice %v\n", &nums, &nums[0], len(nums), cap(nums), nums)
}
