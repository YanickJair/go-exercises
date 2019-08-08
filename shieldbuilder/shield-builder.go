package shieldbuilder

import "strings"

// Shield - object definition
type Shield struct {
	front bool
	back  bool
	right bool
	left  bool
}

// builder - ...
type shBuilder struct {
	code string
}

// NewShieldBuilder - Autogenerate a new ShieldBuilder
func NewShieldBuilder() *shBuilder {
	return new(shBuilder)
}

// RaiseFront - up front shield
func (sb *shBuilder) RaiseFront() *shBuilder {
	sb.code += "F"
	return sb
}

// RaiseBack - up back shield
func (sb *shBuilder) RaiseBack() *shBuilder {
	sb.code += "B"
	return sb
}

// RaiseRight - up right shield
func (sb *shBuilder) RaiseRight() *shBuilder {
	sb.code += "R"
	return sb
}

// RaiseLeft - up left shield
func (sb *shBuilder) RaiseLeft() *shBuilder {
	sb.code += "L"
	return sb
}

// Build - return a shield type
func (sb *shBuilder) Build() *Shield {
	code := sb.code
	return &Shield{
		front: strings.Contains(code, "F"),
		back:  strings.Contains(code, "B"),
		right: strings.Contains(code, "R"),
		left:  strings.Contains(code, "L"),
	}
}
