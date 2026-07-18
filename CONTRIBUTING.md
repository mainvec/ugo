# Contributing

Thanks for contributing to `ugo`.

## Before you start

Search the existing issues before opening a new one. For significant API or
behavior changes, open an issue first so the design and compatibility impact
can be discussed before implementation.

## Development

You need Go 1.22 or newer. Clone your fork, create a focused branch, and make
the smallest change that solves the issue. This module intentionally uses only
the Go standard library; discuss any proposed third-party dependency in an
issue before adding it.

Run the required checks before submitting a pull request:

```sh
test -z "$(gofmt -l .)"
go vet ./...
go test -race ./...
```

Add or update tests for behavior changes. Keep exported APIs documented and
avoid unrelated refactoring in the same pull request.

## Pull requests

Describe what changed, why it changed, and how it was tested. Link the relevant
issue with a closing keyword when appropriate. Pull requests must pass CI and
be licensed under the repository's MIT License.

By submitting a contribution, you agree that your contribution is licensed
under the terms in [LICENSE](./LICENSE).