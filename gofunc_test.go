package gofunc

import (
	"math"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected interface{}
	}{
		{
			name:     "test1",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "test2",
			input:    []int{},
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		slice := test.input
		collection := New(slice)
		require.Equal(t, test.expected, collection.ToSlice())
	}
}

func TestGenerate(t *testing.T) {
	var i int
	tests := []struct {
		name     string
		script   func() int
		limit    int
		expected interface{}
	}{
		{
			name:     "test1",
			script:   func() int { i++; return i },
			limit:    5,
			expected: New([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "test2",
			script:   func() int { i++; return i },
			limit:    0,
			expected: New(make([]int, 0)),
		},
		{
			name:     "test3",
			script:   func() int { i++; return i },
			limit:    -1,
			expected: New(make([]int, 0)),
		},
		{
			name:     "test4",
			script:   nil,
			limit:    5,
			expected: (*collection[int])(nil),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		i = 0
		collection := Generate[int](test.script, test.limit)
		require.Equal(t, test.expected, collection)
	}
}

func TestMap(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func(int) int
		expected *collection[int]
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(i int) int { return i + 1 },
			expected: New([]int{2, 3, 4, 5, 6}),
		},
		{
			name:     "test2",
			input:    New([]int{}),
			script:   func(i int) int { return i + 1 },
			expected: New([]int{}),
		},
		{
			name:     "test3",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   nil,
			expected: New([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		collection := test.input.Map(test.script)
		require.Equal(t, test.expected, collection)
	}
}

func TestFlatMap(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func(int) (int, int)
		expected *collection[int]
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(i int) (int, int) { return i, i + 1 },
			expected: New([]int{1, 2, 2, 3, 3, 4, 4, 5, 5, 6}),
		},
		{
			name:     "test2",
			input:    New([]int{}),
			script:   func(i int) (int, int) { return i, i + 1 },
			expected: New([]int{}),
		},
		{
			name:     "test3",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   nil,
			expected: New([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		collection := test.input.FlatMap(test.script)
		require.Equal(t, test.expected, collection)
	}
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func(int, int) int
		expected int
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(el, accum int) int { return el + accum },
			expected: 15,
		},
		{
			name:     "test2",
			input:    New([]int{}),
			script:   func(el, accum int) int { return el + accum },
			expected: 0,
		},
		{
			name:     "test3",
			input:    New([]int{}),
			script:   nil,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.Reduce(test.script)
		require.Equal(t, test.expected, result)

	}
}

func TestFilter(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func(int) bool
		expected *collection[int]
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(i int) bool { return i%2 == 0 },
			expected: New([]int{2, 4}),
		},
		{
			name:     "test2",
			input:    New([]int{}),
			script:   func(i int) bool { return i%2 == 0 },
			expected: New([]int{}),
		},
		{
			name:     "test3",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   nil,
			expected: New([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		collection := test.input.Filter(test.script)
		require.Equal(t, test.expected, collection)
	}
}

func TestMatch(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func(int) bool
		expected bool
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(i int) bool { return i == 4 },
			expected: true,
		},
		{
			name:     "test2",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(i int) bool { return i == 10 },
			expected: false,
		},
		{
			name:     "test3",
			input:    New([]int{}),
			script:   func(i int) bool { return i == 0 },
			expected: false,
		},
		{
			name:     "test4",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   nil,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.Match(test.script)
		require.Equal(t, test.expected, result)
	}
}

func TestAllMatch(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func(int) bool
		expected bool
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(i int) bool { return i > 0 },
			expected: true,
		},
		{
			name:     "test2",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(i int) bool { return i%2 == 0 },
			expected: false,
		},
		{
			name:     "test3",
			input:    New([]int{}),
			script:   func(i int) bool { return i > 0 },
			expected: false,
		},
		{
			name:     "test4",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   nil,
			expected: false,
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.AllMatch(test.script)
		require.Equal(t, test.expected, result)
	}
}

func TestDistinct(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		expected *collection[int]
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
			expected: New([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "test2",
			input:    New([]int{}),
			expected: New([]int{}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.Distinct()
		require.Equal(t, test.expected, result)
	}
}

func TestLimit(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		limit    int
		expected *collection[int]
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			limit:    5,
			expected: New([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "test2",
			input:    New([]int{}),
			limit:    5,
			expected: New([]int{}),
		},
		{
			name:     "test3",
			input:    New([]int{1, 2, 3, 4, 5}),
			limit:    -1,
			expected: New([]int{}),
		},
		{
			name:     "test4",
			input:    New([]int{1, 2, 3, 4, 5}),
			limit:    10,
			expected: New([]int{1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.Limit(test.limit)
		require.Equal(t, test.expected, result)
	}
}

func TestSkip(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		skip     int
		expected *collection[int]
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			skip:     5,
			expected: New([]int{6, 7, 8, 9, 10}),
		},
		{
			name:     "test2",
			input:    New([]int{}),
			skip:     5,
			expected: New([]int{}),
		},
		{
			name:     "test3",
			input:    New([]int{1, 2, 3, 4, 5}),
			skip:     -1,
			expected: New([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "test4",
			input:    New([]int{1, 2, 3, 4, 5}),
			skip:     10,
			expected: New([]int{}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.Skip(test.skip)
		require.Equal(t, test.expected, result)
	}
}

func TestSort(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func([]int)
		expected *collection[int]
	}{
		{
			name:     "test1",
			input:    New([]int{2, 1, 4, 3, 5}),
			script:   func(arr []int) { sort.Ints(arr) },
			expected: New([]int{1, 2, 3, 4, 5}),
		},
		{
			name:     "test2",
			input:    New([]int{}),
			script:   func(arr []int) { sort.Ints(arr) },
			expected: New([]int{}),
		},
		{
			name:     "test3",
			input:    New([]int{2, 1, 4, 3, 5}),
			script:   nil,
			expected: New([]int{2, 1, 4, 3, 5}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		collection := test.input.Sort(test.script)
		require.Equal(t, test.expected, collection)
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		expected *collection[int]
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			expected: New([]int{5, 4, 3, 2, 1}),
		},
		{
			name:     "test2",
			input:    New([]int{}),
			expected: New([]int{}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		collection := test.input.Reverse()
		require.Equal(t, test.expected, collection)
	}
}

func TestReplace(t *testing.T) {
	tests := []struct {
		name        string
		input       *collection[int]
		targets     []int
		replacement int
		expected    *collection[int]
	}{
		{
			name:        "test1",
			input:       New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
			targets:     []int{1, 2, 3},
			replacement: 10,
			expected:    New([]int{10, 10, 10, 4, 5, 1, 2, 3, 4, 5}),
		},
		{
			name:        "test2",
			input:       New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
			targets:     []int{7, 8, 9},
			replacement: 10,
			expected:    New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
		},
		{
			name:        "test3",
			input:       New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
			targets:     nil,
			replacement: 10,
			expected:    New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		collection := test.input.Replace(test.targets, test.replacement)
		require.Equal(t, test.expected, collection)
	}
}

func TestReplaceAll(t *testing.T) {
	tests := []struct {
		name        string
		input       *collection[int]
		targets     []int
		replacement int
		expected    *collection[int]
	}{
		{
			name:        "test1",
			input:       New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
			targets:     []int{1, 2, 3},
			replacement: 10,
			expected:    New([]int{10, 10, 10, 4, 5, 10, 10, 10, 4, 5}),
		},
		{
			name:        "test2",
			input:       New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
			targets:     []int{7, 8, 9},
			replacement: 10,
			expected:    New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
		},
		{
			name:        "test3",
			input:       New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
			targets:     nil,
			replacement: 10,
			expected:    New([]int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}),
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		collection := test.input.ReplaceAll(test.targets, test.replacement)
		require.Equal(t, test.expected, collection)
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func(int, int) int
		expected int
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(el1, el2 int) int { return int(math.Max(float64(el1), float64(el2))) },
			expected: 5,
		},
		{
			name:     "test2",
			input:    New([]int{}),
			script:   func(el1, el2 int) int { return int(math.Max(float64(el1), float64(el2))) },
			expected: 0,
		},
		{
			name:     "test3",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   nil,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.Max(test.script)
		require.Equal(t, test.expected, result)

	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func(int, int) int
		expected int
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(el1, el2 int) int { return int(math.Min(float64(el1), float64(el2))) },
			expected: 1,
		},
		{
			name:     "test2",
			input:    New([]int{}),
			script:   func(el1, el2 int) int { return int(math.Min(float64(el1), float64(el2))) },
			expected: 0,
		},
		{
			name:     "test3",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   nil,
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.Min(test.script)
		require.Equal(t, test.expected, result)

	}
}

func TestLen(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		expected int
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			expected: 5,
		},
		{
			name:     "test2",
			input:    New([]int{}),
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.Len()
		require.Equal(t, test.expected, result)

	}
}

func TestToSlice(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		expected []int
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "test2",
			input:    New([]int{}),
			expected: []int{},
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		slice := test.input.ToSlice()
		require.Equal(t, test.expected, slice)
	}
}

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    *collection[int]
		script   func(int) string
		expected string
	}{
		{
			name:     "test1",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   func(el int) string { return strconv.Itoa(el) + ", " },
			expected: "1, 2, 3, 4, 5, ",
		},
		{
			name:     "test2",
			input:    New([]int{}),
			script:   func(el int) string { return strconv.Itoa(el) + ", " },
			expected: "",
		},
		{
			name:     "test3",
			input:    New([]int{1, 2, 3, 4, 5}),
			script:   nil,
			expected: "",
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		result := test.input.ToString(test.script)
		require.Equal(t, test.expected, result)

	}
}
