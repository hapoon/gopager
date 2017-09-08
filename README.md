gopager
===
`gopager` is a library for pagination.

## Installation

Make sure that Go is installed on your computer. Type the following command in your terminal:

`go get github.com/hapoon/gopager`

After it the package is ready to use.

Add following line in your `*.go` file:

```go
import "github.com/hapoon/gopager"
```

## Usage

You need to implement the Pageable interface for slices to use paging.
`Len()` returns the length of slices.

```go
type Item struct {
    ID uint
}

type Items []Item

func (i Items) Len() int {
    return len(i)
}
```

```go
items := Items{
    Item{ID: 1},
    Item{ID: 2},
    Item{ID: 3},
    Item{ID: 4},
    Item{ID: 5},
}
pageSize := 2

p := gopager.NewPaginater(items,pageSize)

i := Items{}
for p.HasNext() {
    p.Next(&i)
    fmt.Printf("Page: %v i: %v\n",p.CurrentPage(),i)
}
```

Output:
```
Page: 1 i: [{ID:1} {ID:2}]
Page: 2 i: [{ID:3} {ID:4}]
Page: 3 i: [{ID:5}]
```

## License

[MIT License](LICENSE)
