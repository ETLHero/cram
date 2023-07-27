# Cram
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/ETLHero/cram)

Type conversions. By force if necessary.

Cram will do its best to convert any type given to it. You're unlikely ever to need this library. This is only useful when you're writing dynamic code and must go from an unknown to a known type.

### Example
```go
package main

import (
  "fmt"
  "github.com/ETLHero/cram"
)

func main() {
  var numbers []int
  cram.Into(&numbers, "1,2,3,4,5")
  fmt.Println(numbers) // => [1 2 3 4 5]
}
```

## Features
- Extensible! See below.
- Map (hashing) lookup for conversations. No slow switch cases.
- Gracefully handles pointers
- Uses reflect package rather than casting. For all the above features.

## Extending
Don't like a conversion? Missing a conversion? Add it yourself. The `Conversions` variable holds all the conversions just add, replace or remove what you want and call the `Into` function like normal.

## Conversions
Some conversions might not be obvious, like `[]float` to `bool`. I think you are best off looking at the tests to see how they convert.
- Convert from [Bool](singleboolto_test.go)
- Convert from [Float](singlefloatto_test.go)
- Convert from [Int](singleintto_test.go)
- Convert from [String](singlestringto_test.go)
- Convert from [Bool Slice](multiboolto_test.go)
- Convert from [Float Slice](multifloatto_test.go)
- Convert from [Int Slice](multiintto_test.go)
- Convert from [String Slice](multistringto_test.go)
