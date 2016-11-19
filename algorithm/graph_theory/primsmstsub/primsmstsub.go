package main

import (
    "container/heap"
    "fmt"
)

type Edge struct {
    Vert Vertex
    Len int
}

type Vertex struct {
    Name int
    Adj map[int]bool
    Dist int
    Seen bool
}

// An IntHeap is a min-heap of ints.
type EdgeHeap []Edge

func (h EdgeHeap) Len() int           { return len(h) }
func (h EdgeHeap) Less(i, j int) bool { return h[i].Len < h[j].Len }
func (h EdgeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *EdgeHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}

func (h *EdgeHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

func main() {
    var n, m int
    fmt.Scanf("%v", &n)
    fmt.Scanf("%v\n", &m)

    graph := make(map[int]Edge)
    for i:=1; i<=n; i++ {
        l := make(map[int]bool)
        graph[i] = Edge{
            Vertex{i, l, -1, false},
            0,
        }
    }


    for i:=0; i<m; i++ {
        var node, adj, len int
        fmt.Scanf("%v", &node)
        fmt.Scanf("%v", &adj)
        fmt.Scanf("%v\n", &len)
        fmt.Printf("n: %v a: %v l: %v\n", node, adj, len)
    }

    var s int
    fmt.Scanf("%v\n", &s)
    fmt.Printf("start: %v\n", s)
}
