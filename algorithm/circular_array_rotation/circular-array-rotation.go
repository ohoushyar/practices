package main

import "fmt"

func main() {
    args := make([]int, 3)
    for i:=0; i<3; i++ {
        fmt.Scanf("%v", &args[i])
    }

    n, k, q := args[0], args[1], args[2]
    k = k%n

    arr := make([]int, n)
    for i:=0; i<n; i++ {
        fmt.Scanf("%v", &arr[i])
    }

    var idx int
    for i:=0; i<q; i++ {
        fmt.Scanf("%v", &idx)
        idx = ((n-k)+idx)%n
        fmt.Println(arr[idx])
    }
}

