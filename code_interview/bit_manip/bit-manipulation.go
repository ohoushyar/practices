package main

import "fmt"

func main() {
    var n int
    fmt.Scanf("%v", &n)

    var res, tmp int
    for i:=0; i<n; i++ {
        fmt.Scanf("%v", &tmp)
        fmt.Printf("tmp: %b\n", tmp)
        res ^= tmp
        fmt.Printf("res: %b\n", res)
    }

    fmt.Println(res)
}
