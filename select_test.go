package gaqlbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		builder SelectBuilder
		query   string
	}{{
		builder: NewSelectBuilder().Select("campaign.id", "campaign.name"),
		query:   "SELECT campaign.id, campaign.name",
	}}

	for _, test := range tests {
		assert.Equal(t, test.query, test.builder.Build())
	}
}