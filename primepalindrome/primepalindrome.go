package main

import "fmt"
import "strconv"

func main() {
    lastPrimePalindrome := 0
    for i := 0; i < 1000; i++ {
        if isPalindrome(i) && isPrime(i) {
            lastPrimePalindrome = i
        }
    }
    fmt.Println(lastPrimePalindrome)
}

func isPrime(n int) bool {
    for i := 2; i < n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}

func isPalindrome(n int) bool {
    stringN := strconv.Itoa(n)
    return (stringN == reverseString(stringN))
}

func reverseString(s string) string {
    output := ""
    for i := range s {
        output += string(s[len(s) - 1 - i])
    }
    return output
}