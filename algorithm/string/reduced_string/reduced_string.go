package main

import (
    "fmt"
    "os"
)

func main() {
    var res, str string
    fmt.Fscanln(os.Stdin, &str)

    for {
        res = reduce_str(str)
        if res == str { break }
        str = res
    }

    if len(res) == 0 {
        res = "Empty String"
    }

    fmt.Println(res)
}

func reduce_str(str string) string {
    var res string
    debug(fmt.Sprintf("-> str: %s\n", str))

    strlen := len(str)
    for i:=0; i<strlen; i++ {
        if i < strlen-1 {
            if str[i] == str[i+1] {
                i++
            } else {
                res += string(str[i])
            }
        } else {
            res += string(str[i])
        }
    }

    debug(fmt.Sprintf("<- %s\n", res))
    return res
}

func debug(log interface{}) {
    return
    fmt.Fprintf(os.Stderr, "%s", log)
}
