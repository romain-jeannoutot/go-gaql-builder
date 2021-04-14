package gaqlbuilder

import "fmt"

type resource string

func (r resource) String() string {
	return fmt.Sprintf("FROM %s", string(r))
}
