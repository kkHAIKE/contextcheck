[![CircleCI](https://circleci.com/gh/sylvia7788/contextcheck.svg?style=svg)](https://circleci.com/gh/sylvia7788/contextcheck)


# contextcheck

`contextcheck` is a static analysis tool which checks whether use context.Background() and context.TODO() directly.

## Install

You can get `contextcheck` by `go get` command.

```bash
$ go get -u github.com/jingyugao/rowserrcheck
```

or build yourself.

```bash
$ make build
$ make install
```

## Usage
Invoke `contextcheck` with your package name
```
contextcheck ./...
# or
contextcheck github.com/you/yourproject/...
```
