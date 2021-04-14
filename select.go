package gaqlbuilder

import (
	"bytes"
	"strconv"
	"strings"
)

// SelectBuilder is builder to build SELECT queries
type SelectBuilder struct {
	fields   []string
	resource string
	where    []string
	orderBy  []string
	limit    uint
}

// NewSelectBuilder create a new SELECT builder with default values
func NewSelectBuilder() *SelectBuilder {
	return &SelectBuilder{}
}

// Select clause specifies a set of fields to fetch in the request
func (b *SelectBuilder) Select(fields ...string) *SelectBuilder {
	b.fields = fields
	return b
}

func (b *SelectBuilder) From(resource string) *SelectBuilder {
	b.resource = resource
	return b
}

func (b *SelectBuilder) Where(exprs ...string) *SelectBuilder {
	b.where = append(b.where, exprs...)
	return b
}

func (b *SelectBuilder) OrderBy(cols ...string) *SelectBuilder {
	b.orderBy = cols
	return b
}

func (b *SelectBuilder) Limit(limit uint) *SelectBuilder {
	b.limit = limit
	return b
}

func (b *SelectBuilder) Build() string {
	buffer := &bytes.Buffer{}

	buffer.WriteString("SELECT ")
	buffer.WriteString(strings.Join(b.fields, ", "))

	buffer.WriteString(" FROM ")
	buffer.WriteString(b.resource)

	if len(b.where) > 0 {
		buffer.WriteString(" WHERE ")
		buffer.WriteString(strings.Join(b.where, " AND "))
	}

	if len(b.orderBy) > 0 {
		buffer.WriteString(" ORDER BY ")
		buffer.WriteString(strings.Join(b.orderBy, ", "))
	}

	if b.limit > 0 {
		buffer.WriteString(" LIMIT ")
		buffer.WriteString(strconv.Itoa(int(b.limit)))
	}

	return buffer.String()
}
