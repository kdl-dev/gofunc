package gofunc_test

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/kdl-dev/gofunc"
	"github.com/stretchr/testify/require"
)

// ! don't edit
var (
	intSlice  = []int{3, 1, 6, 9, -5, 0, 11}
	strSlice  = []string{"test2", "test1", "test4", "test3"}
	intSlice2 = []int{3, 1, 6, 9, -5, 0, 11, 1, 1, 5, 9, 14, 3, -5, 0}
	strSlice2 = []string{"test2", "test1", "test4", "test3", "test2", "test5", "test1"}
)

type TestNewStruct[T comparable] struct {
	description string
	input       T
	expected    interface{}
}

func TestNew(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test init collection of ints",
			input:       intSlice,
			expected:    intSlice,
		},
		{
			description: "test init collection of strings",
			input:       strSlice,
			expected:    strSlice,
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			require.Equal(t, test.expected, collection.ToSlice())
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			require.Equal(t, test.expected, collection.ToSlice())
		}
	}
}

func TestMap(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test map for collection of ints",
			input:       intSlice,
			expected:    []int{6, 2, 12, 18, -10, 0, 22},
		},
		{
			description: "test map for collection of strings",
			input:       strSlice,
			expected:    []string{"TEST2", "TEST1", "TEST4", "TEST3"},
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			newCollection := collection.Map(func(el int) int { return el * 2 })
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			newCollection := collection.Map(func(el string) string { return strings.ToUpper(el) })
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		}
	}
}

func TestReduce(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test reduce for collection of ints",
			input:       intSlice,
			expected:    25,
		},
		{
			description: "test reduce for collection of strings",
			input:       strSlice,
			expected:    "test2 test1 test4 test3 ",
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			result := collection.Reduce(func(el, accum int) int { return accum + el })
			require.Equal(t, test.expected, result)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			result := collection.Reduce(func(el, accum string) string { return accum + el + " " })
			require.Equal(t, test.expected, result)
		}
	}
}

func TestFilter(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test filter for collection of ints",
			input:       intSlice,
			expected:    []int{6, 0},
		},
		{
			description: "test filter for collection of strings",
			input:       strSlice,
			expected:    []string{"test2", "test4"},
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			newCollection := collection.Filter(func(el int) bool { return el%2 == 0 })
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			newCollection := collection.Filter(func(el string) bool { return strings.Contains(el, "2") || strings.Contains(el, "4") })
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		}
	}
}

func TestMatch(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test match for collection of ints",
			input:       intSlice,
			expected:    true,
		},
		{
			description: "test match for collection of strings",
			input:       strSlice,
			expected:    false,
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			result := collection.Match(func(el int) bool { return el < 0 })
			require.Equal(t, test.expected, result)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			result := collection.Match(func(el string) bool { return el == "gofunc" })
			require.Equal(t, test.expected, result)
		}
	}
}

func TestAllMatch(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test all match for collection of ints",
			input:       intSlice,
			expected:    false,
		},
		{
			description: "test all match for collection of strings",
			input:       strSlice,
			expected:    true,
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			result := collection.AllMatch(func(el int) bool { return el > 0 })
			require.Equal(t, test.expected, result)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			result := collection.AllMatch(func(el string) bool { return strings.Contains(el, "test") })
			require.Equal(t, test.expected, result)
		}
	}
}

func TestDistinct(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test distinct for collection of ints",
			input:       intSlice2,
			expected:    []int{3, 1, 6, 9, -5, 0, 11, 5, 14},
		},
		{
			description: "test distinct for collection of strings",
			input:       strSlice2,
			expected:    []string{"test2", "test1", "test4", "test3", "test5"},
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			newCollection := collection.Distinct()
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			newCollection := collection.Distinct()
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		}
	}
}

func TestLimit(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test limit for collection of ints",
			input:       intSlice,
			expected:    intSlice[0:3],
		},
		{
			description: "test limit for collection of strings",
			input:       strSlice,
			expected:    strSlice[0:3],
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			newCollection := collection.Limit(3)
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())

			collection.Limit(collection.Len() + 10)
			collection.Limit(-1)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			newCollection := collection.Limit(3)
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		}
	}
}

func TestSkip(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test skip for collection of ints",
			input:       intSlice,
			expected:    intSlice[3:],
		},
		{
			description: "test skip for collection of strings",
			input:       strSlice,
			expected:    strSlice[3:],
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			newCollection := collection.Skip(3)
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())

			collection.Skip(collection.Len() + 10)
			collection.Skip(-1)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			newCollection := collection.Skip(3)
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		}
	}
}

func TestSort(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test sort for collection of ints",
			input:       intSlice,
			expected:    []int{-5, 0, 1, 3, 6, 9, 11},
		},
		{
			description: "test sort for collection of strings",
			input:       strSlice,
			expected:    []string{"test1", "test2", "test3", "test4"},
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			newCollection := collection.Sort(func(arr []int) { sort.Ints(arr) })
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			newCollection := collection.Sort(func(arr []string) { sort.Strings(arr) })
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		}
	}
}

func TestReverse(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test reverse for collection of ints",
			input:       intSlice,
			expected:    []int{11, 0, -5, 9, 6, 1, 3},
		},
		{
			description: "test reverse for collection of strings",
			input:       strSlice,
			expected:    []string{"test3", "test4", "test1", "test2"},
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			newCollection := collection.Reverse()
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			newCollection := collection.Reverse()
			require.Equal(t, test.expected, newCollection.ToSlice())
			require.NotEqual(t, test.expected, collection.ToSlice())
		}
	}
}

func TestMax(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test max for collection of ints",
			input:       intSlice,
			expected:    11,
		},
		{
			description: "test max for collection of strings",
			input:       strSlice,
			expected:    "test4",
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			result := collection.Max(func(firstEl, secondEl int) int { return int(math.Max(float64(firstEl), float64(secondEl))) })
			require.Equal(t, test.expected, result)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			result := collection.Max(func(firstEl, secondEl string) string {
				if firstEl > secondEl {
					return firstEl
				}
				return secondEl

			})
			require.Equal(t, test.expected, result)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test min for collection of ints",
			input:       intSlice,
			expected:    -5,
		},
		{
			description: "test min for collection of strings",
			input:       strSlice,
			expected:    "test1",
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			result := collection.Min(func(firstEl, secondEl int) int { return int(math.Min(float64(firstEl), float64(secondEl))) })
			require.Equal(t, test.expected, result)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			result := collection.Min(func(firstEl, secondEl string) string {
				if firstEl < secondEl {
					return firstEl
				}
				return secondEl

			})
			require.Equal(t, test.expected, result)
		}
	}
}

func TestLen(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test len for collection of ints",
			input:       intSlice,
			expected:    7,
		},
		{
			description: "test len for collection of strings",
			input:       strSlice,
			expected:    4,
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			result := collection.Len()
			require.Equal(t, test.expected, result)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			result := collection.Len()
			require.Equal(t, test.expected, result)
		}
	}
}

func TestToSlice(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test conversion to slice of ints for collection of ints",
			input:       intSlice,
			expected:    intSlice,
		},
		{
			description: "test conversion to slice of strings for collection of strings",
			input:       strSlice,
			expected:    strSlice,
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			result := collection.ToSlice()
			require.Equal(t, test.expected, result)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			result := collection.ToSlice()
			require.Equal(t, test.expected, result)
		}
	}
}

func TestToString(t *testing.T) {
	tests := []TestNewStruct[interface{}]{
		{
			description: "test conversion to string for collection of ints",
			input:       intSlice,
			expected:    "3 1 6 9 -5 0 11 ",
		},
		{
			description: "test conversion to string for collection of strings",
			input:       strSlice,
			expected:    "test2 test1 test4 test3 ",
		},
	}

	for _, test := range tests {
		t.Logf("Start: %s\n", test.description)
		switch test.input.(type) {
		case []int:
			slice := test.input.([]int)
			collection := gofunc.New(slice)
			result := collection.ToString(func(el int) string { return strconv.Itoa(el) + " " })
			require.Equal(t, test.expected, result)
		case []string:
			slice := test.input.([]string)
			collection := gofunc.New(slice)
			result := collection.ToString(func(el string) string { return el + " " })
			require.Equal(t, test.expected, result)
		}
	}
}
