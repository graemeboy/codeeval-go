//Sample code to read in test cases:
package main

import (
    "fmt"
    "log"
    "strings"
    "strconv"
    "bufio"
    "os"
)

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }   
    defer file.Close()
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        //'scanner.Text()' represents the test case, do something with it
        input := scanner.Text()
        inArr := strings.Split(input, " ")
        n, err := strconv.Atoi(inArr[2])
        y, err1 := strconv.Atoi(inArr[1])
        x, err2 := strconv.Atoi(inArr[0])
        if (err != nil) {
            fmt.Printf("Error parsing n")
        } else if (err1 != nil) {
            fmt.Printf("Error parsing y")
        } else if (err2 != nil) {
            fmt.Printf("Error parsing x")
        } else {
            do_fb(n, x, y)
        }
    }
}

func do_fb(n int, x int, y int) {
    output := ""
    for i := 1; i <= int(n); i++ {
        output = ""
        if ( i % int(x) == 0) {
            output += "F"
        }
        if (i % int(y) == 0) {
            output += "B"
        } 

        if (output == "") {
            output += strconv.Itoa(i)
        }
        fmt.Print(output + " ")
    }
    fmt.Println()
}