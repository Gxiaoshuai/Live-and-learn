package main

import "fmt"

func main(){
	for i:=1;i<=3;i++{
		i := i
		defer func(){fmt.Println(i)}()
	}
}
