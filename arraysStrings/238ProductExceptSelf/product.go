package productexceptself

// https://leetcode.com/problems/product-of-array-except-self/
func ProductExceptself(nums []int) []int {

	// Realize the product of all elements except self
	// is the product of all elements to the left of self
	// times the product of all elements to the right of self

	// We can calculate the product of all elements to the left of self
	// by iterating through the array and multiplying the previous element
	// with the current element
	// We can calculate the product of all elements to the right of self
	// by iterating through the array in reverse and multiplying the next element
	// with the current element

	// We can then multiply the two arrays to get the result
	// We can optimize the space complexity by using the result array
	// to store the product of all elements to the left of self
	// and then multiply the product of all elements to the right of self
	// with the result array
	// This way we can avoid using two extra arrays
    pre := make([]int, len(nums))
    post := make([]int, len(nums))
    res := make([]int, len(nums))

    pre[0] = 1
    post[len(post)-1] = 1
    for i := 0; i < len(nums)-1;i++ {
        pre[i+1] = pre[i]*nums[i]
    }
    for i := len(nums)-1; i >= 1; i-- {
        post[i-1] = post[i] * nums[i]
    }
    for i := 0; i < len(nums); i++ {
        res[i] = pre[i]*post[i]
    }

    return res
}