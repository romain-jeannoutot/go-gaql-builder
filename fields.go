package gaqlbuilder

import (
	"fmt"
	"strings"
)

type fields []string

func (fs fields) String() string {
	return fmt.Sprintf("SELECT %s", strings.Join(fs, ", "))
}
