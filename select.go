package gaqlbuilder

import (
	"bytes"
	"strings"
)

// SelectBuilder is builder to build SELECT queries
type SelectBuilder interface {
	Builder

	Select(fields ...string) SelectBuilder
	From(resource string) SelectBuilder
	Where(exprs ...string) SelectBuilder
}

// NewSelectBuilder create a new SELECT builder with default values
func NewSelectBuilder() SelectBuilder {
	return &selectBuilder{}
}

type selectBuilder struct {
	fields   []string
	resource string
	where    []string
}

// Select clause specifies a set of fields to fetch in the request
func (b *selectBuilder) Select(fields ...string) SelectBuilder {
	b.fields = fields
	return b
}

func (b *selectBuilder) From(resource string) SelectBuilder {
	b.resource = resource
	return b
}

func (b *selectBuilder) Where(exprs ...string) SelectBuilder {
	b.where = append(b.where, exprs...)
	return b
}

func (b *selectBuilder) Build() string {
	buffer := &bytes.Buffer{}

	buffer.WriteString("SELECT ")
	buffer.WriteString(strings.Join(b.fields, ", "))

	buffer.WriteString(" FROM ")
	buffer.WriteString(b.resource)

	if len(b.where) > 0 {
		buffer.WriteString(" WHERE ")
		buffer.WriteString(strings.Join(b.where, " AND "))
	}

	return buffer.String()
}
