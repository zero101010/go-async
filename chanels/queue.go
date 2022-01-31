package main
import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan int)

	go func() {
		i:=0
		for {
			time.Sleep(time.Second)
			queue <- i
			fmt.Println("Setando valor para a fila",i)
			i++
		}
	}()

	for x := range queue {
		fmt.Println("Lendo buffer da fila",x)
	}
}