package main

import (
    "container/heap"
    "fmt"
)

type Edge struct {
    Src int
    Dest int
    Len int
}

type TAdj map[int]*Edge

type Vertex struct {
    Name int
    Adj TAdj
    Queued bool
    InQ bool
}

type Graph map[int]*Vertex

func main() {
    var n, m int
    fmt.Scanf("%v", &n)
    fmt.Scanf("%v\n", &m)

    graph := make(Graph)
    for i:=1; i<=n; i++ {
        graph[i] = &Vertex{
            Name: i,
            Adj: make(TAdj),
            Queued: false,
            InQ: false,
        }
    }

    for i:=0; i<m; i++ {
        var src, dest, len int
        fmt.Scanf("%v", &src)
        fmt.Scanf("%v", &dest)
        fmt.Scanf("%v\n", &len)
        // fmt.Printf("n: %v a: %v l: %v\n", src, dest, len)

        e := Edge{ src, dest, len }
        add_adj(graph[src], e)
        // fmt.Printf("graph[%v]: %v\n", src, graph[src])

        // bi-directional
        src, dest = dest, src
        e = Edge{ src, dest, len }
        add_adj(graph[src], e)
        // fmt.Printf("graph[%v]: %v\n", src, graph[src])
    }

    // print_g(graph)

    var s int
    fmt.Scanf("%v\n", &s)
    // fmt.Printf("start: %v\n", s)

    fmt.Println(prim_mst(graph, s))
}

func add_adj(vert *Vertex, edge Edge) {
    dest := edge.Dest
    x_edge, ok := vert.Adj[dest]
    if !ok || (ok && edge.Len < x_edge.Len) {
        vert.Adj[dest] = &edge
    }
}

func print_g(g Graph) {
    for k, v:= range g {
        fmt.Println(k)
        for _, e := range(v.Adj) {
            fmt.Printf("---> dest: %v len: %v\n", e.Dest, e.Len)
        }
        fmt.Println()
    }
}

func prim_mst(g Graph, start int) int {
    // fmt.Println("--- mst ---")
    pq := make(PriorityQueue, 0)
    idx := 0
    heap.Init(&pq)

    mst := make(Graph)

    item := &Item{
        value: &Edge{
            Src: -1,
            Dest: start,
            Len: 0,
        },
        index: idx,
    }
    g[start].InQ = true
    g[start].Queued = true

    // fmt.Printf("%v ", start)
    heap.Push(&pq, item)
    idx++
    // fmt.Println("... pushed")


    res := 0
    for pq.Len() > 0 {
        // printq(&pq)
        item := heap.Pop(&pq).(*Item)
        edge := item.value
        g[edge.Dest].InQ = false
        // _, ok := mst[item.value.Dest.Name]
        // if ok {
        //     continue
        // }

        // fmt.Printf("%v -> %v (%v):", edge.Src, edge.Dest, edge.Len)

        res += edge.Len
        // fmt.Printf("%v\n", res)

        if edge.Src != -1 {
            mst[edge.Src].Adj[edge.Dest] = edge
        }
        ix := edge.Dest
        mst[ix] = &Vertex{
            Name: ix,
            Adj: make(TAdj),
            Queued: false,
            InQ: false,
        }
        // fmt.Println(ix)

        adj := g[edge.Dest].Adj
        for _, edge := range(adj) {
            // fmt.Printf("%v => %v (%v) ", edge.Src, edge.Dest, edge.Len)
            item := &Item{
                value: edge,
                index: idx,
            }

            if !g[edge.Dest].Queued {
                heap.Push(&pq, item)
                g[edge.Dest].Queued = true
                g[edge.Dest].InQ = true
                // fmt.Println("... pushed")
                idx++
            } else if g[edge.Dest].InQ {
                update_edge_inq(&pq, edge.Src, edge.Dest, edge.Len)
                // fmt.Println("Queue updated")
            }
        }
        // fmt.Println()
    }

    // print_g(mst)
    return res
}

func update_edge_inq(pq *PriorityQueue, src int, dest int, len int) {
    queue := *pq
    for _, item := range queue {
        if dest == item.value.Dest && len < item.value.Len {
            // fmt.Printf("Found dest: %v\n", dest)
            value := &Edge{
                Src: src,
                Dest: dest,
                Len: len,
            }
            pq.update(item, value)
            break
        }
    }
}

// An Item is something we manage in a priority queue.
type Item struct {
    value *Edge // The value of the item; arbitrary.
    // The index is needed by update and is maintained by the heap.Interface methods.
    index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
    // We want Pop to give us the highest, not lowest, priority so we use greater than here.
    return pq[i].value.Len < pq[j].value.Len
}

func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
    n := len(*pq)
    item := x.(*Item)
    item.index = n
    *pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    item.index = -1 // for safety
    *pq = old[0 : n-1]
    return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value *Edge) {
    item.value = value
    heap.Fix(pq, item.index)
}

func printq(pq *PriorityQueue) {
    queue := *pq
    for i:=0; i<len(*pq); i++ {
        v := queue[i].value
        fmt.Printf("value: src: %v dest: %v len: %v, index: %v\n", v.Src, v.Dest, v.Len, queue[i].index);
    }
}
