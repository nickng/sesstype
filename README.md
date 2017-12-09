# sesstype [![Build Status](https://travis-ci.org/nickng/sesstype.svg?branch=master)](https://travis-ci.org/nickng/sesstype) [![GoDoc](https://godoc.org/go.nickng.io/sesstype?status.svg)](http://godoc.org/go.nickng.io/sesstype)


`sesstype` is a parser and library for the `sesstype` type language.

The `sesstype` type language for Multiparty Session Types is defined in
[sesstype](https://github.com/nickng/sesstype.rs#parser).

To get:

    go get -u go.nickng.io/sesstype/...

By default the `parse-sesstype` command will also be installed.

## Syntax

The basic syntax of `sesstype` implemented in this parser is given below

```
ident   = [A-Za-z0-9]+
role    = ident
message = ident payload
payload = "()"
        | "(" ident ")"
```

#### Global Types

```
global   = role "->" role ":" interact
         | recur
         | typevar
         | end
interact = sendrecv | "{" sendrecv ("," sendrecv)+ "}"
sendrecv = message "." global
recur    = "*" ident "." global
typevar  = ident
end      = "end"
```

#### Local Types

```
local    = role "&" branch
         | role "+" select
         | lrecur
         | ltypevar
         | end
branch   = recv | "{" recv ("," recv)+ "}"
recv     = "?" message "." local
select   = send | "{" send ("," send)+ "}"
send     = "!" message "." local
lrecur   = "*" ident "." local
ltypevar = ident
lend     = "end"
```
## License

`sesstype` is licensed under the [Apache License](http://www.apache.org/licenses/LICENSE-2.0).
