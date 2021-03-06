package gaqlbuilder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	tests := []struct {
		builder *SelectBuilder
		query   string
	}{{
		builder: NewSelectBuilder().
			Select("campaign.id", "campaign.name").
			From("ad_group").
			Where("metrics.impressions > 0", "segments.device = MOBILE").
			OrderBy("campaign.name ASC", "metrics.impressions DESC").
			Limit(10),
		query: "SELECT campaign.id, campaign.name FROM ad_group WHERE metrics.impressions > 0 AND segments.device = MOBILE ORDER BY campaign.name ASC, metrics.impressions DESC LIMIT 10",
	}, {
		builder: NewSelectBuilder().
			Select("campaign.id", "campaign.name").
			From("campaign"),
		query: "SELECT campaign.id, campaign.name FROM campaign",
	}}

	for _, test := range tests {
		assert.Equal(t, test.query, test.builder.Build())
	}
}
