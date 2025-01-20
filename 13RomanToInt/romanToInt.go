package romantoint

func RomanToInt(s string) int {
    m := make(map[byte]int)
    m['I'] = 1
    m['V'] = 5
    m['X'] = 10
    m['L'] = 50
    m['C'] = 100
    m['D'] = 500
    m['M'] = 1000

    result := 0

    i := 0
    for i = 0; i < len(s)-1; i++ {
       if m[s[i]] < m[s[i+1]] {
        result += m[s[i+1]] - m[s[i]]
        i++
       } else {
        result += m[s[i]]
       }
    }
    for i < len(s) {
        result += m[s[i]]
        i++
    }
    return result
 }