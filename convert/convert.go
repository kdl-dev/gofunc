package convert

import (
	"fmt"
)

type ints interface {
	int8 | int16 | int32 | int64 | int
}

type uints interface {
	uint8 | uint16 | uint32 | uint64 | uint
}

type floats interface {
	float32 | float64
}

type complex interface {
	complex64 | complex128
}

func New[T, V comparable](slice []T, convertFunc func(el T) V) []V {
	if slice == nil || convertFunc == nil {
		return nil
	}

	newSlice := make([]V, len(slice))

	for i := 0; i < len(slice); i++ {
		newSlice[i] = convertFunc(slice[i])
	}

	return newSlice
}

func IntToString[T ints | uints](el T) string {
	return fmt.Sprintf("%d", el)
}

func FloatToString[T floats](el T) string {
	return fmt.Sprintf("%f", el)
}

func BoolToString(el bool) string {
	return fmt.Sprintf("%t", el)
}

func RuneToString(el rune) string {
	return fmt.Sprintf("%q", el)
}

func ComplexToString[T complex](el T) string {
	return fmt.Sprintf("%g", el)
}

/*
func StringToInt[T ints](el string) T {
	newEl, err := strconv.ParseInt(el, 10, 64)
	if err != nil {
		panic(err)
	}

	return T(newEl)
}

func StringToUInt[T uints](el string) T {
	newEl, err := strconv.ParseUint(el, 10, 64)
	if err != nil {
		panic(err)
	}

	return T(newEl)
}

func StringToFloat[T floats](el string) T {
	newEl, err := strconv.ParseFloat(el, 64)
	if err != nil {
		panic(err)
	}

	return T(newEl)
}

func StringToBool(el string) bool {
	newEl, err := strconv.ParseBool(el)
	if err != nil {
		panic(err)
	}

	return newEl
}

func StringToSliceRune[T complex](el string) []rune {
	return []rune(el)
}

func StringToComplex[T complex](el string) T {
	newEl, err := strconv.ParseComplex(el, 128)
	if err != nil {
		panic(err)
	}

	return T(newEl)
}*/
