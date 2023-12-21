package pointers

import "fmt"

type Account struct {
    owner string
    balance int
}

func PointersMain() {

    arturosAcc := *NewAccount("Arturo", 200) 
    fmt.Println(arturosAcc.balance)

    addTransaction(arturosAcc, 400)

    // name := "Arturo"
    // // fmt.Println(&name) // reference to memory address using "&"

    // ptr1 := &name // 0x1400008e040 // referencing
    // ptr2 := &name
    // var ptr3 *string = &name

    // *ptr1 = "Santiago" // previously -> "Arturo" // dereferencing
    // *ptr2 = "John"     // previously -> "Santiago"
    // *ptr3 = "Tom"

    // fmt.Println(*ptr3)
    // fmt.Println(name)
}

func NewAccount(name string, balance int) *Account {
    return &Account {
        owner: name,
        balance: balance,
    }
}

func addTransaction(acc Account, amount int) {
    acc.balance += amount
}

