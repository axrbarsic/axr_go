//Напишите программу, которая выводит числа от 1 до 100, которые делятся на 3. (3, 6, 9, …).
//Напишите программу, которая выводит числа от 1 до 100. Но для кратных трём нужно вывести «Fizz» вместо числа, для кратных пяти - «Buzz», 
//а для кратных как трём, так и пяти — «FizzBuzz».



package main

import "fmt"

func main() {
	i :=1
	for i <= 100 {
		if i % 3 == 0 && i % 5 == 0 {  //тоже пойдет if i % 15 == 0 {
			fmt.Println(i, "FizzBuzz", "число делится на 3 и 5")
		} else if i % 5 == 0 {
			fmt.Println(i,"Buzz", "число делится на 5")
		} else  if i % 3 == 0 {
			fmt.Println(i,"Fizz", "число делится на 3")
			} else {
				fmt.Println(i, "простые числа")
		}
		i = i + 1
	}
}
