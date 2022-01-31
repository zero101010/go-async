package main
// Esse exercício tem como objetivo mostrar como funciona uma thread prenptiva, ou seja enquanto um worker espera o tempo de 1 segundo do time.Sleep(time.Second) ele muda de worker e pega a msg

import (
	"fmt"
	"time"
)
// É uma espécie de consumer
func worker(workerId int,  msg chan int){
	for res := range msg {
		fmt.Println("WorkerId:",workerId," Msg:",res)
		time.Sleep(time.Second)
	}
}

func main(){
	msg := make(chan int)
	// Como o valor de msg não foi setado ele vai ter que esperar para seguir para a próxima rotina
	for i:=0;i<2;i++ {
		go worker(i, msg)
	}
	// Age como uma especie de producer que manda o dado para o canal, onde será consumido pelo worker
	for x:=0;x<10;x++ {
		msg <- x
	}

}