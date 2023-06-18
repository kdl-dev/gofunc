package convert_test

import (
	"strconv"
	"testing"

	"github.com/kdl-dev/gofunc/convert"
	"github.com/stretchr/testify/require"
)

type Test[T, V comparable] struct {
	input    []T
	expected []V
}

var (
	intSlice     = []int{1, 2, 3, 4}
	stringSlice  = []string{"1", "2", "3", "4"}
	boolSlice    = []bool{true, false, true}
	stringSlice2 = []string{"true", "false", "true"}
)

func TestNew(t *testing.T) {
	test1 := Test[int, string]{
		input:    intSlice,
		expected: stringSlice,
	}

	test2 := Test[string, bool]{
		input:    stringSlice2,
		expected: boolSlice,
	}

	actualStringSlice := convert.New(test1.input, func(el int) string { return strconv.Itoa(el) })
	require.Equal(t, test1.expected, actualStringSlice)

	actualBoolSlice := convert.New(test2.input, func(el string) bool {
		newEl, _ := strconv.ParseBool(el)
		return newEl
	})
	require.Equal(t, test2.expected, actualBoolSlice)
}

func TestIntToString(t *testing.T) {
	test := Test[int, string]{
		input:    []int{5},
		expected: []string{"5"},
	}
	actual := convert.IntToString(test.input[0])

	require.Equal(t, test.expected[0], actual)
}

func TestFloatToString(t *testing.T) {
	test := Test[float64, string]{
		input:    []float64{5.0},
		expected: []string{"5.000000"},
	}
	actual := convert.FloatToString(test.input[0])

	require.Equal(t, test.expected[0], actual)
}

func TestBoolToString(t *testing.T) {
	test := Test[bool, string]{
		input:    []bool{false},
		expected: []string{"false"},
	}
	actual := convert.BoolToString(test.input[0])

	require.Equal(t, test.expected[0], actual)
}

func TestRuneToString(t *testing.T) {
	test := Test[rune, string]{
		input:    []rune{65},
		expected: []string{"'A'"},
	}
	actual := convert.RuneToString(test.input[0])

	require.Equal(t, test.expected[0], actual)
}

func TestComplexToString(t *testing.T) {
	test := Test[complex128, string]{
		input:    []complex128{1 + 5i},
		expected: []string{"(1+5i)"},
	}
	actual := convert.ComplexToString(test.input[0])

	require.Equal(t, test.expected[0], actual)
}
