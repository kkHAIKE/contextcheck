[![CircleCI](https://circleci.com/gh/sylvia7788/contextcheck.svg?style=svg)](https://circleci.com/gh/sylvia7788/contextcheck)


# contextcheck

`contextcheck` is a static analysis tool, it is used to check a function whether use context.Background() or context.TODO() directly instead of the input ctx when calling the sub-function, or even don't pass the ctx, which will result in a broken call link.

For example:

```go
func call1(ctx context.Context) {
    ...
    call2(context.Background()) // should use ctx

    call3() // should pass the ctx
    ...
}

func call2(ctx context.Context) {
    ...
}

func call3() {
    ctx := context.TODO()
    call2(ctx)
}
```

## Installation

You can get `contextcheck` by `go get` command.

```bash
$ go get -u github.com/sylvia7788/contextcheck
```

or build yourself.

```bash
$ make build
$ make install
```

## Usage

Invoke `contextcheck` with your package name

```bash
$ contextcheck ./...
$ # or
$ contextcheck github.com/you/yourproject/...
```
