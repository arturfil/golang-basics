package concurrency

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func WaitGroupsMain() {

    start := time.Now()
   
    /*
    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(time.Second * 3)
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
    }()
    

    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(time.Second * 3)
        for i := 10; i < 20; i++ {
            fmt.Println(i)
        }
    }()
    

    wg.Add(1)
    go func() {
        defer wg.Done()
        time.Sleep(time.Second * 3)
        for i := 20; i < 30; i++ {
            fmt.Println(i)
        }
    }()
    */

    /*
    wg.Add(3)

    go loopNums(&wg, 0, 10)
    go loopNums(&wg, 10, 20)
    go loopNums(&wg, 20, 30)
    */
    
    // wg.Wait()

    urls := []string {
        "https://google.com",
        "https://twitter.com",
        "https://youtube.com",
        "https://udemy.com",
    }

    getStatus(&wg, urls)

    fmt.Println("Wait group is done")

    end := time.Since(start)
    fmt.Println("It has taken: ", end)
} 

func loopNums(wg *sync.WaitGroup, init int, end int) {
    defer wg.Done()

    time.Sleep(time.Second * 3)
    for i := init; i < end; i++ {
        fmt.Println(i)
    }
}

func getStatus(wg *sync.WaitGroup, urls []string) {
    for _, url := range urls {
        wg.Add(1)

        go func(url string) {
            time.Sleep(time.Second * 2)
            defer wg.Done()
            res, err := http.Get(url) 
                        
            if err != nil {
                fmt.Printf("[Error: %s x -> %s]\n", err, url) 
            } else {
                fmt.Printf("[%d %s]\n", res.StatusCode, url)
            }
        }(url)
    }

    wg.Wait()
}
