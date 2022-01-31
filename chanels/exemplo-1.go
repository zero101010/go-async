package main

import "fmt"
// Thread 1
func main(){

	msg := make(chan string)
	// função anonima que cria a Thread 2
	go func(){
		// Passando valor para o canal por uma go routine
		msg<-"Hello world"
	}()

	result := <-msg
	fmt.Println(result)
}