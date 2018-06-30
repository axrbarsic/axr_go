package main

import (
	"fmt"
)

func main() {
	three := "Fizz" //переменная кратная трем
	five := "Buzz" //переменная кратная пяти
	i := 1
	for i < 100 {
		i = i + 1
		if i % 3 == 0 { //условие кратное трем
			i := three
			fmt.Println(i)
		} else {
			fmt.Println(i)
			if i % 5 == 0 { // условие кратное пяти
				i := five
			fmt.Println(i)
		} else {
			fmt.Println(three+five)
		}
		}
	}
}



//if i % 3 == 0 {
//fmt.Println(i, "Fizz")
//} else {
//fmt.Println(i, "не четное")