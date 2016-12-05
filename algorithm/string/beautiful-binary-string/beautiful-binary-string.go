package main

import (
    "fmt"
    "os"
)

func main() {
    var st, n int
    var bin string
    fmt.Fscanln(os.Stdin, &n)
    fmt.Fscanln(os.Stdin, &bin)

    for i:=0; i<n; i++ {
        if i<n-2 {
            if bin[i] == '0' && bin[i+1] == '1' && bin[i+2] == '0' {
                i += 2
                st++
            }
        }
    }
    fmt.Println(st)
}
