
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

    quick(&a, 0, len(a)-1)
}

func quick(a *[]int, lo, hi int) {
    if hi - lo < 1 { return; }
    debug(fmt.Sprintf("... lo: %v hi: %v", lo, hi))

    p := partition(a, lo, hi)
    quick(a, lo, p-1)
    quick(a, p+1, hi)
}

func partition(a *[]int, lo, hi int) int {
    i := lo
    pivot := (*a)[hi]

    for j:=lo; j<hi; j++ {
        debug(fmt.Sprintf("+++ pivot: %v a[%v] %v <=> a[%v] %v", pivot, i, (*a)[i], j, (*a)[j]))
        if (*a)[j] <= pivot {
            (*a)[i], (*a)[j] = (*a)[j], (*a)[i]
            i++
        }
        debug( *a )
    }
    (*a)[i], (*a)[hi] = (*a)[hi], (*a)[i]

    print_arr( *a )
    fmt.Println()

    return i
}

func print_arr(a []int) {
    for i:=0; i<len(a); i++ {
        fmt.Printf("%v ", a[i])
    }
}

func debug(a interface{}) {
    fmt.Fprintln(os.Stderr, a)
}

