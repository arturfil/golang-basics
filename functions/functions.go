package functions

import "fmt"

func getInfo() (string, string) {
    var (
        first, last string
    )

    fmt.Println("Enter firstname: ")
    fmt.Scan(&first)
    fmt.Println("Enter lastname:")
    fmt.Scan(&last)

    fmt.Printf("Hello %v %v\n", first, last)
    fullname := fmt.Sprintf("%v %v", first, last)
    fmt.Println("Inside getInfo: ", fullname)
    return first, last 
}

func FunctionsMain() {
    fmt.Println("WOrks")

    var fullname = "John Doe"
    fmt.Println(fullname)

    first, last := getInfo()
    fmt.Printf("The values are %v, %v, %v\n", first, last, fullname)

    fmt.Println(fullname)
}
