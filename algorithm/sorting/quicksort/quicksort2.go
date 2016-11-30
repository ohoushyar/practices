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

    quick(a)
}

func quick(a []int) []int {
    if len(a) == 0 {
        return make([]int, 0)
    }
    debug(fmt.Sprintf("... %v", a))

    pivot := a[0]
    left := make([]int, 0)
    right := make([]int,0)
    equal := make([]int,1)
    equal[0] = pivot

    for i:=1; i<len(a); i++ {
        if a[i] < pivot {
            left = append(left, a[i])
        } else if a[i] > pivot {
            right = append(right, a[i])
        } else {
            equal = append(equal, a[i])
        }
    }
    left = quick(left)
    right = quick(right)

    sorted := merge( left, equal, right )
    if len(sorted) > 1 {
        print_arr( sorted )
        fmt.Println()
    }

    return sorted
}

func merge(l, e, r []int) []int {
    merged := make([]int, 0)

    for i:=0; i<len(l); i++ {
        merged = append(merged, l[i])
    }
    for i:=0; i<len(e); i++ {
        merged = append(merged, e[i])
    }
    for i:=0; i<len(r); i++ {
        merged = append(merged, r[i])
    }

    return merged
}

func print_arr(a []int) {
    for i:=0; i<len(a); i++ {
        fmt.Printf("%v ", a[i])
    }
}

func debug(a interface{}) {
    fmt.Fprintln(os.Stderr, a)
}

