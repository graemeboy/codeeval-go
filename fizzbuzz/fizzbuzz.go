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
        output := ""
        inArr := strings.Split(input, " ")
        n, err := strconv.ParseInt(inArr[2], 10, 64)
        y, err := strconv.ParseInt(inArr[1], 10, 64)
        x, err := strconv.ParseInt(inArr[0], 10, 64)
        if (err != nil) {
            fmt.Printf("Found error.")
        } else {
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
                fmt.Printf(output + " ")
            }
            fmt.Printf("\n")
        }
    }   
}