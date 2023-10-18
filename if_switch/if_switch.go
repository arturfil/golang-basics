package ifswitch 

import "fmt"

// lowercase var => private
// uppercase var => public

func IfSwitchMain() {

    // arturosName := "Arturo"
    arturosAge := 17 

    if arturosAge < 18  {
        fmt.Println("You cannot drink")

    } else if arturosAge < 21 {
        fmt.Println("You can drink, under parent supervision")

    } else {
        fmt.Println("You can drink")
    }

    arturosProgammingGrade := "A-"

    switch arturosProgammingGrade {

    case "A": {

        fmt.Println("96%")
    }
    case "A-": {
        fmt.Println("You got 91%")
    }
    case "B+":
        fmt.Println("90%")
    case "B":
        fmt.Println("87%")
    case "B-":
        fmt.Println("82%")

    default:
       fmt.Println("Please test a valid class grade") 
         
    }
 
}
