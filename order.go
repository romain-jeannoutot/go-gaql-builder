package gaqlbuilder

const (
	// OrderAsc order result ascending
	OrderAsc Order = "ASC"
	// OrderDesc order result descending
	OrderDesc Order = "DESC"
)

type Order string

func (o Order) String() string {
	return string(o)
}
