# Unquote

Simple program to unquote a string from stdin or arguments. Designed for JSON use.

# Installation

```sh
go get github.com/TigorLazuardi/unquote
```

# Usage Example

## From pipe/stdin:

```sh
echo "{\"foo\":\"bar\"}" | unquote
```

## From Args

```sh
unquote "{\"foo\":\"bar\"}"
```
