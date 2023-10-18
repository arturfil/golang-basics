package main

import "fmt"

func stringFormattingLogic() {
    // name := "Arturo"
    // sport := "Soccer"
    // fmt.Print("Hello I am", name, "and I like to play", sport)
    host := "localhost"
    username := "arturo"
    password := "Password123"
    dbName := "dbNameExample"

    dbConnection := fmt.Sprintf("%v:/%v@%v/%v=exampleDbConnectionString", host, username, password, dbName)
    
    fmt.Printf("-> Connection String: %T ", dbConnection) // & -> address in memory 

    // fmt.Printf("Hello I am %s, and I like to play %v", name, sport)

}
