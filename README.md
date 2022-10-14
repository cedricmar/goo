# GOo

<img src="blob_960_720.webp" width="100">

GOo aims to port javascript's underscore functionnalities in Golang, because reinventing the wheel is no fun :(

Since it uses generics it is really simple to use.

[Custom foo description](#Each\(\))

[Custom foo description](#EachMap())


## Each()

**Each** loops through a slice of \<any> type.

### Example:

```go
sl := []int{1, 2, 3}
Each(sl, func(el int, i int, s []int) {
    s[i] = el * 2
})

-> [2 4 6]
```

The function ```func(el T, i int, s []T) {}``` is common modifier pattern for slice operations.

```T``` can be **any** type.

```el T``` is the slice element being modified.

```i int``` is the index of the element ```el``` in the slice.

```s []T``` is the slice itself.

## EachMap()

**EachMap** loops through a map, order is not guaranteed.

### Example:

```go
IncrementFunc := func(el int, k string, m map[string]int) {
    m[k] += 1
}

EachMap(myMap, IncrementFunc)
```

the function ```func(el T, k U, m map[U]T)``` is the equivalent of the slices modifier function for maps.

## EachMapOrder()

Same as **EachMap** but loops using the keys alphanumeric order.

## Map()

Map each values of a slice through a modifier.

```func Map(s []T, fn func(el T) T)```

## Keys() / KeysOrdered()

Collects the map keys in a slice.

If you need the keys ordered use **KeysOrdered**.

```func Keys(m map[U]T) []U```

```func KeysOrdered(m map[U]T) []U```

## Clone()

Makes a copy of a slice and returns it.

```func Clone(s []T) []T```


## Unique()

Removes identical elements from a slice.

```func Unique(s []T) []T```

## FoundIn()

Returns a boolean if the element k was found in the slice s

```func FoundIn(k T, s []T) bool```

## FoundAt()

Returns the index of the found element in the slice or -1.

It also returns a boolean to indicate if the element k was found.

```func FoundAt(k T, s []T) (int, bool)```

## Index()

Turns a slice into a map using a function on each slice element to infer its key.

```func Index(s []T, fn func(el T) K) map[K]T```

### Example:

```go
Index(myDBEntities, fun(el DBEntity) int {
    return el.ID
})

-> 1: {ID: 1, Name: "One"}
   2: {ID: 2, Name: "Two"}
   3: ...
```

## LookupTable()

Slice elements become the keys in the returned map.

```func LookupTable(s []T) map[T]struct{}```

### Example:

```go
sl := []int{1, 2, 3}
LookupTable(sl)

-> 1: struct{}
   2: struct{}
   3: ...
```
