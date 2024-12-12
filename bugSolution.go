```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 5) // Buffered channel to prevent blocking

        for i := 0; i < 5; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        fmt.Printf("Goroutine %d received: %d\n", i, <-ch)
                }(i)
        }

        for i := 0; i < 5; i++ {
                ch <- i
        }

        wg.Wait()
}
```