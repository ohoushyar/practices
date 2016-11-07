package main

import "fmt"

func main() {
    a := read_trip()
    b := read_trip()

    var alice, bob int
    for i, _ := range(a) {
        if a[i] > b[i] {
            alice += 1
        } else if a[i] < b[i] {
            bob += 1
        }
    }

    fmt.Println(alice, bob)
}

func read_trip() [3]int {
    var a [3]int
    for i:=0; i<3; i++ {
        fmt.Scanf("%v", &a[i])
    }
    return a
}

