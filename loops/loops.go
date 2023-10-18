package loops

import "fmt"


func LoopsMain() {
    // videogames := []string {"Zelda BOTW", "Zelda TOTK", "Witcher 3", "Metal Gear Solid 3",}
   
    // normal forloop 
    /*for i := 0; i < 100; i++ {
        fmt.Println("We are in round:", i + 1)
    }*/
   
    // for each
    /*for _, game := range videogames {
        fmt.Println(game)
    }*/

    // while loops
    rounds := 0 
    for rounds < 8 {
        fmt.Println(rounds)
        rounds++
    }
    
}
