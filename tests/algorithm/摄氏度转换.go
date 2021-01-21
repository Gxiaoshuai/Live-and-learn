package main

import "fmt"

const boilingF = 212.0
func main(){
	var f  = boilingF
	var c = (f - 32) * 5/9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	const freezingF,boilingF = 32.0,212.0
	fmt.Printf("%.1f°F = %.1f°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%.1f°F = %.1f°C\n", freezingF, fToC(boilingF))
}

func fToC(f float64)(float64){
	return (f - 32) * 5 / 9
}