package main

import "fmt"

var cont int

func main() {
    numDiscos := 3
    hanoi(numDiscos, "Torre A", "Torre C", "Torre B")
    fmt.Printf("foi feita em %d rodadas.\n", cont)
}

func hanoi(n int, torreA, torreC, torreB string) {
    if n == 1 {
    cont++
    fmt.Printf("Disco 1 de %s pra %s\n", torreA, torreC)
    return
    }

    hanoi(n-1, torreA, torreB, torreC)
    cont++
    fmt.Printf("Disco %d de %s pra %s\n", n, torreA, torreC)
    hanoi(n-1, torreB, torreC, torreA)
}