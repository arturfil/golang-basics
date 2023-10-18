package concurrency

import (
	"fmt"
	"os"
	"sync"
	"time"
)


func ChannelsMain() {
   // unbufferedChannels() 
   // bufferedChannels()
   selecChannels()
}

var wgc sync.WaitGroup

func unbufferedChannels() {
    c := make(chan int) 
    start := time.Now()

    completed := []int{}
    
    for w := 1; w <= 3; w++ {
        wgc.Add(1)
        go func() {
            defer wgc.Done()
            for val := range c {
                time.Sleep(time.Millisecond * 1000)
                fmt.Println(val)
                val *= val
                val += 1
                completed = append(completed, val)
            } 
        }()
    }
    
    for i := 1; i <= 10; i++ {
        c <- i 
    }

    close(c)
    wgc.Wait()

    end := time.Since(start)
    fmt.Println("Completed", completed)
    fmt.Println("Time taken:", end)
}

func bufferedChannels() {
    c := make(chan int, 2)    

    for w := 1; w <= 6; w++ {
        wg.Add(1)
         go func() {
            defer wg.Done()
            for val := range c {
                fmt.Println("Read\t", val, "\tfrom channel")
                time.Sleep(time.Second * 2)
            }
        }()
   }

   for i := 1; i <= 10; i++ {
        c <- i
        fmt.Println("Wrote\t", i, "\tto channel")
    }

    close(c)
    wg.Wait()
}

func selecChannels() {
    c1 := make(chan string)    
    c2 := make(chan string)    
    c3 := make(chan string)    

    go func() {
        for {
            time.Sleep(time.Millisecond * 500)
            c1 <- "Ever 500ms"
        }
    }()
    go func() {
        for {
            time.Sleep(time.Millisecond * 2000)
            c2 <- "Ever 2 seconds"
        }
    }()
    go func() {
        for {
            time.Sleep(time.Second * 8)
            c3 <- "Processes finished"
        }
    }()

    for {
        select {
        case ms1 := <- c1:
            fmt.Println(ms1)
        case ms2 := <- c2:
            fmt.Println(ms2)
        case ms3 := <- c3:
            fmt.Println(ms3)
            os.Exit(0)
        }
    } 
}

