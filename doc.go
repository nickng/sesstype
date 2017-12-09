// Copyright 2017 Nicholas Ng <nickng@nickng.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package sesstype provides a parser and library for the sesstype mini language.
//
// Syntax
//
// The basic syntax of the sesstype language is as follows, G and L denote
// Global Types and Local Types respectively:
//
//    P,Q ::= alphanum               Role names
//    l   ::= alphanum               Message label
//    U   ::= alphanum               Payload type
//    T   ::= alphanum               Type variable label
//
//    G   ::= P->Q: { l(U).G , ... } Interaction between P and Q with message l(U)
//          | *T.G                   Recursion with label T, body G
//          | T                      Type variable
//          | end                    End type
//
//    L   ::= Q &{ ?l(U).L, ... }    Branching, receive l(U) from role Q
//          | P +{ !l(U).L, ... }    Selection, send l(U) to role P
//          | *T.L                   Recursion with label T, body L
//          | T                      Type variable
//          | end                    End type
//
// As a syntactic sugar, branching and receiving with only a single branch, i.e.
//
//     Q &{ ?l(U).L }
//     P +{ !l(U).L }
//
// can be written without the braces to denote receiving and sending:
//
//     Q ?l(U).L
//     P !l(U).L
//
package sesstype // import "go.nickng.io/sesstype"
