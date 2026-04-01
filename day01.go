package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main(){
	fmt.Println("AoC-2025")
	// content := readFile("day01.txt")
    // result := countZeroVisits(50, content, 100)
    // fmt.Println(result)
	fmt.Println(countZerosCrossedBackward(0, 5, 100))
    content := readFile("day01.txt")
    result := countAllZerosCrude(50, content, 100)
    fmt.Println(result)
}


func moveForward(position int,steps int, dial int) int{
	res := (position + steps)% dial;
	return res
}

func moveBackward(position int, steps int, dial int) int{
	res := (dial + (position - steps)) % dial
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

// func countZerosCrossedBackward(position, steps, dial int) int {
// 	// Special case: starting at 0
// 	if position == 0 {
// 		if steps == 0 {
// 			return 0
// 		}
// 		return (steps - 1) / dial
// 	}

// 	// If we don't even reach zero once
// 	if steps <= position {
// 		return 0
// 	}

// 	// First crossing + additional full wraps
// 	return 1 + (steps-position-1)/dial
// }

// func countZerosCrossedBackward(position int, steps int, dial int) int {
//     if position == 0 {
//         if steps < dial {
//             return 0
//         }
//         result := (steps - dial) / dial + 1
//         if steps%dial == 0 {
//             result--  // landing handled separately
//         }
//         return result
//     }
//     if steps > position {
//         result := (steps - position) / dial + 1
//         if (position-steps+dial)%dial == 0 {
//             result--
//         }
//         return result
//     }
//     return 0
// }

// func countZerosCrossedForward(position int, steps int, dial int) int {
//     if position == 0 && steps < dial {
//         return 0
//     }
//     result := (position + steps) / dial
//     if (position+steps)%dial == 0 {
//         result--
//     }
//     return result
// }
func countZerosCrossedForward(position int, steps int, dial int) int {
	return (position + steps) / dial
}

 func countZerosCrossedBackward(position int, steps int, dial int) int {
	result := (steps - position + dial) / dial

	if steps >= position && (steps-position)%dial == 0 {
		result--
	}

	return result
}
 

func countZerosDuringMoveCrude(position, steps, dial int, forward bool) int {
	count := 0

	for i := 0; i < steps; i++ {
		if forward {
			position = (position + 1) % dial
		} else {
			position = (position - 1 + dial) % dial
		}

		if position == 0 {
			count++
		}
	}

	return count
}

func countAllZerosCrude(position int, inputs []string, dial int) int {
	total := 0

	for _, input := range inputs {
		steps := getTheInt(input)
		forward := isForward(getTheFirstCharacter(input))

		total += countZerosDuringMoveCrude(position, steps, dial, forward)

		position = pointAt(position, input, dial)
	}

	return total
}

func countZerosDuringMove(position, steps, dial int, forward bool) int {
	count := 0

	for i := 0; i < steps; i++ {
		if forward {
			position = countZerosCrossedForward(position,steps, dial)
		} else {
			position = countZerosCrossedBackward(position, steps, dial)
		}

		if position == 0 {
			count++
		}
	}

	return count
}

func countAllCrossedZeros(position int, inputs []string, dial int) int {
	count := 0

	for _, input := range inputs {
		steps := getTheInt(input)
		forward := isForward(getTheFirstCharacter(input))

		if forward {
			count += countZerosCrossedForward(position, steps, dial)
			position = moveForward(position, steps, dial)
		} else {
			count += countZerosCrossedBackward(position, steps, dial)
			position = moveBackward(position, steps, dial)
		}
	}

	return count
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

