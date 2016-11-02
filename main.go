package main

import (
    "fmt"
)

func main() {
    fmt.Println("Starting Set 1!")
    // Set 1 Challenge 1
    fmt.Println("-- Challenge 1 --")
    basics_str := Basics_Chall1()
    fmt.Println(basics_str + "\n")

    // Set 1 Challenge 2
    fmt.Println("-- Challenge 2 --")
    basics_str = Basics_Chall2()
    fmt.Println(basics_str + "\n")

    // Set 1 Challenge 3
    fmt.Println("-- Challenge 3 --")
    plaintext, plain_val := Basics_Chall3()
    fmt.Printf("Plaintext: %v\nValue: %v\n", plaintext, plain_val)

    // Set 1 Challenge 4
    fmt.Println("-- Challenge 4 --")
    plaintext, plain_val = Basics_Chall4()
    fmt.Printf("Plaintext: %v\nValue: %v\n", plaintext, plain_val)

    // Set 1 Challenge 5
    fmt.Println("-- Challenge 5 --")
    ciphertext := Basics_Chall5()
    fmt.Printf("Ciphertext: %v\n", ciphertext)
}
