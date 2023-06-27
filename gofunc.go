package gofunc

type collection[T comparable] struct {
	data []T
}

/*
Returns a collection (collection is a wrapper over a slice).
*/
func New[T comparable](arr []T) *collection[T] {
	var newCollection collection[T]
	newCollection.data = make([]T, len(arr))
	copy(newCollection.data, arr)

	return &newCollection
}

/*
Generates a collection based on the received function.
The number of elements is given by the second input argument.
*/
func Generate[T comparable](script func() T, limit int) *collection[T] {
	if script == nil {
		return nil
	}

	if limit <= 0 {
		return New(make([]T, 0))
	}

	newCollection := New(make([]T, limit))

	for i := 0; i < limit; i++ {
		newCollection.data[i] = script()
	}

	return newCollection
}

// Performs an action for each element of this collection.
func (c *collection[T]) ForEach(consume func(el T)) {
	if consume == nil {
		return
	}

	for _, value := range c.data {
		consume(value)
	}
}

/*
Returns a collection consisting of the results of applying
the given function to the elements of this collection.
*/
func (c *collection[T]) Map(predicate func(el T) T) *collection[T] {
	if predicate == nil {
		return New(c.data)
	}

	newcollection := New[T](c.data)

	for i := 0; i < len(newcollection.data); i++ {
		newcollection.data[i] = predicate(newcollection.data[i])
	}

	return newcollection
}

/*
Returns a collection consisting of the results of replacing
each element of this collection with the contents of a mapped
collection produced by applying the provided mapping function to
each element.
*/
func (c *collection[T]) FlatMap(predicate func(el T) (T, T)) *collection[T] {
	if predicate == nil {
		return New(c.data)
	}

	newcollection := New[T](make([]T, len(c.data)*2))

	for i := 0; i < len(newcollection.data)-1; i += 2 {
		newcollection.data[i], newcollection.data[i+1] = predicate(c.data[i/2])
	}

	return newcollection
}

/*
Performs a reduction on the elements of this stream,
using the provided identity value and an associative
accumulation function, and returns the reduced value.
*/
func (c *collection[T]) Reduce(binaryOperator func(el, accum T) T) T {
	var accum T

	if c.Len() == 0 || binaryOperator == nil {
		return accum
	}

	for _, value := range c.data {
		accum = binaryOperator(value, accum)
	}

	return accum
}

/*
Returns a collection consisting of the elements
of this collection that match the given condition.
*/
func (c *collection[T]) Filter(filter func(el T) bool) *collection[T] {
	if filter == nil {
		return New(c.data)
	}

	newcollection := New[T](make([]T, 0, len(c.data)))

	for _, value := range c.data {
		if filter(value) {
			newcollection.data = append(newcollection.data, value)
		}
	}

	return newcollection
}

/*
Returns whether any elements of this collection
match the provided condition.
*/
func (c *collection[T]) Match(predicate func(el T) bool) bool {
	if c.Len() == 0 || predicate == nil {
		return false
	}

	for _, value := range c.data {
		if predicate(value) {
			return true
		}
	}

	return false
}

/*
Returns whether all elements of this collection
match the provided condition.
*/
func (c *collection[T]) AllMatch(predicate func(el T) bool) bool {
	if c.Len() == 0 || predicate == nil {
		return false
	}

	for _, value := range c.data {
		if !predicate(value) {
			return false
		}
	}

	return true
}

/*
Returns a collection consisting of the distinct elements.
*/
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

/*
Returns a collection consisting of the elements of this collection,
truncated to be no longer than n in length.
*/
func (c *collection[T]) Limit(n int) *collection[T] {
	if n > c.Len() {
		n = c.Len()
	} else if n < 0 {
		n = 0
	}

	return New[T](c.data[:n])
}

/*
Returns a collection consisting of the remaining elements
of this collection after discarding the first n elements
of the collection.
*/
func (c *collection[T]) Skip(n int) *collection[T] {
	if n > c.Len() {
		n = c.Len()
	} else if n < 0 {
		n = 0
	}

	return New[T](c.data[n:])
}

/*
Returns a collection consisting of the elements
of this collection, sorted according by input function.
*/
func (c *collection[T]) Sort(sort func(arr []T)) *collection[T] {
	if sort == nil {
		return New(c.data)
	}

	newcollection := New[T](c.data)

	sort(newcollection.data)

	return newcollection
}

/*
Returns a collection consisting of the elements
of this collection, in reverse order.
*/
func (c *collection[T]) Reverse() *collection[T] {
	newcollection := New[T](make([]T, len(c.data)))

	for i := 0; i < len(newcollection.data); i++ {
		newcollection.data[i] = c.data[len(c.data)-i-1]
	}

	return newcollection
}

/*
Replaces all the first matching elements of the collection
passed to targets with the element passed to replacement.
*/
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

/*
Replaces all elements of the collection passed to the target
objects with the element passed to replace.
*/
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

/*
Returns the maximum element of this collection
according to the provided compare function.
*/
func (c *collection[T]) Max(compare func(firstEl, secondEl T) T) T {
	var resultMax, currentMax T

	if compare == nil {
		return resultMax
	}

	if len(c.data) > 1 {
		resultMax = compare(c.data[0], c.data[1])
	}

	for i := 1; i < len(c.data)-1; i++ {
		currentMax = compare(c.data[i], c.data[i+1])
		resultMax = compare(resultMax, currentMax)
	}

	return resultMax
}

/*
Returns the minimum element of this collection
according to the provided compare function.
*/
func (c *collection[T]) Min(compare func(firstEl, secondEl T) T) T {
	var resultMin, currentMin T

	if compare == nil {
		return resultMin
	}

	if len(c.data) > 1 {
		resultMin = compare(c.data[0], c.data[1])
	}

	for i := 1; i < len(c.data)-1; i++ {
		currentMin = compare(c.data[i], c.data[i+1])
		resultMin = compare(resultMin, currentMin)
	}

	return resultMin
}

/*
Returns the count of elements in collection.
*/
func (c *collection[T]) Len() int {
	return len(c.data)
}

/*
Converts a collection to a slice of elements.
*/
func (c *collection[T]) ToSlice() []T {
	return c.data
}

/*
Converts a collection to a string.
*/
func (c *collection[T]) ToString(convert func(el T) string) string {
	var resultStr string

	if convert == nil {
		return resultStr
	}

	for _, value := range c.data {
		resultStr += convert(value)
	}

	return resultStr
}
