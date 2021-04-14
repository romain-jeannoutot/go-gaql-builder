package gaqlbuilder

import (
	"fmt"
	"strings"
)

// SelectBuilder is builder to build SELECT queries
type SelectBuilder struct {
	fields   fields
	resource resource
	where    where
	orderBy  orderBy
	limit    limit
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

func (b *SelectBuilder) From(res string) *SelectBuilder {
	b.resource = resource(res)
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

func (b *SelectBuilder) Limit(nb uint) *SelectBuilder {
	b.limit = limit(nb)
	return b
}

func (b *SelectBuilder) build(strs ...fmt.Stringer) string {
	var ss []string
	for _, str := range strs {
		if s := str.String(); s != "" {
			ss = append(ss, s)
		}
	}

	return strings.Join(ss, " ")
}

func (b *SelectBuilder) Build() string {
	return b.build(b.fields, b.resource, b.where, b.orderBy, b.limit)
}
