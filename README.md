![image](img/gopher.png)

![Go Version](https://img.shields.io/badge/go%20version-1.16-61CFDD.svg?style=flat-square)

<div id="content-section>

# Content
1. [Install](#install-section)
2. [What is Gofunc ?](#whatis-section)
3. [Example](#example-section)
4. [Functions](#functions-section)
5. [Methods](#methods-section)

<div>

<div id="install-section">

## Install 

```shell 
go get github.com/kdl-dev/gofunc
```
</div> 

<div id="whatis-section">

## What is Gofunc ?
---
<p>
Gofunc is a free library that allows you to write clean, elegant code for working with arrays of data. With Gofunc, you can process a slice in a functional style, which is very convenient and more readable. Gofunc will save you time, which you can use for more serious stuff in your code.
</p>
</div>

---

<div id="example-section">

## Example

<p>
Below you can see user slicing processing without Gofunc and with Gofunc. Algorithm of operations for working with users:

1. for each user, increase the age by 1;
2. leave only those users who are 18 or older;
3. sort users by age;
4. reverse users;
5. print information to the console about each user;
</p>

---

### Input data
```go
type User struct {
	Id   int64
	Name string
	Age  int
}
```

```go
var Users = []User{
	{1, "Kate", 25}, {2, "John", 17},
	{3, "Sam", 22}, {4, "Marina", 15},
	{5, "Nikita", 32}, {6, "Ksenia", 16},
	{7, "Alex", 11}, {8, "Tony", 50},
	{9, "Max", 32}, {10, "Veronika", 4},
}
```

---

### Without Gofunc

```go
func WithoutGofunc() {
	copyUsers := make([]User, len(Users))
	filteredUsers := make([]User, 0, len(Users))

	copy(copyUsers, Users)

	for i := 0; i < len(copyUsers); i++ {
		copyUsers[i].Age++
	}

	for i := 0; i < len(copyUsers); i++ {
		if copyUsers[i].Age >= 18 {
			filteredUsers = append(filteredUsers, copyUsers[i])
		}
	}

	QuickSortByUserAge(filteredUsers)

	for i := 0; i < len(filteredUsers)/2; i++ {
		filteredUsers[i], filteredUsers[len(filteredUsers)-i-1] =
			filteredUsers[len(filteredUsers)-i-1], filteredUsers[i]
	}

	for i := 0; i < len(filteredUsers); i++ {
		fmt.Printf("%+v\n", filteredUsers[i])
	}
}
```

---

### With Gofunc

```go
func WithGofunc() {
	collection := gofunc.New(Users)
	collection.
		Map(func(el User) User { el.Age++; return el }).
		Filter(func(el User) bool { return el.Age >= 18 }).
		Sort(func(arr []User) { QuickSortByUserAge(arr) }).
		Reverse().
		ForEach(func(el User) { fmt.Printf("%+v\n", el) })
}
```
---
<p>As you can see, the code looks shorter and prettier in the second variant.</p>

<p>You can play with the code in goplay space: 
<a href="https://goplay.space/#P1mrWp7ZePV">gofunc example</a>.</p>

<div>

---

<div id="functions-section">


## Functions

* `New[T comparable](arr []T) *collection[T]`
<p>	
	Returns a collection (collection is a wrapper over a slice).
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5}
	collection := gofunc.New(slice)
}
```
<div>

<div id="methods-section">

---

## Methods 

`(c *collection[T])`

<br>

* `ForEach(f func(el T))`
<p>
	Performs an action for each element of this collection.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5}
	collection := gofunc.New(slice)
	collection.
		ForEach(func(el int) { fmt.Printf("%d, ", el) }) // 1, 2, 3, 4, 5,
}
```

<br>

* `Map(f func(el T) T) *collection[T]`
<p>
	Returns a collection consisting of the results of applying the given function to the elements of this collection.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5}
	collection := gofunc.New(slice)
	collection.
		Map(func(el int) int { return el * el }).
		ForEach(func(el int) { fmt.Printf("%d, ", el) }) // 1, 4, 9, 16, 25,
}
```

<br>

* `Reduce(f func(el, accum T) T) T`
<p>
	Performs a reduction on the elements of this stream, using the provided identity value and an associative accumulation function, and returns the reduced value.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5}
	collection := gofunc.New(slice)
	sum := collection.
		Reduce(func(el, accum int) int { return el + accum })

	fmt.Println(sum) // 15
}
```

<br>

* `Filter(f func(el T) bool) *collection[T]`
<p>
	Returns a collection consisting of the elements of this collection that match the given condition.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5}
	collection := gofunc.New(slice)
	collection.
		Filter(func(el int) bool { return el%2 != 0 }).
		ForEach(func(el int) { fmt.Printf("%d, ", el) }) // 1, 3, 5,
}
```

<br>

* `Match(f func(el T) bool) bool`
<p>
	Returns whether any elements of this collection match the provided condition.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5}
	collection := gofunc.New(slice)
	isMatched := collection.
		Match(func(el int) bool { return el < 0 })

	fmt.Println(isMatched) // false
}
```

<br>

* `AllMatch(f func(el T) bool) bool`
<p>
	Returns whether all elements of this collection match the provided condition.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5}
	collection := gofunc.New(slice)
	isAllMatched := collection.
		AllMatch(func(el int) bool { return el > 0 })

	fmt.Println(isAllMatched) // true
}
```

<br>

* `Distinct() *collection[T]`
<p>
	Returns a collection consisting of the distinct elements.
</p>

```go
{
	slice := []int{1, 2, 1, 4, 2, -2, 10, 1}
	collection := gofunc.New(slice)
	collection.
		Distinct().
		ForEach(func(el int) { fmt.Printf("%d, ", el) }) // 1, 2, 4, -2, 10,
}
```

<br>

* `Limit(n int) *collection[T]`
<p>
	Returns a collection consisting of the elements of this collection, truncated to be no longer than n in length.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	collection := gofunc.New(slice)
	collection.
		Limit(6).
		ForEach(func(el int) { fmt.Printf("%d, ", el) }) // 1, 2, 3, 4, 5, 6,
}
```

<br>

* `Skip(n int) *collection[T]`
<p>
	Returns a collection consisting of the remaining elements of this collection after discarding the first n elements of the collection.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	collection := gofunc.New(slice)
	collection.
		Skip(6).
		ForEach(func(el int) { fmt.Printf("%d, ", el) }) // 7, 8, 9, 
}
```

<br>

* `Sort(f func(arr []T)) *collection[T]`
<p>
	Returns a collection consisting of the elements of this collection, sorted according by input function.
</p>

```go
{
	slice := []int{0, -5, -7, 1, 3, 2, 11, 8, 4}
	collection := gofunc.New(slice)
	collection.
		Sort(func(arr []int) { sort.Ints(arr) }).
		ForEach(func(el int) { fmt.Printf("%d, ", el) }) // -7, -5, 0, 1, 2, 3, 4, 8, 11,
}
```

<br>

* `Reverse() *collection[T]`
<p>
	Returns a collection consisting of the elements of this collection, in reverse order.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	collection := gofunc.New(slice)
	collection.
		Reverse().
		ForEach(func(el int) { fmt.Printf("%d, ", el) }) // 9, 8, 7, 6, 5, 4, 3, 2, 1,
}
```

<br>

* `Max(compareFunc func(firstEl, secondEl T) T) T`
<p>
	Returns the maximum element of this collection according to the provided compare function.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	collection := gofunc.New(slice)
	max := collection.
		Max(func(firstEl, secondEl int) int {
			return int(math.Max(float64(firstEl), float64(secondEl)))
		})

	fmt.Println(max) // 9
}
```

<br>

* `Min(compareFunc func(firstEl, secondEl T) T) T`
<p>
	Returns the minimum element of this collection according to the provided compare function.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	collection := gofunc.New(slice)
	min := collection.
		Min(func(firstEl, secondEl int) int {
			return int(math.Min(float64(firstEl), float64(secondEl)))
		})

	fmt.Println(min) // 1
}
```

<br>

* `Len() int`
<p>
	Returns the count of elements in collection.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	collection := gofunc.New(slice)

	fmt.Println(collection.Len()) // 10
}
```

<br>

* `ToSlice() []T`
<p>
	Converts a collection to a slice of elements.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	collection := gofunc.New(slice)

	newSlice := collection.
		Map(func(el int) int { return el * 3 }).
		Filter(func(el int) bool { return el%2 == 0 }).
		ToSlice()

	fmt.Printf("%T: %v", newSlice, newSlice) // []int: [6 12 18 24 30]
}
```

<br>

* `ToString(f func(el T) string) string`
<p>
	Converts a collection to a string.
</p>

```go
{
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	collection := gofunc.New(slice)

	newSlice := collection.
		Map(func(el int) int { return el * 3 }).
		Filter(func(el int) bool { return el%2 == 0 }).
		ToString(func(el int) string { return strconv.Itoa(el) + " " })

	fmt.Printf("%T: %v", newSlice, newSlice) // string: 6 12 18 24 30
}
```

<br>

<div>

---

[Back to content](#content-section)