# jqs

`jqs` is a small command line tool to convert JSON objects to url encoded query strings.

## Usage

`jqs` can read from piped input, or the first argument it is passed. E.g.

```
$ echo '{"hello": "world"}' | jqs
hello=world
```

or

```
$ jqs '{"hello": "world"}'
hello=world
```

## Purpose

I often find myself needing to do this converstion when trying to replay a network request with `cURL`, given the `Copy JSON` data from the IntelliJ debugger.

## Build

```
$ go build jqs.go
```