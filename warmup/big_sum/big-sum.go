package main

import "fmt"

func main() {
    var l int
    fmt.Scanf("%v\n", &l)

    arr := make( []int32, l )
    for i:=0; i<l; i++ {
        fmt.Scanf("%v", &arr[i])
    }

    var sum int64
    for _, val := range(arr) {
        sum += int64(val)
    }

    fmt.Println(sum);
}
