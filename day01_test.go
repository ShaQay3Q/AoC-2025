package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMoveForward01(t *testing.T) {
    // start at 0, move 1
    result := moveForward(0, 1)
		require.Equal(t, 1, result)
}

func TestMovingForward02(t *testing.T) {
	result := moveForward(0, 1)
		require.Equal(t, 1, result)
}

func TestMovingForward03(t *testing.T) {
	result := moveForward(1, 1)
		require.Equal(t, 2, result)
}


func TestMovingForward04(t *testing.T) {
	result := moveForward(9, 1)
		require.Equal(t, 0, result)
}

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

