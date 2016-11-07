package main

import "fmt"

func main() {
    var l int
    fmt.Scanf("%v\n", &l)

    arr := make( []int, l )
    for i:=0; i<l; i++ {
        fmt.Scanf("%v", &arr[i])
    }

    var sum int
    for _, val := range(arr) {
        sum += val
    }

    fmt.Println(sum);
}
