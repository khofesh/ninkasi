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

## evaluator

```shell
go test ./evaluator
```

## object

```shell
go test ./object
```

## parser

```shell
go test ./parser
```

## running the REPL

```shell
[fahmad@ryzen ninkasi]$  go run main.go
Hello fahmad! This is the Ninkasi programming language!
.help won't help you. We don't have enough information
>> print("hola mundo!")
hola mundo!
null
>> let workers = [{"name": "yana", "age": 30}, {"name": "david", "age": 25}];
>> workers[0]
{age: 30, name: yana}
>> workers[0]["name"]
yana
>> workers[0]["age"] + workers[1]["age"];
55
>> let getAge = funk(worker) {worker["age"];};
>> getAge(workers[0]);
30
>>
```

## MACROS

```shell
[fahmad@ryzen ninkasi]$  go run main.go
Hello fahmad! This is the Ninkasi programming language!
.help won't help you. We don't have enough information
>> let unless = macro(condition, consequence, alternative) {quote(if (!(unquote(condition))) { unquote(consequence);} else { unquote(alternative); }); };
>> unless(10 > 5, print("not greater"), print("greater"));
greater
null
```
