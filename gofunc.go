package gofunc

type collection[T comparable] struct {
	data []T
}

func New[T comparable](arr []T) *collection[T] {
	var newCollection collection[T]
	newCollection.data = make([]T, len(arr))
	copy(newCollection.data, arr)

	return &newCollection
}

func Generate[T comparable](script func() T, limit int) *collection[T] {
	if limit <= 0 {
		return New(make([]T, 0))
	}

	newCollection := New(make([]T, limit))

	for i := 0; i < limit; i++ {
		newCollection.data[i] = script()
	}

	return newCollection
}

func (c *collection[T]) Map(predicate func(el T) T) *collection[T] {
	newcollection := New[T](c.data)

	for i := 0; i < len(newcollection.data); i++ {
		newcollection.data[i] = predicate(newcollection.data[i])
	}

	return newcollection
}

func (c *collection[T]) FlatMap(predicate func(el T) (T, T)) *collection[T] {
	newcollection := New[T](make([]T, len(c.data)*2))

	for i := 0; i < len(newcollection.data)-1; i += 2 {
		newcollection.data[i], newcollection.data[i+1] = predicate(c.data[i/2])
	}

	return newcollection
}

func (c *collection[T]) Reduce(binaryOperator func(el, accum T) T) T {
	var accum T

	for _, value := range c.data {
		accum = binaryOperator(value, accum)
	}

	return accum
}

func (c *collection[T]) Filter(filter func(el T) bool) *collection[T] {
	newcollection := New[T](make([]T, 0, len(c.data)))

	for _, value := range c.data {
		if filter(value) {
			newcollection.data = append(newcollection.data, value)
		}
	}

	return newcollection
}

func (c *collection[T]) Match(predicate func(el T) bool) bool {
	for _, value := range c.data {
		if predicate(value) {
			return true
		}
	}

	return false
}

func (c *collection[T]) AllMatch(predicate func(el T) bool) bool {
	for _, value := range c.data {
		if !predicate(value) {
			return false
		}
	}

	return true
}

func (c *collection[T]) Distinct() *collection[T] {
	newcollection := New[T](make([]T, 0, len(c.data)))
	unique := make(map[T]bool)

	for _, value := range c.data {
		if _, isExists := unique[value]; !isExists {
			unique[value] = true
			newcollection.data = append(newcollection.data, value)
		}
	}

	return newcollection
}

func (c *collection[T]) Limit(n int) *collection[T] {
	if n > c.Len() {
		n = c.Len()
	} else if n < 0 {
		n = 0
	}

	return New[T](c.data[:n])
}

func (c *collection[T]) Skip(n int) *collection[T] {
	if n > c.Len() {
		n = c.Len()
	} else if n < 0 {
		n = 0
	}

	return New[T](c.data[n:])
}

func (c *collection[T]) ForEach(consume func(el T)) {
	for _, value := range c.data {
		consume(value)
	}
}

func (c *collection[T]) Sort(sort func(arr []T)) *collection[T] {
	newcollection := New[T](c.data)

	sort(newcollection.data)

	return newcollection
}

func (c *collection[T]) Reverse() *collection[T] {
	newcollection := New[T](make([]T, len(c.data)))

	for i := 0; i < len(newcollection.data); i++ {
		newcollection.data[i] = c.data[len(c.data)-i-1]
	}

	return newcollection
}

func (c *collection[T]) Replace(targets []T, replacement T) *collection[T] {
	newcollection := New[T](c.data)

Exit:
	for i := 0; i < len(newcollection.data); i++ {
		for j := 0; j < len(targets); j++ {
			if newcollection.data[i] == targets[j] {
				newcollection.data[i] = replacement
				targets = append(targets[:j], targets[j+1:]...)

				if len(targets) == 0 {
					continue Exit
				}

				break
			}
		}
	}

	return newcollection
}

func (c *collection[T]) ReplaceAll(targets []T, replacement T) *collection[T] {
	newcollection := New[T](c.data)

	for i := 0; i < len(newcollection.data); i++ {
		for j := 0; j < len(targets); j++ {
			if newcollection.data[i] == targets[j] {
				newcollection.data[i] = replacement

				break
			}
		}
	}

	return newcollection
}

func (c *collection[T]) Max(compare func(firstEl, secondEl T) T) T {
	var resultMax, currentMax T

	if len(c.data) > 1 {
		resultMax = compare(c.data[0], c.data[1])
	}

	for i := 1; i < len(c.data)-1; i++ {
		currentMax = compare(c.data[i], c.data[i+1])
		resultMax = compare(resultMax, currentMax)
	}

	return resultMax
}

func (c *collection[T]) Min(compare func(firstEl, secondEl T) T) T {
	var resultMin, currentMin T

	if len(c.data) > 1 {
		resultMin = compare(c.data[0], c.data[1])
	}

	for i := 1; i < len(c.data)-1; i++ {
		currentMin = compare(c.data[i], c.data[i+1])
		resultMin = compare(resultMin, currentMin)
	}

	return resultMin
}

func (c *collection[T]) Len() int {
	return len(c.data)
}

func (c *collection[T]) ToSlice() []T {
	return c.data
}

func (c *collection[T]) ToString(convert func(el T) string) string {
	var resultStr string
	for _, value := range c.data {
		resultStr += convert(value)
	}

	return resultStr
}
