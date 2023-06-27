package convert_test

import (
	"strconv"
	"testing"

	"github.com/kdl-dev/gofunc/convert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		script   func(el int) string
		expected []string
	}{
		{
			name:     "test1",
			slice:    []int{1, 2, 3, 4, 5},
			script:   func(el int) string { return strconv.Itoa(el) },
			expected: []string{"1", "2", "3", "4", "5"},
		},
		{
			name:     "test2",
			slice:    []int{},
			script:   func(el int) string { return strconv.Itoa(el) },
			expected: []string{},
		},
		{
			name:     "test3",
			slice:    nil,
			script:   func(el int) string { return strconv.Itoa(el) },
			expected: nil,
		},
		{
			name:     "test4",
			slice:    []int{1, 2, 3, 4, 5},
			script:   nil,
			expected: nil,
		},
	}

	for _, test := range tests {
		t.Log(test.name)
		slice := convert.New(test.slice, test.script)
		require.Equal(t, test.expected, slice)
	}

}

type Test[T, V comparable] struct {
	input    T
	expected V
}

func TestIntToString(t *testing.T) {
	test := Test[int, string]{
		input:    5,
		expected: "5",
	}
	actual := convert.IntToString(test.input)

	require.Equal(t, test.expected, actual)
}

func TestFloatToString(t *testing.T) {
	test := Test[float64, string]{
		input:    5.0,
		expected: "5.000000",
	}
	actual := convert.FloatToString(test.input)

	require.Equal(t, test.expected, actual)
}

func TestBoolToString(t *testing.T) {
	test := Test[bool, string]{
		input:    false,
		expected: "false",
	}
	actual := convert.BoolToString(test.input)

	require.Equal(t, test.expected, actual)
}

func TestRuneToString(t *testing.T) {
	test := Test[rune, string]{
		input:    65,
		expected: "'A'",
	}
	actual := convert.RuneToString(test.input)

	require.Equal(t, test.expected, actual)
}

func TestComplexToString(t *testing.T) {
	test := Test[complex128, string]{
		input:    1 + 5i,
		expected: "(1+5i)",
	}
	actual := convert.ComplexToString(test.input)

	require.Equal(t, test.expected, actual)
}
