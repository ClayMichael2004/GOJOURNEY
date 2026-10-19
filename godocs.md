### GO TOOLS
- godoc
- playground
- refactor
- txtar
- analysis
- intsets
- cover
- copyright

### go datatypes
- basic
- composite
- reference
- interface

# basic
- numbers i.e int, float and complex
- strings
- boolean

# reference types
 - slice
 - maps
 - channels
 - pointers
 - function types

go is **statically types** meaning you have to define what type it at compile time; we must define the type before running.

# maps
need low memory compared to slices and arrays.
store data in key value pairs

array: [X]Y
slice: []Y
maps: map[Z]Y

**where Y is the data type**
and in maps **Z** is the arbitrary 

**chanels** are used as a cintainer of element values for the execution of the goroutines.
**pointers** rep the memory address of the element value.

If you create a slice using make, we have the capacity(cap) and the length(len)
**make([]int, 5, 100) means it is  a slice of length:5 and a capacity of 100.**

**len** - number of elements currently in the slice
**cap** - total number of elements the slice can hold

**Untyped constants** written without ecplicitly giving it a type. goves it precision and allows you to adapt to any type you want later without stricyt type errors.

# switch case usage

the syntax is: 
switch variable{ case values: ... case values: ... default: }