package collector

import (
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestDevicesCollector(t *testing.T) {
	client := &fakeClient{}
	c := NewDevicesCollector(client)
	problems, err := testutil.CollectAndLint(c,
		`orbi_connected_device_downrate`,
		`orbi_connected_device_uprate`,
		`orbi_connected_device_info`,
	)

	assert.NoError(t, err)
	assert.Empty(t, problems)

	assert.NoError(t, testutil.CollectAndCompare(c, strings.NewReader(`
# HELP orbi_connected_device_downrate Unknown
# TYPE orbi_connected_device_downrate gauge
orbi_connected_device_downrate{connected_orbi="Orbi Router",connection_type="5G",device_type="Laptop",host="http://fakehost.local",ip="10.0.0.10",model="Generic",name="My Laptop"} 0
# HELP orbi_connected_device_info is the device active
# TYPE orbi_connected_device_info gauge
orbi_connected_device_info{connected_orbi="Orbi Router",connection_type="5G",device_type="Laptop",host="http://fakehost.local",ip="10.0.0.10",model="Generic",name="My Laptop"} 1
# HELP orbi_connected_device_uprate Unknown
# TYPE orbi_connected_device_uprate gauge
orbi_connected_device_uprate{connected_orbi="Orbi Router",connection_type="5G",device_type="Laptop",host="http://fakehost.local",ip="10.0.0.10",model="Generic",name="My Laptop"} 0
`)))
}
