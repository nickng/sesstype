package global

import "go.nickng.io/sesstype"

type parseSendRecvs struct {
	m sesstype.Message
	g Type
}

type parseContext struct {
	Type     Type
	TypeVars map[string]bool
	stack    *parseContext
}

func newParseContext() *parseContext {
	return &parseContext{TypeVars: make(map[string]bool)}
}

func (ctx *parseContext) push() *parseContext {
	pc := newParseContext()
	pc.stack = ctx
	return pc
}

func (ctx *parseContext) pop() *parseContext {
	return ctx.stack
}
