package main

import "fmt"

func main(){
    var fullNumChan chan int = MakeFullNum(100)
    //循环取素数
    for{
    	prime,ok := <- fullNumChan //因为第一个数为2  为素数 所以直接取出来 然后去筛
    	if(ok == false ){
    		break
		}
		fmt.Println(prime,ok)
    	fullNumChan = PrimeFilter(fullNumChan,prime)// 取数、筛数,并且将筛完后的chan继续赋值给fullNumChan  继续下次循环筛
	}
}
// 生成所有自然数
func MakeFullNum(max int) (chan int){
	var fullNumChan = make(chan int)
	var i int
	go func() {
		for  i = 2;i<=max;i++ {
			fullNumChan <- i
		}
		close(fullNumChan)
	}()
	return  fullNumChan
}
//对channel中所有的数  过滤掉所有非素数
func PrimeFilter (in <- chan  int,prime int ) chan int{
	out := make(chan int)
	go func() {
		for{
			i,ok := <- in
			if(!ok){
				break;
			}
			if  ok&&i%prime !=0 {
				out <- i
			}
		}
		close(out)
	}()
	return out
}