package reversepolishnotation

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation/
func EvalRPN(tokens []string) int {
	
	// Approach
	// create a stack
    // iterate over the tokens
    // push each integerinto stack
    // if a operand is encountered do pop, pop, push answer
	// result will be the only/last element in the stack

    stack := make([]int, 0)
    for _, v := range tokens {
        if v == "+" || v == "-" || v == "*" || v == "/" {
            // pop
            a := stack[len(stack)-1]
            stack = stack[:len(stack)-1]
            // pop again
            b := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            switch v {
            case "+":
                stack = append(stack, a+b)
            case "-":
                stack = append(stack, b-a)
            case "*":
                stack = append(stack, a*b)
            case "/":
                stack = append(stack, b/a)
            }
        } else {
            num, _ := strconv.Atoi(v)
            stack = append(stack, num)
        }
    }
    return stack[len(stack)-1]
}