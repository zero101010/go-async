package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        fmt.Println(s,i)
		time.Sleep(time.Second)
    }
}
func main() {
	go say("a")
	go say("b")
	time.Sleep(time.Second *10)
}