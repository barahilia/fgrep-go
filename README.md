# fgrep-go
Aho-Corasick algorithm implementation in Go.

The time has come to learn Golang. For things to be interesting and fruitful, I
chose to implement a light `fgrep`. The algorithm details were taken from the
wonderful Wikipedia page
[Aho–Corasick algorithm](https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm).

## Usage

From the command line:
```
cd $GOPATH/src/fgrep-go
go test
go install
$GOPATH/bin/fgrep-go -f WORDS FILE
```

From code:
```go
package main

import (
    "fmt"
    "github.com/barahilia/fgrep-go"
)

func main() {
    text := "abc"
    matches := fgrep.Search(text, "a", "b", "c")

    for match := range matches {
        fmt.Print("%s appears in text[%d: %d]\n",
            match.word, match.start, match.end)
    }
}
```
