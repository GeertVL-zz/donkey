package object

import (
	"fmt"
	"github.com/geertvl/donkey/ast"
	"bytes"
	"strings"
)

type ObjectType string

const (
	IntegerObj = "INTEGER"
	BooleanObj = "BOOLEAN"
	NullObj	   = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj   = "ERROR"
	FunctionObj = "FUNCTION"
)

type Object interface {
	Type() 		ObjectType
	Inspect()	string
}

type Integer struct {
	Value	int64
}
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return IntegerObj }

type Boolean struct {
	Value 	bool
}
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return BooleanObj }

type Null struct {}
func (n *Null) Inspect() string { return "null" }
func (n *Null) Type() ObjectType { return NullObj }

type ReturnValue struct {
	Value Object
}
func (rv *ReturnValue) Type() ObjectType { return ReturnValueObj }
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

type Error struct {
	Message string
}
func (e *Error) Type() ObjectType { return ErrorObj }
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}
func (f *Function) Type() ObjectType { return FunctionObj }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString("} {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}