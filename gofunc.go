package gofunc

type Interface interface {
	comparable
}

type Collection[T Interface] struct {
	data []T
}

func New[T Interface](arr []T) *Collection[T] {
	var newArr Collection[T]
	newArr.data = make([]T, len(arr))
	copy(newArr.data, arr)

	return &newArr
}

func (c *Collection[T]) Map(f func(el T) T) *Collection[T] {
	newCollection := New[T](c.data)

	for i := 0; i < len(newCollection.data); i++ {
		newCollection.data[i] = f(newCollection.data[i])
	}

	return newCollection
}

func (c *Collection[T]) Reduce(f func(el, accum T) T) T {
	var accum T

	for i := 0; i < len(c.data); i++ {
		accum = f(c.data[i], accum)
	}

	return accum
}

func (c *Collection[T]) Filter(f func(el T) bool) *Collection[T] {
	newCollection := New[T](make([]T, 0, len(c.data)))

	for i := 0; i < len(c.data); i++ {
		if f(c.data[i]) {
			newCollection.data = append(newCollection.data, c.data[i])
		}
	}

	return newCollection
}

func (c *Collection[T]) Match(f func(el T) bool) bool {
	for i := 0; i < len(c.data); i++ {
		if f(c.data[i]) {
			return true
		}
	}

	return false
}

func (c *Collection[T]) AllMatch(f func(el T) bool) bool {
	for i := 0; i < len(c.data); i++ {
		if !f(c.data[i]) {
			return false
		}
	}

	return true
}

func (c *Collection[T]) Limit(limit int) *Collection[T] {
	return New[T](c.data[:limit])
}

func (c *Collection[T]) Skip(skip int) *Collection[T] {
	return New[T](c.data[skip:])
}

func (c *Collection[T]) ForEach(f func(el T)) {
	for i := 0; i < len(c.data); i++ {
		f(c.data[i])
	}
}

func (c *Collection[T]) Sort(f func(arr []T)) *Collection[T] {
	newCollection := New[T](c.data)

	f(newCollection.data)

	return newCollection
}

func (c *Collection[T]) Reverse() *Collection[T] {
	newCollection := New[T](make([]T, len(c.data)))

	for i := 0; i < len(c.data); i++ {
		newCollection.data[i] = c.data[len(c.data)-i-1]
	}

	return newCollection
}

func (c *Collection[T]) ToString(f func(el T) string) string {
	var resultStr string
	for i := 0; i < len(c.data); i++ {
		resultStr += f(c.data[i])
	}

	return resultStr
}

func (c *Collection[T]) Max(compareFunc func(firstEl, secondEl T) T) T {
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

func (c *Collection[T]) Min(compareFunc func(firstEl, secondEl T) T) T {
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

func (c *Collection[T]) Len() int {
	return len(c.data)
}

func (c *Collection[T]) ToSlice() []T {
	return c.data
}
