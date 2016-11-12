package main

import (
    "bufio"
    "fmt"
    "math/big"
    "os"
)

func main() {
    var t1, t2, n int
    in := bufio.NewReader( os.Stdin )
    fmt.Fscan( in, &t1, &t2, &n )
    fibonacci := fib(t1, t2)

    res := new(big.Int)
    for i:=2; i<n; i++ {
        res = fibonacci()
    }

    fmt.Println(res)
}

func fib(t1, t2 int) func() *big.Int {
    ti := big.NewInt( int64(t1) )
    ti1 := big.NewInt( int64(t2) )

    return func() *big.Int {
        ti2 := new(big.Int)
        ti1sqr := new(big.Int)
        ti2.Add(ti, ti1sqr.Mul(ti1, ti1))
        ti, ti1 = ti1, ti2

        return ti2
    }
}
