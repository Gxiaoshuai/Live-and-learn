package main

import (
	"fmt"
	"math/rand"
)

func main(){
	var list = []int {93,402,392,83,9,42,42,383,39,32,12,43,43,78,34,99,191}//声明一个乱序数组
	list = fastSort(list) //调用递归排序方法
	fmt.Println(list)
}
func fastSort(arr []int)([]int){
	if(len(arr)<= 1){// 如果传入的数组只有一个元素  直接返回当前数组
		return arr
	}
	var leftArr = make([]int,0)// 声明两个空的切片
	var rightArr = make([]int,0)
	var randIndex = rand.Intn(len(arr));//获取一个随机的index下标
	var middleVal = arr[randIndex]// 取出数组中下标对应的值 作为中间值
	for key,value := range arr{ //循环判断数组中的元素 跟中间值比大小
		if(key == randIndex){
			continue
		}
		if(value <= middleVal ){// 小的值存左边切片，大的值在右边切片
			leftArr = append(leftArr,value)
		}else{
			rightArr = append(rightArr,value)
		}
	}
	//fmt.Println(leftArr)
	//fmt.Println(rightArr)
	leftArr = fastSort(leftArr) //递归排左边的数组
	rightArr = fastSort(rightArr) //递归排右边的数组
	return append(append(leftArr,middleVal),rightArr...) //返回整个数组
}
