package main

import (
	"fmt"
	"strings"
)

/*
* 3. Manacher 算法
*
*
*/

func main(){
	charlist := string("abladabcbaadslfskfakkkksldlwekldklsdfjj");
	fmt.Println(charlist)
	fmt.Println(len(charlist))
	var maxLen = manacher(charlist)
	fmt.Println(maxLen)
}
//3. Manacher 算法
// 算法核心思想
// 记录上一次扩展到最右边的回文串的位置 （中间状态pos、最右边字符maxRight）
// 如果  当前i位置在maxRight的左边，找出关于pos对称的位置 j 因为j已经计算过最大回文串，所以可以直接拿来用 但是需要判断是否超过了最右边maxRight，所以取 min(RL[pos-(i - pos)],maxRight - 1)最小的一个 然后开始比较、计数
func manacher(charList  string)(int){
	var charListBytes =  strings.Split(charList,"")
	charList = strings.Join(charListBytes,"#")
	charList = "#"+charList+"#"
	var RL = make([]int,len(charList))
	var maxRight int = 0
	var pos = 0
	var maxLen = 0
	for i,_ := range charList{
		if i < maxRight{
			// 如果 当前i位置在maxRight的左边，找出关于pos对称的位置 j 因为j已经计算过最大回文串，所以可以直接拿来用 但是需要判断是否超过了最右边maxRight，所以取 min(RL[pos-(i - pos)],maxRight - 1)最小的一个
			RL[i] = min(RL[2*pos - i],maxRight - 1)
		}else{
			RL[i] = 1
		}
		for{
			//无出界 情况 并且 左字符串等于右字符串，计算+1 否则就退出
			if i - RL[i] >=0 && i + RL[i] < len(charList) && charList[i - RL[i]] == charList[i+RL[i]]{
				RL[i] ++
			}else{
				break
			}
		}
		// 如果当前的右边界大于maxRight  重新记录maxRight 和pos
		if RL[i] + i -1 > maxRight{
			maxRight = RL[i] + i -1
			pos = i
		}
		maxLen = max(maxLen,RL[i])
		fmt.Println(i,maxLen,charList[i-RL[i]+1:pos+RL[i]])
	}
	return maxLen - 1
}
//
func min(x,y int) (int){
	if x < y {
		return x
	}
	return y
}
func max(x,y int)(int){
	if x > y {
		return x
	}
	return y
}