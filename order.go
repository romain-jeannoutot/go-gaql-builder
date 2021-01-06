package gaqlbuilder

const (
	// OrderAsc order result ascending
	OrderAsc order = "ASC"
	// OrderDesc order result descending
	OrderDesc order = "DESC"
)

type order string

func (o order) String() string {
	return string(o)
}
