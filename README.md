# ninkasi

## lexer_test

```shell
$ go test ./lexer/
ok      github.com/khofesh/ninkasi/lexer        0.002s
```

## parser

```shell
$ go test ./parser/
ok  	github.com/khofesh/ninkasi/parser	0.002s
```

trace and untrace

```shell
go test -v -run TestOperatorPrecedenceParsing ./parser
```

## ast

```shell
$ go test ./ast/
ok  	github.com/khofesh/ninkasi/ast	0.002s
```
