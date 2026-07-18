# ugo

[![CI](https://github.com/mainvec/ugo/actions/workflows/ci.yml/badge.svg)](https://github.com/mainvec/ugo/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/mainvec/ugo.svg)](https://pkg.go.dev/github.com/mainvec/ugo)

`ugo` is a collection of small, standard-library-only Go utilities used across
Mainvec projects. Each package can be imported independently, so applications
only need to adopt the pieces they use.

## Requirements

- Go 1.22 or newer
- No third-party runtime dependencies

## Installation

Install the package you need through the module:

```sh
go get github.com/mainvec/ugo@latest
```

Then import it directly. For example:

```go
package main

import (
	"fmt"

	"github.com/mainvec/ugo/registry"
)

func main() {
	services := registry.NewRegistry[string]()
	services.Register("search", "https://search.example.com")

	endpoint, ok := services.Lookup("search")
	fmt.Println(endpoint, ok)
}
```

## Packages

| Package | Purpose |
| --- | --- |
| [`assert`](./assert) | Minimal panic-based assertions for runtime invariants. |
| [`cli`](./cli) | A zero-dependency command and flag framework. |
| [`collections`](./collections) | Generic helpers for slices and maps. |
| [`oencoding`](./oencoding) | A registry for pluggable object encodings. |
| [`oencoding/json`](./oencoding/json) | A JSON encoding driver registered through a blank import. |
| [`omap`](./omap) | Deterministic map iteration by key or value. |
| [`registry`](./registry) | A concurrent generic registry of named values. |
| [`validate`](./validate) | Composable validation rules and error aggregation. |

API details and examples are available on
[pkg.go.dev](https://pkg.go.dev/github.com/mainvec/ugo).

## Stability

The project currently uses `v0` releases. APIs may change between minor
versions until `v1.0.0`; release tags follow semantic versioning.

## Contributing

See [CONTRIBUTING.md](./CONTRIBUTING.md) for the development workflow and
required checks. Please report security issues according to
[SECURITY.md](./SECURITY.md).

## License

This project is available under the [MIT License](./LICENSE). The `cli` package
contains code adapted from `tractordev/toolkit-go`; see
[THIRD_PARTY_NOTICES.md](./THIRD_PARTY_NOTICES.md).
