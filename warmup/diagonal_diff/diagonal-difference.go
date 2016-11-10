package main

import (
    "fmt"
    "math"
)

func main() {
    var n int
    fmt.Scanf("%v\n", &n)

    // read the matrix
    var val, diag1, diag2 int
    for i:=0; i<n; i++ {
        for j:=0; j<n; j++ {
            fmt.Scanf("%v", &val)

            if i == j {
                diag1 += val
            }
            if i+j == n-1 {
                diag2 += val
            }
        }
    }


    // fmt.Printf("diag1: %v diag2: %v\n", diag1, diag2)

    fmt.Println( math.Abs(float64(diag1-diag2)) )
}
