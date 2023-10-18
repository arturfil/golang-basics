package main

import "fmt"

/*
   bool

   string

   int int8 int16 int32 int64
   uint uint8 uint16 uint32 uint64

   uintptr // memory address pointer representation

   byte // alias for int32

   rune // alias for int32
       // represents a unicode code point

   float32 float64

   complex64 complex128 // example of -4 and sqrt -> imaginary complex number
*/

func runPrimitivesLogic() {
    // var name string = "Arturo" 
    // var awake bool = true

    var bankBalance int8 = 127 // -> -128 + 128 => 256 
    
    var mem uintptr = 0x1400012c002

    fmt.Println("The memory address is: ", &bankBalance)
    
    fmt.Println("uintptr mem", mem)
    
    var weight float32 = 175.8
    fmt.Println("Currently my weight is: ", weight)
}
