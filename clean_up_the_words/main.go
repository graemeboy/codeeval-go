package main

import "fmt"
import "strings"
import "unicode"
import "log"
import "bufio"
import "os"

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }   
    defer file.Close()
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {

        str := scanner.Text()
        output := ""
        addedSpace := true

        for _, r := range str {
            if unicode.IsLetter(r) {
                output += strings.ToLower(string(r))
                addedSpace = false
            } else if (!addedSpace) {
                output += " "
                addedSpace = true
            }
        }

        fmt.Println(output)
    }   
}