# Unquote

Simple program to unquote a string from stdin or arguments. Designed for JSON use.

# Usage Example

## From pipe/stdin:

```sh
echo "{\"foo\":\"bar\"}" | unquote
```

## From Args

```sh
unquote "{\"foo\":\"bar\"}"
```
