package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// a general testing functions for all the forward-going digit tests
func TestMovingForward(t *testing.T){
	tests := []struct{
		position int
		steps int
		expectation int
	}{
		{0, 1, 1},
        {0, 1, 1},
        {1, 1, 2},
        {9, 1, 0},
	}
	for _, tt := range tests{
		result := moveForward(tt.position, tt.steps)
		require.Equal(t, tt.expectation, result)
	}
}

func TestMoveBackward(t *testing.T){
	tests := []struct{
		position int
		steps int
		expectation int
	}{
		{1, 1, 0},
        {2, 1, 1},
        {0, 1, 9},
	}
	for _, tt := range tests{
		result := moveBackward(tt.position, tt.steps)
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


// func TestPoinAt01(t *testing.T){
// 	result := pointAt("R1")
// 	require.Equal(t,1,result)
// }



// func TestPoinAt02(t *testing.T){
// 	result := pointAt("R2")
// 	require.Equal(t,2,result)
// }