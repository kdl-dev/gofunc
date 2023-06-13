package gofunc_test

import (
	"strings"
	"testing"

	"github.com/kdl-dev/gofunc"
	"github.com/stretchr/testify/require"
)

var (
	intSlice = []int{3, 1, 6, 9, -5, 0, 11}
	strSlice = []string{"test1", "test2", "test3", "test4"}
)

type TestNewStruct[T comparable] struct {
	description string
	input       T
	expected    T
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
			expected:    []string{"TEST1", "TEST2", "TEST3", "TEST4"},
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
			expected:    "test1 test2 test3 test4 ",
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
