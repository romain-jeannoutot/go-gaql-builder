package gaqlbuilder

import "fmt"

type limit uint

func (l limit) String() string {
	if l > 0 {
		return fmt.Sprintf("LIMIT %d", l)
	}

	return ""
}
