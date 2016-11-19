package main

import "fmt"

func main() {
    m := make(map[int]string)
    m[1] = "one"
    m[2] = "two"
    m[3] = "three"

    for i, v := range(m) {
        fmt.Printf("i: %v v: %v\n", i, v)
    }
}
