# hyperlink

hyperlink ([OSC8](https://gist.github.com/egmontkob/eb114294efbcd5adb1944c9f3cb5feda)) utility script with tmux support.

## Usage

```
$ hyperlink <link> [<text>]

$ hyperlink http://example.com "This is a link"

# If you omit <text>, <link> is also used as <text>.
$ hyperlink http://example.com
```

## Go

### Library

[![GoDoc](https://godoc.org/github.com/haya14busa/hyperlink?status.svg)](https://godoc.org/github.com/haya14busa/hyperlink)

```go
import "github.com/haya14busa/hyperlink"
```

### Install

```
go get -u github.com/haya14busa/hyperlink/cmd/hyperlink
```
