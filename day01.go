package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("AoC-2025")
}

func moveForward(initial_digit int,steps int, dial int) int{
	res := (initial_digit + steps)% dial;
	return res
}

func moveBackward(initial_digit int, steps int, dial int) int{
	res := (dial + (initial_digit - steps)) % dial
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

func pointAt(position int, input string, dial int)int{
	if isForward(getTheFirstCharacter(input)){
		return moveForward(position,getTheInt(input), dial)
	}
	return moveBackward(position,getTheInt(input), dial)
}

func pointAtAll(position int, inputs []string, dial int) int {
	for _, input := range inputs {
		position = pointAt(position, input, dial)
	}
	return position
}
func countZeroVisits(position int, inputs []string, dial int) int {
	count := 0
	for _, input := range inputs {
		position = pointAt(position, input, dial)
		if position == 0 {
			count++
		}
	}
	return count
}


