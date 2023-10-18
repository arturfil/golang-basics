package functions

import "fmt"

func getInfo(first string, last string) (string, string) {
    fmt.Printf("Hello %v %v\n", first, last)
    fullname := fmt.Sprintf("%v %v", first, last)
    fmt.Println("Inside getInfo: ", fullname)
    return first, last 
}

func FunctionsMain() {
    fmt.Println("WOrks")

    var fullname = "John Doe"
    fmt.Println(fullname)

    first, last := getInfo("Arturo", "Filio")
    fmt.Printf("The values are %v, %v, %v\n", first, last, fullname)

    fmt.Println(fullname)
}
