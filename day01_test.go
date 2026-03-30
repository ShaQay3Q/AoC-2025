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