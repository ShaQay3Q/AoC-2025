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