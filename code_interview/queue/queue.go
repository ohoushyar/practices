package main

import (
    "fmt"
    "container/list"
)

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
    // fmt.Printf("Dequeued %v\n", e.Value)
    return e
}

func (s *Queue) _move() {
    if s.OUT.Len() == 0 {
        for s.IN.Len() > 0 {
            e := s.IN.Front()
            s.OUT.PushFront(s.IN.Remove(e))
        }
    }
}

func (s *Queue) head() int {
    s._move()
    return s.OUT.Front().Value.(int)
}

func (s *Queue) printq() {
    for e := s.OUT.Front(); e != nil; e = e.Next() {
        fmt.Printf("%v <- ", e.Value)
    }
    for e := s.IN.Back(); e != nil; e = e.Prev() {
        fmt.Printf("%v <- ", e.Value)
    }
    fmt.Println()
}

func main() {
    var q Queue
    q.init()

    var query int
    fmt.Scanf("%v", &query)
    for i:=0; i<query; i++ {
        var cmd int
        var x int
        fmt.Scanf("%v", &cmd)

        // fmt.Println("-----")
        switch cmd {
        case 1:
            fmt.Scanf("%v", &x)
            // fmt.Printf("Enq %v\n", x)
            q.enqueue(x)
        case 2:
            // fmt.Println("Deq ...")
            q.dequeue()
        case 3:
            fmt.Println(q.head())
        }
        // q.printq()
    }

}
