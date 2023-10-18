package concurrency

import (
	"fmt"
	"time"
)

func WorkerMain() {
    const numbJobs = 10
    completed := []int{} 

    jobsChan := make(chan int, numbJobs)
    jobsCompleted := make(chan int, numbJobs)

    for w := 1; w <= 3; w++ {
       go worker(w, jobsChan, jobsCompleted) 
    }

    // blocking code
    for j := 1; j <= 10; j++ {
        jobsChan <- j
    }

    close(jobsChan)

    for k := 1; k <= numbJobs; k++ {
        processed := <- jobsCompleted
        completed = append(completed, processed)
    }

    fmt.Println("Completed & processed slice:", completed)
}
                            // read only,               // write only
func worker(id int, jobsChan <- chan int, jobsCompleted chan <- int) {
    for job := range jobsChan {
        time.Sleep(time.Second * 1)
        fmt.Println("worker\t", id, "\t[started]\t job", job, "with", len(jobsChan), "jobs left to process")
        fmt.Println("worker\t", id, "\t[finished]\t job", job)
        jobsCompleted <- job * job 
    }

}
