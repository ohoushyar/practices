package main

import "fmt"
import "os"


func main() {
    var n int
    fmt.Fscan(os.Stdin, &n)

    str_map_func := get_str_map_func()
    for z:=0; z<n; z++ {
        var str string
        fmt.Fscan(os.Stdin, &str)
        debug(fmt.Sprintf("-> str: %s", str))

        var unanag_cnt int
        perms := get_perms2(str)
        debug(perms)
        for _, strs := range perms {
            strslen := len(strs)
            for i, str1 := range strs {
                debug(str1)
                for j:=i+1; j<strslen; j++ {
                    str2 := strs[j]
                    if is_unanag(str_map_func, str1, str2) { unanag_cnt++; debug(fmt.Sprintf(" ... str1: %s str2: %s",str1, str2)); }
                }
            }
        }
        fmt.Println(unanag_cnt)
    }
}

func is_unanag(str_map func(string) map[byte]int, str1, str2 string) bool {
    str1len := len(str1)
    str2len := len(str2)
    if str1len != str2len { return false }
    s1 := str_map(str1)
    s2 := str_map(str2)

    for chr, s1cnt := range s1 {
        s2cnt, ok := s2[chr]
        if !ok || s2cnt != s1cnt {
            return false
        }
    }

    return true
}

func get_str_map_func() func(str string) map[byte]int {
    strmap := make(map[string]map[byte]int)

    return func(str string) map[byte]int {
        smap, ok := strmap[str]
        if ok { debug("hit: "+str); return smap }
        debug("missed: "+str)

        strlen := len(str)
        smap = make(map[byte]int)
        for i:=0; i<strlen; i++ {
            smap[str[i]]++
        }

        strmap[str] = smap
        return smap
    }
}

type PERMS map[int][]string

func get_perms2(str string) PERMS {
    perms := make(PERMS)
    strlen := len(str)
    for i:=0; i<strlen; i++ {
        for j:=i; j<strlen; j++ {
            var perm string
            for z:=i; z<=j; z++ {
                perm += string(str[z])
            }
            perms[len(perm)] = append(perms[len(perm)], perm)
        }
    }
    return perms
}

func debug(log interface{}) {
    return
    fmt.Fprintln(os.Stdout, log)
}
