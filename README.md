# tiny-http-echo

A minimal Go HTTP server that locally echoes the requests it receives.

This program can be useful in debugging HTTP clients and/or server components of distributed systems.

---

## Example

Using 5 panes in tmux with 1 client (top left), 1 test server(bottom left), 3 tiny-http-echo servers.

![](https://i.imgur.com/6u5HCiq.png)

![](https://i.imgur.com/nYjCa77.png)

![](https://i.imgur.com/KGHBOPT.png)

## Installation

> Must have Go and Git installed

**1. Retrieve source code:**

* `go get -u github.com/stvhwrd/tiny-http-echo`

**2. Build binary:**

* `go install github.com/stvhwrd/tiny-http-echo`

## Configuration

There is only one configuration option, taken as a runtime flag:
```
  -port string
        The port for this HTTP server to listen on, eg. -port=8080
```

## Usage

Use `go run` and include the [port flag](#configuration):

* `go run $GOPATH/github.com/stvhwrd/tiny-http-echo -port=YOUR_PORT_GOES_HERE`

or, if you have `$GOPATH/bin` in your `$PATH`, you can use the shorthand:

* `tiny-http-echo -port=YOUR_PORT_GOES_HERE`
