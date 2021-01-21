package main

import "fmt"

type People interface {
	Speak (string)string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

func main() {
	var peo People = & Student{}
	// 或者使用new来创建 var peo People = new (Student)
	think := "speak"
	fmt.Println(peo.Speak(think))
}
/**


*/
