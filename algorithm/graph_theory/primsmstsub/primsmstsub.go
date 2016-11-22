package main

import (
    "container/heap"
    "fmt"
)

type Edge struct {
    Src *Vertex
    Dest *Vertex
    Len int
}

type Vertex struct {
    Name int
    Adj []Edge
    Dist int
    Seen bool
}

func main() {
    var n, m int
    fmt.Scanf("%v", &n)
    fmt.Scanf("%v\n", &m)

    graph := make(map[int]*Vertex)
    for i:=1; i<=n; i++ {
        graph[i] = &Vertex{
            Name: i,
            Adj: make([]Edge, 0),
            Dist: -1,
            Seen: false,
        }
    }

    for i:=0; i<m; i++ {
        var src, dest, len int
        fmt.Scanf("%v", &src)
        fmt.Scanf("%v", &dest)
        fmt.Scanf("%v\n", &len)
        fmt.Printf("n: %v a: %v l: %v\n", src, dest, len)

        e := Edge{ graph[src], graph[dest], len }
        graph[src].Adj = append(graph[src].Adj, e)
        // fmt.Printf("graph[%v]: %v\n", src, graph[src])

        // bi-directional
        src, dest = dest, src
        e = Edge{ graph[src], graph[dest], len }
        graph[src].Adj = append(graph[src].Adj, e)
        // fmt.Printf("graph[%v]: %v\n", src, graph[src])
    }

    print_g(graph)

    var s int
    fmt.Scanf("%v\n", &s)
    fmt.Printf("start: %v\n", s)

    fmt.Println(prim_mst(graph, s))
}

func print_g(g map[int]*Vertex) {
    for k, v:= range g {
        fmt.Println(k)
        for i := range(v.Adj) {
            e := v.Adj[i]
            fmt.Printf("dest: %v len: %v - ", e.Dest.Name, e.Len)
        }
        fmt.Println()
    }
}

func prim_mst(g map[int]*Vertex, start int) int {
    fmt.Println("--- mst ---")
    pq := make(PriorityQueue, 0)
    idx := 0
    heap.Init(&pq)

    mst := make(map[int]*Vertex)
    mst[start] = g[start]
    mst[start].Seen = true
    fmt.Println(start)
    for i:=0; i<len(mst[start].Adj); i++ {
        fmt.Printf("%v ", mst[start].Adj[i].Dest.Name)
        if !mst[start].Adj[i].Dest.Seen {
            mst[start].Adj[i].Dest.Seen = true
            item := &Item{
                value: &(mst[start].Adj[i]),
                index: idx,
            }
            heap.Push(&pq, item)
            idx++
        }
    }
    fmt.Println()

    res := 0
    for pq.Len() > 0 {
        item := heap.Pop(&pq).(*Item)

        edge := item.value
        fmt.Printf("%v -> %v (%v):", edge.Src.Name, edge.Dest.Name, edge.Len)
        res += edge.Len
        fmt.Printf("%v\n", res)


        ix := int(edge.Dest.Name)
        fmt.Println(ix)
        mst[ix] = edge.Dest
        for i:=0; i<len(mst[ix].Adj); i++ {
            fmt.Printf("%v ", mst[ix].Adj[i].Dest.Name)
            if !mst[ix].Adj[i].Dest.Seen {
                mst[ix].Adj[i].Dest.Seen = true
                item := &Item{
                    value: &(mst[ix].Adj[i]),
                    index: idx,
                }
                heap.Push(&pq, item)
                idx++
            }
        }
        fmt.Println()
    }

    fmt.Printf("res: %v\n", res)
    fmt.Println()
    return res
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    *Edge // The value of the item; arbitrary.
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
