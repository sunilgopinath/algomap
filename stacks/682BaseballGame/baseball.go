package baseballgame

import "strconv"

// https://leetcode.com/problems/baseball-game/
func CalPoints(operations []string) int {
    stack := make([]int, 0) // Initialize stack

    for _, v := range operations {
        switch v {
        case "+":
            if len(stack) >= 2 {
                stack = append(stack, stack[len(stack)-1]+stack[len(stack)-2])
            }
        case "D":
            if len(stack) > 0 {
                stack = append(stack, 2*stack[len(stack)-1])
            }
        case "C":
            if len(stack) > 0 {
                stack = stack[:len(stack)-1] // Remove the last score
            }
        default:
            d, _ := strconv.Atoi(v) // Convert string to integer
            stack = append(stack, d)
        }
    }

    sum := 0
    for _, v := range stack {
        sum += v
    }
    return sum
}