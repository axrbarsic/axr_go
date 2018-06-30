package main

import "fmt"

func main()  {
	fmt.Print("Введите в меня фаренгейт и получите цельсий: ")
	var f float64
	fmt.Scanf("%F", &f)
	c := (f - 32) *5/9
	//var c float64 = c bool()
	fmt.Println(c)
}

// (C = (F - 32) * 5/9)