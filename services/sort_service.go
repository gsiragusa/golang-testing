package services

import "golang-testing/utils/sort"

func Sort(elements []int) {
	sort.BubbleSort(elements)
}

func GolangSort(elements []int) {
	sort.Sort(elements)
}
