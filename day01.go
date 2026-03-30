package main

import (
	"fmt"
	"strconv"
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

func isForward(direction string) bool {
    return direction == "R"
}

func getTheFirstCharacter(input string)string{
	return string(input[0])
}

func getTheInt(input string) int {
	numStr := input[1:]
	num, _ := strconv.Atoi(numStr)
	return  num
}

// func pointAt(input string)int{
// 	if isForward(getTheFirstCharacter(input)){
// 		moveForward(0,)
// 	}
// 	return 1
// }
