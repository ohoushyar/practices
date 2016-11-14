package main

import (
    // "bufio"
    "container/list"
    "fmt"
    // "os"
)

type Shop struct {
    name string
    fishes []int
}

type Edge struct {
    dest Vertex
    cost int
}

type Vertex struct {
    name string
    adj *list.List
    shop Shop
}

func main() {
    var n, m, k int
    fmt.Scanf("%v", &n)
    fmt.Scanf("%v", &m)
    fmt.Scanf("%v", &k)
    fmt.Printf("n: %v m: %v k: %v\n", n, m, k)

    graph := make(map[string]Vertex)
    for i:=1; i<=n; i++ {

        var f int
        fmt.Scanf("%v", &f);
        fishes := make([]int, f)
        for j:=0; j<f; j++ {
            fmt.Scanf("%v", &fishes[j])
        }
        // fmt.Printf("fishes: %v\n", fishes)

        istr := fmt.Sprintf("%v", i)
        l := list.New()
        shop := Shop{ istr, fishes }
        graph[istr] = Vertex{
            istr, l, shop,
        }

        fmt.Printf("graph[%v]: %v\n", istr, graph[istr])
    }

    // fmt.Printf("graph: %v\n", graph)
    for i:=1; i<=m; i++ {

        var src, dest, cost int
        fmt.Scanf("%v", &src )
        fmt.Scanf("%v", &dest )
        fmt.Scanf("%v", &cost )

        srcstr := fmt.Sprintf("%v", src)
        deststr := fmt.Sprintf("%v", dest)
        // fmt.Printf("Adding src[%v] to dest[%v] with cost[%v]\n", srcstr, deststr, cost)

        graph[srcstr].adj.PushBack(Edge{ graph[deststr], cost })
        // fmt.Printf("graph[%v]: adjlen: %v\n", istr, graph[istr].adj.Len())
    }

    // for i:=1; i<=n; i++ {
    //     istr:= fmt.Sprintf("%v", i)
    //     fmt.Printf("graph[%v]:\n", istr)
    //     l := graph[istr].adj
    //     // Iterate through list and print its contents.
    //     for e := l.Front(); e != nil; e = e.Next() {
    //         fmt.Println(e.Value)
    //     }
    // }
}


