package main

import "fmt"

func main() {
    var n int
    fmt.Scanf("%v", &n)

    for i:=1; i<=n; i++ {
        for j:=1; j<=n; j++ {
            if j<=n-i {
                fmt.Printf("%s", " ")
            } else {
                fmt.Printf("%s", "#")
            }

            if j==n && i<n {
                fmt.Println()
            }
        }
    }
}
