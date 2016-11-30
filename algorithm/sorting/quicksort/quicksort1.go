package main

import (
    "fmt"
    "os"
)

func main() {
    var n int
    fmt.Scan(&n)
    debug(fmt.Sprintf("-> %v", n))

    a := make([]int, n)
    var num int;
    for i:=0; i<n; i++ {
        fmt.Scan(&num)
        a[i] = num
    }
    debug(fmt.Sprintf("-> %v", a))

    pivot := a[0]
    right := make([]int,0)
    equal := make([]int,1)
    equal[0] = pivot
    for i:=1; i<n; i++ {
        if a[i] < pivot {
            fmt.Printf("%v ", a[i])
        } else if a[i] > pivot {
            right = append(right, a[i])
        } else {
            equal = append(equal, a[i])
        }
    }

    print_arr(equal)
    print_arr(right)
}

func print_arr(a []int) {
    for i:=0; i<len(a); i++ {
        fmt.Printf("%v ", a[i])
    }
}

func debug(a interface{}) {
    fmt.Fprintln(os.Stderr, a)
}

