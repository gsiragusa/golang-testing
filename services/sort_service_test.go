package services

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// this is an integration test
func TestSort(t *testing.T) {
	elements := []int{5, 1, 6, 2, 3, 9, 8, 0, 7, 4}

	Sort(elements)

	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, elements)
}

func TestGolangSort(t *testing.T) {
	elements := []int{5, 1, 6, 2, 3, 9, 8, 0, 7, 4}

	GolangSort(elements)

	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, elements)
}
