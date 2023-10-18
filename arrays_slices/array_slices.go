package main

import "fmt"


func arraySlicesLogic() {
    
    // var videoGames  = [7]string {"Zelda", "TLOU", "", "", "", "", ""} // <- an array
    // videoGames[4] = "Batman Arkham City"
    
    videoGames  := []string {"Zelda", "TLOU"} // <- a slice
    videoGames = append(videoGames, "Batman Arkham City") 
    
    movies := []string {"Start Wars", "Avatar"}
    

    fmt.Println(videoGames) 
  
    // deleting The Last of Us
    videoGames = append(videoGames[0:1],videoGames[2:]...)
    gamesAndMovies := []string{} 
    gamesAndMovies = append(videoGames, movies...)

    fmt.Println(gamesAndMovies)
}
