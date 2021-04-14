package gaqlbuilder

import (
	"fmt"
	"strings"
)

type orderBy []string

func (ob orderBy) String() string {
	if len(ob) > 0 {
		return fmt.Sprintf("ORDER BY %s", strings.Join(ob, ", "))
	}

	return ""
}
