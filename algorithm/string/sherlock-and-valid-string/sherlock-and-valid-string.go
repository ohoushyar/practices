package main

import "fmt"
import "os"
import "math"

func main() {
    var str string
    fmt.Fscan(os.Stdin, &str)
    debug("-> str: "+str)

    occmap := get_str_occmap(str)

    yes := true
    occmap_len := len(occmap)
    debug(occmap)
    if occmap_len > 2 {
        yes = false
    } else if occmap_len == 2 {
        big, small := 0, 0
        for occ, cnt := range occmap {
            if big == 0 {
                big = occ
            } else if cnt >= occmap[big] {
                small = big
                big = occ
            } else {
                small = occ
            }
        }
        debug(fmt.Sprintf("big: %v small: %v", big, small))
        if occmap[small] != 1 {
            no_switch := occmap[small] * int(math.Abs(float64(big-small)))
            debug(fmt.Sprintf("no switch: %v", no_switch))
            if no_switch > 1 { yes = false }
        }
    }

    if yes {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
}

func get_str_occmap(str string) map[int]int {
    strmap := make(map[byte]int)
    strlen := len(str)
    for i:=0; i<strlen; i++ {
        strmap[str[i]]++
    }
    debug(strmap)
    occmap := make(map[int]int)
    for _, occ := range strmap {
        occmap[occ]++
    }
    return occmap
}

func debug(log interface{}) {
    fmt.Fprintln(os.Stderr, log)
}
