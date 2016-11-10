package main

import "fmt"

func main() {
    var h, m, s int
    var mid string
    fmt.Scanf("%2d:%2d:%2d%s", &h, &m, &s, &mid)

    if mid == "PM" {
        if h < 12 {
            h += 12
        }
    } else {
        if h == 12 {
            h = 0
        }
    }

    fmt.Printf("%02d:%02d:%02d", h, m, s)
}
