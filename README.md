# Collections
 
This module is for useful collection algorithms implemented in go.

## Installing


```bash
go get github.com/Hilson-Alex/collections
```

## Usage

Once you installed the module, import the package you want to use.

### Streams
the package `streams` implements functional methods to go slices, allowing
a JavaScript-like syntax with method chaining. Ex:

```go
package pkg

import (
	"fmt"
	"strconv"

	"github.com/Hilson-Alex/collections/streams"
)

func main() {
	var arr = []string{"1", "2", "5", "3", "7", "12"}
	var example = streams.Transform[string, int](arr).
		Map(func(item string, _ int) int {
			var result, _ = strconv.Atoi(item)
			return result
		}).
		Filter(func(item int, _ int) bool {
			return item >= 5
		})
	fmt.Println(example)
	// example = []int{5, 7, 12}   
}
```

### streams.Stream

The `Stream[T]` type is a wrapper for `T[]` slices that includes immutable methods,
which means that methods like `Stream.Filter` will produce a new `Stream[T]` with
the filtered entries.

A stream method cannot produce a stream of a type different from `T`. Every `Stream[T]` 
method can only produce a `Stream[T]`, `T` or a primitive type (like `bool`). The only
exception is `Wrap`, that receives a callback to change the stream into an`interface{}`, 
which allow you to change the stream in any way you like. 

If you need to change `T` into another type, look [Transform](#streamstransform).

#### Callback methods:

| Method  |            Callback             | Description                                                                                                                                                     |
|---------|:-------------------------------:|-----------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Foreach |   (item T, index int) => void   | Apply the function to every item                                                                                                                                |
|         |                                 |                                                                                                                                                                 |
| Some    |   (item T, index int) => bool   | Return true if any match                                                                                                                                        |
| None    |   (item T, index int) => bool   | Return true if none match                                                                                                                                       |
| Every   |   (item T, index int) => bool   | Return true if all match                                                                                                                                        |
| Reduce  | (acc T, item T, index int) => T | Apply the callback on each element from index 1, passing the return<br/>of the last iteration (using the element at 0 on the first execution)                   |
|         |                                 |                                                                                                                                                                 |
| Map     |    (item T, index int) => T     | Changes each item in another item of type `T`                                                                                                                   |
| Filter  |   (item T, index int) => bool   | Return a stream with only the elements that passes the provided<br/>test                                                                                        |
| Sort    |       (i T, j T) => bool        | Sorts the array with the supplied function. If true, i comes before j                                                                                           |
|         |                                 |                                                                                                                                                                 |
| Wrap    |      ([]T) => interface{}       | This is a special function. Wrap receives the entire stream an can<br/>change it in any way, returning an interface{} to allow the change of<br/>the slice type |

<a id="wrap-func" name="wrap-func"></a>
> The `Wrap` function was implemented to keep the fluent interface while adapting the 
> entire stream. An example of the intended use of the `Wrap` function is shown bellow,
> where a `Stream` is changed into a `Transform`
> ```go
> var arr = streams.Stream[int]{1, 2, 5, 3, 7, 12}.
>       Filter(func(item int, _ int) bool {
>           return item >= 5
>       }).
>       Wrap(streams.AsTransform[int, string]).(streams.Transform[int,string]).
>       Map(func(i int, _ int) string {
>           return strconv.Itoa(i)
>       })
>  // In this example the Wrap was used to perform a type cast, then just a type assert is needed to chain
>  // with the Transform methods.
> ```
 
#### Modifying (non-mutating) methods:

Those methods make a shallow copy of the stream and then modify the copy adding or 
switching elements, returning the new modified stream:

| Method  | Description                                                                       |
|---------|-----------------------------------------------------------------------------------|
| Concat  | Creates a new stream with all passed items on the end of it                       |
| Prepend | Creates a new stream with all passed items on the beginning                       |
| AddAt   | Creates a new stream an then adds all the items on the given index                |
| With    | Creates a new stream and then changes the element of the given on the given index |


### streams.Transform

The `Transform[T, U]` type is an enhanced `Stream[T]` that provides methods to change a
`Stream[T]` into `U` or a `Stream[U]`. Like a stream, a transform will not change the 
initial stream, but create a new modified stream.

The transform's methods return a `Stream` rather than a `Transform`, because we can't
track future transformations. If you want to transform a `Stream[A]` into a `Stream[B]`
and then on a `Stream[C]`, you'll need a `Transform[A, B]` that will return a `Stream[B]`
and then use a [Wrap](#wrap-func) to create a `Transform[B, C]`, like the example bellow:

```go
var streamC = streams.Transform[A, B](arr).
        Map(callbackToB).
        // Call any stream methods needed, like filter or sort
        Wrap(streams.AsTransform[B, C]).(streams.Transform[B, C]). 
        Map(callbackToC)

// Or you want to chain transformations right away without calling any methods 
// between transformations, you can just do all transformations in a single callback

streamC = streams.Transform[A, C](arr).
        Map(func(item A, index int) C {
            // this avoids all the wrapping, and assertion, you just can't
            // do stuff that would handle a Stream[B]
            var mappedToB = callbackToB(item, index)
            return callbackToC(mappedToB, index)
        })
```

#### Methods:

| Method  |             Callback             | Description                                                                                                                                                                                                                               |
|---------|:--------------------------------:|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| Map     |     (item T, index int) => U     | Changes each item in an item of type `U`                                                                                                                                                                                                  |
| FlatMap |    (item T, index int) => []U    | Changes each item in a slice of `U`, then join all slices in a<br/>`Stream[U]`                                                                                                                                                            |
| Reduce  | (acc U, item T, index int) =>  U | Unlike the stream, the transform reduce receives an initial value<br/>of type `U`, and apply the callback to each element from index 0,<br/>passing the return of the last iteration (using the initial value on<br/>the first iteration) |


## Contribute

## License

