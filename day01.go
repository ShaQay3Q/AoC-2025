package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	fmt.Println("AoC-2025")
	content := readFile("day01.txt")
    result := countZeroVisits(50, content, 100)
    fmt.Println(result)
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
		// fmt.Println("input:", input, "→ position:", position)  // trace each step
		if position == 0 {
			count++
		}
	}
	return count
}

func countZerosCrossedBackward(position int, steps int, dial int) int {
	if steps > position {
		return ((steps - position)/dial + 1)
	}
	return 0
}

func countZerosCrossedForward(position int, steps int, dial int) int {
	return ((position + steps) / 100)
}




func readFile(filename string) []string {
    file, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        if line != "" {  // skip empty lines
            lines = append(lines, line)
        }
    }
    return lines
}

