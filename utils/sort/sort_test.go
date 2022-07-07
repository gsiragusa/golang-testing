package sort

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// this is a unit test
func TestBubbleSort(t *testing.T) {
	elements := []int{5, 1, 6, 2, 3, 9, 8, 0, 7, 4}

	BubbleSort(elements)

	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, elements)
}

func TestSort(t *testing.T) {
	elements := []int{5, 1, 6, 2, 3, 9, 8, 0, 7, 4}

	Sort(elements)

	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, elements)
}

func TestBubbleSortWithTimeout(t *testing.T) {
	elements := []int{5, 1, 6, 2, 3, 9, 8, 0, 7, 4}

	timeoutChan := make(chan bool, 1)
	defer close(timeoutChan)

	go func() {
		BubbleSort(elements)
		timeoutChan <- false
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		timeoutChan <- true
	}()

	if <-timeoutChan {
		require.Fail(t, "bubble sort took more than 500ms")
	}

	require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, elements)
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := []int{5, 1, 6, 2, 3, 9, 8, 0, 7, 4}

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := []int{5, 1, 6, 2, 3, 9, 8, 0, 7, 4}

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
