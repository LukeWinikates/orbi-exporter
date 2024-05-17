package collector

import (
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStatisticsCollector(t *testing.T) {
	client := &fakeClient{}
	c := NewStatisticsCollector(client)
	problems, err := testutil.CollectAndLint(c,
		`orbi_received_bytes_per_second`,
		`orbi_transmitted_bytes_per_second`,
		`orbi_received_packets`,
		`orbi_transmitted_packets`,
		`orbi_collisions`,
		`orbi_uptime_seconds`,
		`orbi_system_uptime_seconds`,
	)

	assert.NoError(t, err)
	assert.Empty(t, problems)

	assert.NoError(t, testutil.CollectAndCompare(c, strings.NewReader(expectedResponse)))
}
