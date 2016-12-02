package main

import (
    "fmt"
    "os"
)

var cnt int
func main() {
    var n int
    fmt.Scan(&n)
    debug(fmt.Sprintf("-> %v", n))

    iarr := make([]int, n)
    qarr := make([]int, n)
    var num int;
    for i:=0; i<n; i++ {
        fmt.Scan(&num)
        iarr[i] = num
        qarr[i] = num
    }
    debug(fmt.Sprintf("-> %v", iarr))

    debug("--- insertion sort ---")
    debug(iarr)
    insert(&iarr, 1)
    debug(fmt.Sprintf("cnt: %v res: %v", cnt, iarr))

    debug("--- quick sort ---")
    debug(qarr)
    quicksort(&qarr, 0, len(qarr)-1)
    debug(fmt.Sprintf("cnt: %v res: %v", cnt, qarr))

    fmt.Println(cnt)
}

func quicksort(a *[]int, lo, hi int) {
    debug(fmt.Sprintf("lo: %v hi: %v", lo, hi))
    if hi-lo <= 0 { return }

    p := partition( a, lo, hi )
    quicksort(a, lo, p-1)
    quicksort(a, p+1, hi)
}

func partition(a *[]int, lo, hi int) int {
    pivot := (*a)[hi]
    i := lo
    debug(fmt.Sprintf("pivot: %v", pivot))

    for j:=lo; j<=hi-1; j++ {
        if (*a)[j] <= pivot {
            (*a)[i], (*a)[j] = (*a)[j], (*a)[i]
            i++
            debug(*a)
            cnt--
        }
    }
    (*a)[i], (*a)[hi] = (*a)[hi], (*a)[i]
    cnt--
    debug(*a)

    return i
}

func insert(a *[]int, idx int) {
    if idx > len(*a) { return }
    tmp := (*a)[idx-1]
    debug(fmt.Sprintf("tmp: %v idx: %v", tmp, idx))

    var i int
    for i:=idx-2; i>=0; i-- {
        if (*a)[i] >= tmp {
            (*a)[i+1] = (*a)[i]
            cnt++
        } else {
            (*a)[i+1] = tmp
            debug(*a)
            break
        }
        debug(*a)
    }

    if i==0 && (*a)[i] > tmp {
        (*a)[i] = tmp
        debug(*a)
    }

    insert(a, idx+1)
}

func debug(val interface{}) {
    return
    fmt.Fprintf(os.Stderr, "%v\n", val)
}
