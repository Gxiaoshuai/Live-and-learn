package main

import (
	"fmt"
	"time"
)

/**
两个协程，交叉打印：1A2B3C4D....10I
 */
func main(){
	var index int = 1
	go func(index *int){
		var char uint8 = 65
		for{
			if char > 73 {
				break
			}
			if *index % 2 == 0 {
				fmt.Printf("%c",char)
				char ++
				*index ++
			}
		}
	}(&index)

	go func(index *int){
		var char int = 1
		for{
			if char > 9 {
				break
			}
			if (*index+1) % 2 == 0 {
				fmt.Println(char)
				char ++
				*index ++
			}
		}
	}(&index)
	time.Sleep(1 * time.Second)
}
