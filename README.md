# sesstype [![Build Status](https://travis-ci.org/nickng/sesstype.svg?branch=master)](https://travis-ci.org/nickng/sesstype) [![GoDoc](https://godoc.org/go.nickng.io/sesstype?status.svg)](http://godoc.org/go.nickng.io/sesstype)


`sesstype` is a parser and library for the `sesstype` type language.

The `sesstype` type language for Multiparty Session Types is defined in
[sesstype](https://github.com/nickng/sesstype.rs#parser).

To get:

    go get -u go.nickng.io/sesstype/...

By default the `parse-sesstype` command will also be installed.

## Syntax

The basic syntax of `sesstype` language is given below, for details see
[godoc](https://godoc.org/go.nickng.io/sesstype)

#### Global Types

    G   ::= P->Q: { l(U).G , ... } Interaction between P and Q with message l(U)
          | *T.G                   Recursion with label T, body G
          | T                      Type variable
          | end                    End type

### Local Types

    L   ::= Q &{ ?l(U).L, ... }    Branching, receive l(U) from role Q
          | P +{ !l(U).L, ... }    Selection, send l(U) to role P
          | *T.L                   Recursion with label T, body L
          | T                      Type variable
          | end                    End type

## License

`sesstype` is licensed under the [Apache License](http://www.apache.org/licenses/LICENSE-2.0).
