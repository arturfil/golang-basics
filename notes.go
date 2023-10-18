package pointers

import "fmt"

// Notes

/*

    how to present:
    - how to see a memory address,
    - how to create a pointer
    - show Pointers Main first
    - show javascript example
    - show chaning variable from pointers

    - why use pointers
        > stack vs heap
            > heap is cheaper in memory
        > easy to change variables
        > not sure about the data ? // slices vs arrays
        > pointers could be more expensive if used too much.
        > https://go.dev/doc/faq#stack_or_heap

    memory address

    var name string = "Arturo"

    a pointer is a variable that wholes an address value

    var *ptr_name string = &name // using the "*" is dereferencing

    // why use pointers?
    - We can change a variable from various different pointers without having to 
    duplicate data, it's just faster and "cheaper for memory"

    // stack vs heap -> not the heap data structure

    // heap can be more memory effective but could be more expensive for the garbage collector
    // - (getting rid of memory)
   
*/

type soccerPlayer struct {
    name string
    team string
}

// returninig value vs returning pointer
/*func initSoccerPlayer(name string, team string) soccerPlayer { // this return the copy of the value
    player := soccerPlayer{name: name, team: team}
    return player
}*/

func initSoccerPlayer(name string, team string) *soccerPlayer { // this return the copy of the value
    player := soccerPlayer{name: name, team: team}
    fmt.Printf("\ninitSoccerPlayer --> %p\n", &player)
    return &player
}

func PointersMain() {
    fmt.Println("Assigning variables")
    fmt.Println("-------------------------------------------------")
    name := "Arturo"
    fmt.Printf("'name' var:%-15s\n", name)

    // these three are the same
    name_add := &name // 0x14000254210
    var name_ptr *string = &name
    var name_ptr2 *string = &name

    // printing the values
    fmt.Printf("name_ptr\t%-15p \tvar: name\n", name_ptr) 
    fmt.Printf("addr ptr\t%-15s \tvar: name\n", *name_add) // dereferencing
    fmt.Printf("name_ptr\t%-15s \tvar: name\n", *name_ptr) // dereferencing too of value atm "Arturo"

    fmt.Println("\nchaning variables")
    fmt.Println("-------------------------------------------------")

    *name_add = "Santiago"
    fmt.Printf("My name is\t%-15s \tvar: name\n", name)

    *name_ptr = "Jorge"
    fmt.Printf("My name is\t%-15s \tvar: name\n", name)

    *name_ptr2 = "Armin Van Buuren"
    fmt.Printf("My name is\t%-15s \tvar: name\n", name)

    // fmt.Println("\nsoccer player", initSoccerPlayer("Frenkie De Jong", "Barcelona"))
    fmt.Printf("\nmain func --> player %p", initSoccerPlayer("Frenkie De Jong", "Barcelona"))
}

