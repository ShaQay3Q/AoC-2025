package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// a general testing functions for all the forward-going digit tests
func TestMovingForward(t *testing.T){
	tests := []struct{
		position int
		steps int
		expectation int
		dial int
	}{
		{0, 1, 1, 10},
        {0, 1, 1, 10},
        {1, 1, 2, 10},
        {9, 1, 0, 10},
	}
	for _, tt := range tests{
		result := moveForward(tt.position, tt.steps, tt.dial)
		require.Equal(t, tt.expectation, result)
	}
}

func TestMoveBackward(t *testing.T){
	tests := []struct{
		position int
		steps int
		expectation int
		dial int
	}{
		{1, 1, 0, 10},
        {2, 1, 1, 10},
        {0, 1, 9, 10},
	}
	for _, tt := range tests{
		result := moveBackward(tt.position, tt.steps, tt.dial)
		require.Equal(t, tt.expectation, result)
	}
}

func TestIsForward(t *testing.T) {
    result := isForward("R")
    require.True(t, result)
}

func TestGetTheFirstCharacter(t *testing.T){
	tests :=[]struct{
		input string
		expectation string
	}{
		{"R3", "R"},
		{"L3", "L"},
	}
			for _, tt := range tests{
			result := getTheFirstCharacter(tt.input)
			require.Equal(t, tt.expectation, result)
		}
}

func TestGetTheInt(t *testing.T){
	tests := []struct{
		input string
		expectation int
	}{
		{"R3", 3},
		{"R1", 1},
		{"L6", 6},
		{"L1", 1},
	}
	for _, tt := range tests {
		result := getTheInt(tt.input)
		require.Equal(t, tt.expectation, result)
	}
}

func TestPoinAt(t *testing.T){
	tests := []struct{
		input string
		expectation int
		dial int
	}{
		{"R1", 1, 10},
		{"R2", 2, 10},
		{"L3", 7, 10},
		{"L1", 9, 10},
	}
	for _, tt := range tests {
		result := pointAt(0, tt.input, tt.dial)
		require.Equal(t, tt.expectation, result)
	}
}

func TestPointAtAll(t *testing.T) {
	tests := []struct{
		position int
		inputs []string
		expectation int
		dial int
	}{
		{0, []string{"R1", "L4", "L5", "R6"}, 8, 10},
		{0, []string{"R6", "L2"}, 4, 10},
	}
	for _, tt := range tests{
		result := pointAtAll(tt.position, tt.inputs, tt.dial)
		require.Equal(t, tt.expectation, result)
	}
}

func TestCountZeroVisits(t *testing.T){
   tests := []struct {
		position int
        inputs      []string
        expectation int
		dial int
    }{
        {0, []string{"R1", "L1"}, 1, 10},
        {0, []string{"R1", "L1", "R2", "L2"}, 2, 10},
		{0, []string{"R1", "L1"}, 1, 100},
		{2, []string{"R1", "L2"}, 0, 10},
		{50, []string{"L68","L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}, 3, 100},
    }
		for _, tt := range tests {
			count := countZeroVisits(tt.position, tt.inputs, tt.dial)
			require.Equal(t, tt.expectation, count)
		}
}

func TestReadFile(t *testing.T) {
    lines := readFile("day01.txt")
    require.Greater(t, len(lines), 0)  // at least one line was read
    fmt.Println(lines[0])              // print first line to see what it looks like
}

func TestCountZerosCrossedBackward(t *testing.T){
	tests := []struct{
		position int
		steps int
		dial int
		expectation int
	}{
		{50, 68, 100, 1},
		{50, 30, 100, 0},
		{0, 99, 100, 1},
		{0, 118, 100, 2},
		{95, 5, 100, 0},
	}
	for _, tt := range tests{
		result := countZerosCrossedBackward(tt.position, tt.steps, tt.dial)
		require.Equal(t, tt.expectation, result)
	}
}

func TestCountZerosCrossedForward(t *testing.T){
	tests := []struct{
		position int
		steps int
		dial int
		expectation int
	}{
		{50, 68, 100, 1},
		{50, 30, 100, 0},
		{0, 99, 100, 0},
		{0, 1018, 100, 10},
	}
	for _, tt := range tests{
		result := countZerosCrossedForward(tt.position, tt.steps, tt.dial)
		require.Equal(t, tt.expectation, result)
	}
}

func TestCountAllCrossedZeros(t *testing.T){
	tests := []struct {
		position int
        inputs      []string
        expectation int
		dial int
    }{
        // {0, []string{"R1", "L2"}, 1, 10},
        // {0, []string{"R1", "L2", "R2", "L2"}, 3, 10},
		// {0, []string{"R1", "L2"}, 1, 100},
		// {2, []string{"R1", "L2"}, 0, 10},
		// {50, []string{"L68", "L30", "R48"}, 2, 100},
		{50, []string{"L68","L30", "R48", "L5", "R60", "L55", "L1", "L99", "R14", "L82"}, 6, 100},
	}
	for _, tt := range tests{
		// for _, entry := range tt.inputs{
		// 	if isForward(entry){
		// 		coun
		// 	}
		// }
		result := countAllCrossedZeros(tt.position, tt.inputs, tt.dial)
		require.Equal(t, tt.expectation, result)
	}
}

func TestDebugTrace(t *testing.T) {
    type step struct {
        input    string
        position int
        crossed  int
        landed   int
        total    int
    }

    expected := []step{
        {"L68", 82, 1, 0, 1},
        {"L30", 52, 0, 0, 1},
        {"R48", 0,  0, 1, 2},
        {"L5",  95, 0, 0, 2},
        {"R60", 55, 1, 0, 3},
        {"L55", 0,  0, 1, 4},
        {"L1",  99, 0, 0, 4},
        {"L99", 0,  0, 1, 5},
        {"R14", 14, 0, 0, 5},
        {"L82", 32, 1, 0, 6},
    }

    position := 50
    dial := 100
    count := 0

    for i, tt := range expected {
        steps := getTheInt(tt.input)
        direction := getTheFirstCharacter(tt.input)

        crossed := 0
        if isForward(direction) {
            crossed = countZerosCrossedForward(position, steps, dial)
        } else {
            crossed = countZerosCrossedBackward(position, steps, dial)
        }
        count += crossed
        position = pointAt(position, tt.input, dial)
        landed := 0
        if position == 0 {
            landed = 1
            count++
        }
fmt.Printf("step %d: input=%q direction=%q steps=%d\n", i+1, tt.input, direction, steps)
        fmt.Printf("step %d: %s pos:%d crossed:%d landed:%d total:%d\n", i+1, tt.input, position, crossed, landed, count)

        require.Equal(t, tt.position, position,  "position wrong at step %d", i+1)
        require.Equal(t, tt.crossed,  crossed,   "crossed wrong at step %d",  i+1)
        require.Equal(t, tt.landed,   landed,     "landed wrong at step %d",   i+1)
        require.Equal(t, tt.total,    count,      "total wrong at step %d",    i+1)
    }
}