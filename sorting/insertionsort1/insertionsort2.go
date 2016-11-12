package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scanf("%v", &n)

    arr := make([]int, n)
    for i:=0; i<n; i++ {
        fmt.Scanf("%v", &arr[i])
    }

    insert_sort(arr, 1)
}

func insert_sort(arr []int, idx int) []int {

    if idx == len(arr) {
        return arr
    }

    finished := false
    tmp := arr[idx]
    for i:=idx-1; i>=0 && !finished; i-- {
        if arr[i]>tmp {
            arr[i+1] = arr[i]
        } else {
            arr[i+1] = tmp
            finished = true
        }
    }

    if !finished {
        arr[0] = tmp
    }

    print_arr(arr)
    return insert_sort(arr, idx+1)
}

func print_arr(arr []int) {
    len := len(arr)
    for i, v := range(arr) {
        fmt.Printf("%v", v)
        if i+1 != len {
            fmt.Printf(" ")
        } else {
            fmt.Printf("\n")
        }
    }
}
