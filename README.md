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

### Nested JSON Objects

jqs will treat any nested objects as strings, and urlencode them. E.g.

```
$ echo '{"hello": { "foo": { "bar": "baz" }  }}' | jqs
hello=%7B%22foo%22%3A%7B%22bar%22%3A%22baz%22%7D%7D
```

URL decoded, the above is equivalent to:
```
hello={"foo":{"bar":"baz"}}
```

## Purpose

I often find myself needing to do this converstion when trying to replay a network request with `cURL`, given the `Copy JSON` data from the IntelliJ debugger.

## Build

```
$ go build jqs.go
```