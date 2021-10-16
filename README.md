# WASM-FETCH
[![GoDoc](https://godoc.org/mlctrez/wasm-fetch?status.svg)](https://godoc.org/mlctrez/wasm-fetch)

A go-wasm library that wraps the [Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API)

### Install
`go get github.com/mlctrez/wasm-fetch`

### Motivation
Importing net/http adds ~4 MBs to your wasm binary. If that's an issue for you, you can use this
library to make fetch calls.

### Fork
Forked from [https://github.com/marwan-at-work/wasm-fetch](https://github.com/marwan-at-work/wasm-fetch) to add allow
use in [https://github.com/maxence-charriere/go-app](github.com/maxence-charriere/go-app). 


### Example

```golang
package main

import (
    "context"
    "time"

    "github.com/mlctrez/wasm-fetch"
)

ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
resp, err := fetch.Fetch("/some/api/call", &fetch.Opts{
    Body:   strings.NewReader(`{"one": "two"}`),
    Method: fetch.MethodPost,
    Signal: ctx,
})
// use response...
```


### Status
GO-WASM is currently experimental and therefore this package is experimental as well, things can break unexpectedly. 