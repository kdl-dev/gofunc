package gofunc

type collection[T comparable] struct {
	data []T
}

func New[T comparable](arr []T) *collection[T] {
	var newArr collection[T]
	newArr.data = make([]T, len(arr))
	copy(newArr.data, arr)

	return &newArr
}

func (c *collection[T]) Map(f func(el T) T) *collection[T] {
	newcollection := New[T](c.data)

	for i := 0; i < len(newcollection.data); i++ {
		newcollection.data[i] = f(newcollection.data[i])
	}

	return newcollection
}

func (c *collection[T]) Reduce(f func(el, accum T) T) T {
	var accum T

	for i := 0; i < len(c.data); i++ {
		accum = f(c.data[i], accum)
	}

	return accum
}

func (c *collection[T]) Filter(f func(el T) bool) *collection[T] {
	newcollection := New[T](make([]T, 0, len(c.data)))

	for i := 0; i < len(c.data); i++ {
		if f(c.data[i]) {
			newcollection.data = append(newcollection.data, c.data[i])
		}
	}

	return newcollection
}

func (c *collection[T]) Match(f func(el T) bool) bool {
	for i := 0; i < len(c.data); i++ {
		if f(c.data[i]) {
			return true
		}
	}

	return false
}

func (c *collection[T]) AllMatch(f func(el T) bool) bool {
	for i := 0; i < len(c.data); i++ {
		if !f(c.data[i]) {
			return false
		}
	}

	return true
}

func (c *collection[T]) Limit(limit int) *collection[T] {
	return New[T](c.data[:limit])
}

func (c *collection[T]) Skip(skip int) *collection[T] {
	return New[T](c.data[skip:])
}

func (c *collection[T]) ForEach(f func(el T)) {
	for i := 0; i < len(c.data); i++ {
		f(c.data[i])
	}
}

func (c *collection[T]) Sort(f func(arr []T)) *collection[T] {
	newcollection := New[T](c.data)

	f(newcollection.data)

	return newcollection
}

func (c *collection[T]) Reverse() *collection[T] {
	newcollection := New[T](make([]T, len(c.data)))

	for i := 0; i < len(c.data); i++ {
		newcollection.data[i] = c.data[len(c.data)-i-1]
	}

	return newcollection
}

func (c *collection[T]) ToString(f func(el T) string) string {
	var resultStr string
	for i := 0; i < len(c.data); i++ {
		resultStr += f(c.data[i])
	}

	return resultStr
}

func (c *collection[T]) Max(compareFunc func(firstEl, secondEl T) T) T {
	var resultMax, currentMax T

	if len(c.data) > 1 {
		resultMax = compareFunc(c.data[0], c.data[1])
	}

	for i := 1; i < len(c.data)-1; i++ {
		currentMax = compareFunc(c.data[i], c.data[i+1])
		resultMax = compareFunc(resultMax, currentMax)
	}

	return resultMax
}

func (c *collection[T]) Min(compareFunc func(firstEl, secondEl T) T) T {
	var resultMin, currentMin T

	if len(c.data) > 1 {
		resultMin = compareFunc(c.data[0], c.data[1])
	}

	for i := 1; i < len(c.data)-1; i++ {
		currentMin = compareFunc(c.data[i], c.data[i+1])
		resultMin = compareFunc(resultMin, currentMin)
	}

	return resultMin
}

func (c *collection[T]) Len() int {
	return len(c.data)
}

func (c *collection[T]) ToSlice() []T {
	return c.data
}
