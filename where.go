package gaqlbuilder

import (
	"fmt"
	"strings"
)

type where []string

func (w where) String() string {
	if len(w) > 0 {
		return fmt.Sprintf("WHERE %s", strings.Join(w, " AND "))
	}

	return ""
}
