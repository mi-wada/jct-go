# jct-go

A Go library for calculating Japanese Consumption Tax. Inspired by [jct](https://github.com/moneyforward/jct).

## Install

```shell
go get github.com/mi-wada/jct-go
```

## Requirements

Go 1.11 or later.

## Usage

```go
package main

import (
	"fmt"
	"time"

	"github.com/mi-wada/jct-go"
)

func main() {
	amount := int64(100)
	at := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)

	tax := jct.Tax(amount, at)
	fmt.Println(tax)
	// Output: 10

	total := jct.Total(amount, at)
	fmt.Println(total)
	// Output: 110
}
```
