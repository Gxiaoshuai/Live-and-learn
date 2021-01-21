package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
var total uint64
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total)
}

func worker (wg *sync.WaitGroup){
	defer  wg.Done()
	for i := 0;i< 100000;i++{
		atomic.AddUint64(&total,1)
	}
}
