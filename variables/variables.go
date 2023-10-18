package variables 

import "fmt"

var name string = "Arturo"
var age int8 = 30

const (
	firstName = "Arturo"
	lastName  = "Filio"
)

var (
	currentProjectStack string = "Golang, Postgres, React..."
	currentLaptop       string = "Macbook"
)

func RunVariableLogic() {
	fmt.Println("My first name is", firstName)
	fmt.Println("My current laptop is", currentLaptop)
	
    var myTeam string
	myTeam = "Manchester United"
    
    fmt.Println(myTeam)

	myFavSoccerTeam := "Manchester United"
    myFavSoccerTeam = "Bayern Munich"

    fmt.Println(myFavSoccerTeam)
}
