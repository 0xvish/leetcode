package main

func main() {
	answer := waysToSplitArray([]int{1, 2, 2, 2, 5, 0})
	println(answer)
}

//specify the problem statement in comments below
// Given an array nums of n integers, find the number of ways to split the array into three non-empty subarrays (left, mid, right) such that:

func waysToSplitArray(nums []int) int {
	count := 0
	totalSum := sumOfArray(nums)
	leftSum := 0

	n := len(nums)
	for i := 0; i < n-1; i++ {
		// if (sumOfArray(nums[0 : i+1]) >= sumOfArray(nums[i+1 : n])) {
		//     count ++
		// }

		leftSum += nums[i]
		rightSum := totalSum - leftSum

		if leftSum >= rightSum {
			count++
		}
	}

	return count
}

func sumOfArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}

	return sum
}