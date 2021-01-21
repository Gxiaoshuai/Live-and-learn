package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

func main(){
	var jsonStr string = "{\"test\":\"a\",\"a\":{\"cc\":1,\"b\":2,\"cc\":{\"ccc\":1,\"bb\":{\"ccc\":\"\"}}}}"
	var jsonObj  interface{}
	err := json.Unmarshal([]byte(jsonStr),&jsonObj)
	if err !=nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	depLen := getDepLen(jsonObj)
	fmt.Println(depLen)
	//info(jsonObj)


}
func getDepLen(data interface{})int{
	var depLen int = 1
	var tempLen int = 1
	var maxLen int  = 1
	var refType = reflect.TypeOf(data)
	if refType.Kind() != reflect.Map && refType.Kind()!= reflect.Slice && refType.Kind() != reflect.Struct{
		return depLen
	}
	var refVal = reflect.ValueOf(data)
	var tempRefType  reflect.Type
	if refType.Kind() == reflect.Struct{
		for i:= 0;i < refType.NumField();i++{
			tempRefType = reflect.TypeOf(refVal.Field(i).Pointer())

			if tempRefType.Kind()== reflect.Map || tempRefType.Kind()== reflect.Slice || tempRefType.Kind() == reflect.Struct{
				tempLen = getDepLen(refVal.Field(i).Interface())
			}else {
				maxLen = 1
			}
			if tempLen > maxLen{
				maxLen = tempLen
			}
		}
	}else{
		key := refVal.MapRange()
		for key.Next() {
			tempRefType = reflect.TypeOf(key.Value())

			if tempRefType.Kind()== reflect.Map || tempRefType.Kind()== reflect.Slice || tempRefType.Kind() == reflect.Struct{
				tempLen = getDepLen(key.Value())
			}else {
				maxLen = 1
			}
			if tempLen > maxLen{
				maxLen = tempLen
			}
			fmt.Println(tempLen)
		}
	}
	depLen += maxLen
	fmt.Println(maxLen,depLen)
	return depLen
}
