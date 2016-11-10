package main

import "fmt"

func main() {
    s, t := read_input()
    a, b := read_input()
    m, n := read_input()

    apple := in_house(s, t, a, m)
    fmt.Println(apple)

    orange := in_house(s, t, b, n)
    fmt.Println(orange)
}

func in_house(s, t, tree, len int) int {
    in := 0
    for i:=0; i<len; i++ {
        var p int
        fmt.Scanf("%v", &p)
        if tree+p >=s && tree+p <= t {
            in++
        }
    }
    return in
}

func read_input() (int, int) {
    arr := make([]int, 2)
    for i:=0; i<2; i++ {
        fmt.Scanf("%v", &arr[i])
    }
    return arr[0], arr[1]
}
