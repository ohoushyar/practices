package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    var str string
    fmt.Fscanln(os.Stdin, &str)

    var wc int
    strlen := len(str)
    if strlen>0 { wc = 1 }
    for i:=0; i<strlen; i++ {
        if string(str[i]) == strings.ToUpper( string(str[i]) ) {
            wc++
        }
    }

    fmt.Println(wc)
}
