package main

import (
	"learn1/bubble"
	"learn1/insertion"
	"learn1/selection"
	"reflect"
	"runtime"
	"testing"
)

type Sorter func([]int) []int

var testCases = []struct {
	name string
	args []int
	want []int
}{
	{
		name: "Test 1",
		args: []int{5, 3, 1, 4, 2},
		want: []int{1, 2, 3, 4, 5},
	},
	{
		name: "Test 2",
		args: []int{10, 20, 30, 40, 50},
		want: []int{10, 20, 30, 40, 50},
	},
	{
		name: "Test 3",
		args: []int{100, 50, 25, 75, 0},
		want: []int{0, 25, 50, 75, 100},
	},
	{
		name: "Test 4",
		args: []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000},
		want: []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000},
	},
	{
		name: "Test 5",
		args: []int{1000, 900, 800, 700, 600, 500, 400, 300, 200, 100},
		want: []int{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000},
	},
	{
		name: "Test 6",
		args: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		want: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	},
	{
		name: "Test 7",
		args: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		want: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10},
	},
	{
		name: "Test 8",
		args: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
		want: []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6, 7, 7, 8, 8, 9, 9, 10, 10},
	},
}

var sorters = []Sorter{
	insertion.Sort,
	selection.Sort,
	bubble.Sort,
}

func TestSort(t *testing.T) {
	for _, sorter := range sorters {
		sorterName := runtime.FuncForPC(reflect.ValueOf(sorter).Pointer()).Name()
		t.Run(sorterName, func(t *testing.T) {
			for _, testCase := range testCases {

				// Create a copy of tt.args
				argsCopy := make([]int, len(testCase.args))
				copy(argsCopy, testCase.args)

				t.Run(testCase.name, func(t *testing.T) {
					if got := sorter(argsCopy); !reflect.DeepEqual(got, testCase.want) {
						t.Errorf("SortInts() = %v, want %v", got, testCase.want)
					}
				})
			}
		})
	}
}
