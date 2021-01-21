package main

import "fmt"


var msg string
func main(){
	var done  = make(chan bool)
	go aGoroutine(done)
	for _ = range done{

	}
	fmt.Println(msg)
}
func aGoroutine(done chan bool){
	msg = "hello word"
	done <- true
	close(done)
}