package parenthesis

// https://leetcode.com/problems/valid-parentheses/
func IsValid(s string) bool {
    stack := make([]rune, 0)

    m := make(map[rune]rune)
    m['}'] = '{'
    m[')'] = '('
    m[']'] = '['

    for _, v := range s {
        if v == '(' || v == '{' || v == '[' {
            stack = append(stack, v)
        } else {
            if len(stack) == 0 {
                return false
            } else {
                // pop
                top := stack[len(stack)-1]
                if m[v] != top {
                    return false
                }
                stack = stack[:len(stack)-1]
            }
        }
    }
    return len(stack) == 0
}