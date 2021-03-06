[![CircleCI](https://circleci.com/gh/spatialcurrent/goreplacer/tree/master.svg?style=svg)](https://circleci.com/gh/spatialcurrent/goreplacer/tree/master) [![Go Report Card](https://goreportcard.com/badge/spatialcurrent/goreplacer)](https://goreportcard.com/report/spatialcurrent/goreplacer)  [![GoDoc](https://godoc.org/github.com/spatialcurrent/goreplacer?status.svg)](https://godoc.org/github.com/spatialcurrent/goreplacer) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/spatialcurrent/goreplacer/blob/master/LICENSE)

# goreplacer

## Description

**goreplacer** is a simple tool for efficiently replacing bytes within a stream of bytes.

## Platforms

The following platforms are supported.  Pull requests to support other platforms are welcome!

| GOOS | GOARCH |
| ---- | ------ |
| darwin | amd64 |
| linux | amd64 |
| windows | amd64 |
| linux | arm64 |

## Releases

Find releases at [https://github.com/spatialcurrent/goreplacer/releases](https://github.com/spatialcurrent/goreplacer/releases).  You might want to rename your binary to just `goreplacer`.  See the **Building** section below to build from scratch.

**Darwin**

- `goreplacer_darwin_amd64` - CLI for Darwin on amd64 (includes `macOS` and `iOS` platforms)

**Linux**

- `goreplacer_linux_amd64` - CLI for Linux on amd64
- `goreplacer_linux_amd64` - CLI for Linux on arm64

**Windows**

- `goreplacer_windows_amd64.exe` - CLI for Windows on amd64

## Usage

See the usage below or the following examples.

```shell
goreplacer is a simple tool for for efficiently replacing bytes within a stream of bytes.
Replacements are encoded as a 2 dimensional JSON array of a series of replacement pairs ([OLD, NEW]).
If the positional argument is "-" or not given, then reads from stdin.

Usage:
  goreplacer [--lines] [--max MAX] [--replacements] '[[OLD1, NEW1], [OLD2, NEW2], ...]' [-|FILE]

Flags:
  -h, --help                  help for goreplacer
  -l, --lines                 process as lines
  -n, --max int               maximum number of replacements (default -1)
  -r, --replacements string   replacements encoded as 2-dimensional JSON array, i.e., [[OLD1, NEW1], [OLD2, NEW2], ...]
```

# Examples

**Replace words**

```shell
cat FILE | goreplacer -r '[["world", "planet"]]'
```

**Update Copyright**

```shell
cat FILE | goreplacer -n 1 -r '[["Copyright (C) 2019", "Copyright (C) 2020"]]'
```

**Unescape Double Quotes**

```shell
$ echo '{""hello"":""world""}' | goreplacer -r '[["\"\"","\""]]'
{"hello":"world"}
```

**Replace Single Quotes with Double Quotes**

```shell
$ echo "{'hello':'world'}" | goreplacer -r "[[\"'\",\"\\\"\"]]"
{"hello":"world"}
```

## Building

Use `make help` to see help information for each target.

**CLI**

The `make build_cli` script is used to build executables for Linux and Windows.  Use `make install` for standard installation as a go executable.

**Changing Destination**

The default destination for build artifacts is `bin`, but you can change the destination with an environment variable.  For building on a Chromebook consider saving the artifacts in `/usr/local/go/bin`, e.g., `DEST=/usr/local/go/bin make build_cli`

## Testing

**CLI**

To run CLI testes use `make test_cli`, which uses [shUnit2](https://github.com/kward/shunit2).  If you recive a `shunit2:FATAL Please declare TMPDIR with path on partition with exec permission.` error, you can modify the `TMPDIR` environment variable in line or with `export TMPDIR=<YOUR TEMP DIRECTORY HERE>`. For example:

```
TMPDIR="/usr/local/tmp" make test_cli
```

**Go**

To run Go tests use `make test_go` (or `bash scripts/test.sh`), which runs unit tests, `go vet`, `go vet with shadow`, [errcheck](https://github.com/kisielk/errcheck), [ineffassign](https://github.com/gordonklaus/ineffassign), [staticcheck](https://staticcheck.io/), and [misspell](https://github.com/client9/misspell).

## Contributing

[Spatial Current, Inc.](https://spatialcurrent.io) is currently accepting pull requests for this repository.  We'd love to have your contributions!  Please see [Contributing.md](https://github.com/spatialcurrent/goreplacer/blob/master/CONTRIBUTING.md) for how to get started.

## License

This work is distributed under the **MIT License**.  See **LICENSE** file.
