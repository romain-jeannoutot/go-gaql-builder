package gaqlbuilder

type resource string

func (r resource) String() string {
	return string(r)
}
