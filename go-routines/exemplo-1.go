package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        fmt.Println(s,i)
    }
}
func main() {
	say("Sem go routine")
    go say("Com go routine")
	fmt.Println("Teste")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Chegou no fim")
}