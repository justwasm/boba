# boba - browser-oriented bubbletea adapter

<p>
    <a href="https://justwasm.github.io/boba/"><img src="https://img.shields.io/badge/Command%20Ref-6B2DAD" alt="Command Reference"></a>
    <a href="https://github.com/justwasm/boba/tags"><img src="https://img.shields.io/github/tag/justwasm/boba.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/justwasm/boba?tab=doc"><img src="https://pkg.go.dev/badge/github.com/justwasm/boba?utm_source=godoc" alt="GoDoc"></a>
    <a href="https://github.com/justwasm/boba/blob/main/CODE_OF_CONDUCT.md"><img src="https://img.shields.io/badge/Contributor%20Covenant-2.1-4baaaa.svg"  alt="Code Of Conduct"></a>
</p>

`boba` is a Golang module that facilitates embedding [BubbleTea](https://github.com/charmbracelet/bubbletea) Terminal User Interfaces (TUIs) into a Web Browser.

<img src="./etc/boba.png" alt="booba mascot" width="512">

## Installation

**Go** (server-side library and CLI tools):

```sh
go get github.com/justwasm/boba
```

## How and What?

The primary enabling technologies of this are:

 * [`BubbleTea`](https://github.com/charmbracelet/bubbletea) - Terminal UI framework for Go
 * [`WebAssembly`](https://webassembly.org) - For running Go code in browsers

## Embedding a BubbleTea Application in a Web Browser

We can take entire BubbleTea applications and embed them into a Web Browser. The primary limitation is that all of its dependencies can also be compiled to WebAssembly.

### Quickstart

The top-level `boba.Run` picks the right runtime for the build target, so a single `main.go` works for both the native terminal and the browser:

```go
package main

import (
    "log"

    boba "github.com/justwasm/boba"
)

func main() {
    if err := boba.Run(initialModel()); err != nil {
        log.Fatal(err)
    }
}
```

For easier porting from Bubble Tea — or when you need the program handle for `Send`, `Quit`, etc. — use `boba.NewProgram`:

```go
func main() {
    bp := boba.NewProgram(initialModel())
    if _, err := bp.Run(); err != nil {
        log.Fatal(err)
    }
}
```

Build and run natively with `go run ./cmd/myapp`. Build for the browser with `go run github.com/justwasm/boba/cmd/boba-wasm-build -o web/app.wasm ./cmd/myapp/`.

For finer control, the [`wasm`](./wasm) subpackage exposes the browser bridge directly, and native code can construct a `tea.Program` the usual way.

## Credits

- https://github.com/charmbracelet/bubbletea
- https://github.com/tmc/bubbweb
  - https://github.com/myka0/bubbweb-v2
- https://github.com/NimbleMarkets/go-booba
- https://github.com/BigJk/bubbletea-in-wasm
