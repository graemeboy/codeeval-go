package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "strconv"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }   
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        input := scanner.Text()
        sum := 0
        digit := 0
        for i := range input {
            digit,err = strconv.Atoi(string(input[i]))
            if (err == nil) {
                sum += digit
            }
        }
        fmt.Println(sum)
    }   
}