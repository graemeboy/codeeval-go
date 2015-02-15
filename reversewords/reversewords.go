package main

import (
    "fmt"
    "log"
    "bufio"
    "os"
    "strings"
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
        words := strings.Split(input, " ")
        for i := range words {
            fmt.Print(words[len(words) - 1 - i] + " ")
        }
        fmt.Println()
    }
}