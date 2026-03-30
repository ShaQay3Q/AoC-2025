package main

import (
	"fmt"
)

func main(){
	fmt.Println("AoC-2025")
}

func moveForward(initial_digit int,steps int) int{
	res := (initial_digit + steps)%10;
	return res
}

func moveBackward(initial_digit int, steps int) int{
	res := (10 + (initial_digit - steps)) % 10
	return res
}