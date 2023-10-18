package mapsstructs

import "fmt"

type Product struct {
    id string
    name string 
    price float64
    category string
}

func (p *Product) returnInfo() string {
   info := fmt.Sprintf("\nid:\t %v\nname:\t %v\nprice:\t %v\ncategory:\t %v", p.id, p.name, p.price, p.category) 
   fmt.Println(info)
   return info
}

func createMap(products []Product) map[string]int {
    invMap := make(map[string]int)

    for _, p := range products {
        p.returnInfo()

        if _, key_exists := invMap[p.category]; key_exists {
            invMap[p.category] += 1
        } else {
            invMap[p.category] = 1
        }

        fmt.Println("*** product processed ***")
    }

    return invMap
}

func MapStructs() {

    products := []Product {
        {"afsd-809234", "Nintendo Switch", 199.55, "Videogames"},
        {"afsd-809234", "PS5", 450.20, "Videogames"},
        {"af12341234", "Macbook Pro", 1899.99, "Computers"},
        {"3jjlkj12j4", "Mech Keyboard", 234.33, "Accesories"},
        {"3jjlkj12j4", "Raspberry Pi", 40, "Computers"},
        {"1234-89jasdf", "Fender Guitar", 740, "Music Instrument"},
    }

    inventory := createMap(products)
   
    fmt.Println("\n---------- inventory ----------")
    fmt.Println("\nIn Inventory: ", inventory)
}
