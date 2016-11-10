package main

import "fmt"

func main() {
    var len int
    fmt.Scanf("%v", &len)
    var val int
    var pos, neg, zer float32
    for i:=0;i<len;i++ {
        fmt.Scanf("%v", &val)
        if val > 0 {
            pos++
        } else if val < 0 {
            neg++
        } else {
            zer++
        }
    }
    fmt.Printf("%.6f\n%.6f\n%.6f", pos/float32(len), neg/float32(len), zer/float32(len))
}
