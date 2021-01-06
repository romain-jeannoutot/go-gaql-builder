package gaqlbuilder

// SelectBuilder is builder to build SELECT queries
type SelectBuilder interface {
	Select(fields ...string) SelectBuilder
}

// NewSelectBuilder create a new SELECT builder with default values
func NewSelectBuilder() SelectBuilder {
	return &selectBuilder{}
}

type selectBuilder struct {
	fields []string
}

// Select clause specifies a set of fields to fetch in the request
func (b *selectBuilder) Select(fields ...string) SelectBuilder {
	b.fields = fields
	return b
}
