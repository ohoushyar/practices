package main
import "fmt"

func main() {
    var n int
    fmt.Scanf("%v", &n)

    arr := make([]int, n)
    for i:=0; i<n; i++ {
        fmt.Scanf("%v", &arr[i])
    }

    finished := false
    tmp := arr[n-1]
    for i:=n-2; i>=0 && !finished; i-- {
        if arr[i]>tmp {
            arr[i+1] = arr[i]
        } else {
            arr[i+1] = tmp
            finished = true
        }
        print_arr(arr)
    }

    if !finished {
        arr[0] = tmp
        print_arr(arr)
    }
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
