package main

import (
    "container/list"
    "fmt"
)

const DEF_EDGE = 6
const INF = -1

type Queue struct {
    IN *list.List
    OUT *list.List
}

func (s *Queue) init() {
    s.IN = list.New()
    s.OUT = list.New()
}

func (s *Queue) enqueue(e int) *list.Element {
    return s.IN.PushFront(e)
}

func (s *Queue) dequeue() *list.Element {
    s._move()
    e := s.OUT.Front()
    s.OUT.Remove(e)
    return e
}

func (s *Queue) is_empty() bool {
    s._move()
    return s.OUT.Len() == 0
}

func (s *Queue) _move() {
    if s.OUT.Len() == 0 {
        for s.IN.Len() > 0 {
            e := s.IN.Front()
            s.OUT.PushFront(s.IN.Remove(e))
        }
    }
}

func (s *Queue) Len() int64 {
    return int64(s.IN.Len() + s.OUT.Len())
}

type Vertex struct {
    name int
    // adj *list.List
    adj map[int]bool
    dist int
    seen bool
}

func main() {
    var q int
    fmt.Scanf("%v\n", &q)
    for i:=0; i<q && i<10; i++ {
        var n, m int
        fmt.Scanf("%v", &n)
        fmt.Scanf("%v\n", &m)

        graph := make(map[int]Vertex)
        for j:=1; j<=n; j++ {
            l := make(map[int]bool)
            graph[j] = Vertex{
                j, l, -1, false,
            }
        }

        for j:=0; j<m; j++ {
            var node, adj int
            fmt.Scanf("%v", &node)
            fmt.Scanf("%v\n", &adj)

            _, ok := graph[node].adj[adj]
            if !ok {
                graph[node].adj[adj] = true
            }

            // it's bi-directional so we need to add both
            node, adj = adj, node
            _, ok = graph[node].adj[adj]
            if !ok {
                graph[node].adj[adj] = true
            }
        }

        var s int
        fmt.Scanf("%v\n", &s)

        graph = bfs(graph, s)
        for j:=1; j<=n; j++ {
            if j == s {
                continue
            }
            fmt.Printf("%v ", graph[j].dist)
        }
        fmt.Println()
    }
}

func bfs(g map[int]Vertex, s int) map[int]Vertex {
    var queue Queue
    queue.init()

    root := g[s]
    root.dist = 0
    root.seen = true
    queue.enqueue(root.name)
    g[s] = root

    for !queue.is_empty() {
        idx := queue.dequeue().Value.(int)
        current := g[idx]
        for k, _ := range current.adj {
            n := g[k]

            curr_dist := current.dist + DEF_EDGE
            if  n.dist == INF || curr_dist < n.dist {
                n.seen = true
                n.dist = curr_dist
                queue.enqueue(n.name)
                g[n.name] = n
            }
        }
    }

    return g
}

