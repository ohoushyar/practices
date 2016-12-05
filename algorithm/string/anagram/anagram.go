package main

import "fmt"
import "os"

func main() {
    var t int
    fmt.Fscan(os.Stdin, &t)

    for i:=0; i<t; i++ {
        var test string
        fmt.Fscan(os.Stdin, &test)
        debug( fmt.Sprintf("-> test: %s", test) )

        tlen := len(test)
        if tlen % 2 != 0 {
            fmt.Println(-1)
            continue
        }

        str1 := make(map[byte]int)
        str2 := make(map[byte]int)

        maxj := tlen/2
        for j:=0; j<maxj; j++ {
            str1[test[j]]++
            str2[test[j+maxj]]++
        }
        debug(fmt.Sprintf("str1: %v", str1))
        debug(fmt.Sprintf("str2: %v", str2))

        var repcnt int
        for k, xcnt := range str1 {
            ycnt, ok := str2[k]
            if ok && xcnt > ycnt {
                repcnt += xcnt - ycnt
            } else if !ok {
                repcnt += xcnt
            }
        }
        fmt.Println(repcnt)
    }

}

func debug(log interface{}) {
    return
    fmt.Fprintln(os.Stderr, log)
}
