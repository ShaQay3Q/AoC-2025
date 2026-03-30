package main

import (
	"fmt"
)

func main(){
	fmt.Println("AoC-2025")
}

func moveForward(initial_digit int,move_by int) int{
	res := initial_digit + move_by;
	return res
}