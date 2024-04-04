package main

import (
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
}

var sorters = []Sorter{
	insertion.Sort,
	selection.Sort,
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
