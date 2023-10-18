package concurrency

import (
	"fmt"
	"sync"
)

type transaction struct {
    description string
    amount int
}

var (
    balance int = 0
    wg2 sync.WaitGroup
    mutex sync.Mutex

    deposits = []transaction {
        {"deposit", 500},
        {"deposit", 200},
    }

    widthrawls = []transaction {
        {"widthrawl", 500},
        {"widthrawl", 200},
    }
)

func MutexMain() {

    for _, trx := range deposits {
        for i := 0; i < 50000; i++ {
            go deposit(trx.amount, &wg2)
        }
    }

    for _, trx := range widthrawls {
        for i := 0; i < 50000; i++ {
            widthrawl(trx.amount, &wg2)
        }
    }

    wg2.Wait()

    fmt.Println("Balance: ", balance)
}


func deposit(value int, wg *sync.WaitGroup) {
    wg.Add(1)
    mutex.Lock()
    balance += value
    fmt.Printf("[Deposit]\t | %d\t | to account\t | with balance %d\n", value, balance)
    mutex.Unlock() 
    wg.Done()
}

func widthrawl(value int, wg *sync.WaitGroup) {
    wg.Add(1)
    mutex.Lock()
    balance -= value
    fmt.Printf("[Withdrawing]\t | %d\t | from account\t | with balance %d\n", value, balance)
    mutex.Unlock()
    wg.Done()
}
