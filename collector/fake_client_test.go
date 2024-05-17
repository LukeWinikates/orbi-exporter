package collector

import "github.com/LukeWinikates/orbi-exporter/orbi"

var cannedResponse = map[string]*orbi.Metric{
	"sys_uptime": {
		Name:  "sys_uptime",
		Value: "3069182",
	},
	"lan_status": {
		Name:  "lan_status",
		Value: "Link up",
	},
	"lan_txpkts": {
		Name:  "lan_txpkts",
		Value: "88670935",
	},
	"lan_rxpkts": {
		Name:  "lan_rxpkts",
		Value: "43826450",
	},
	"lan_collisions": {
		Name:  "lan_collisions",
		Value: "0",
	},
	"lan_txbs": {
		Name:  "lan_txbs",
		Value: "1345",
	},
	"lan_rxbs": {
		Name:  "lan_rxbs",
		Value: "729",
	},
	"lan_systime": {
		Name:  "lan_systime",
		Value: "3069182",
	},
	"wan_status": {
		Name:  "wan_status",
		Value: "1000M/Full",
	},
	"wan_txpkts": {
		Name:  "wan_txpkts",
		Value: "121869567",
	},
	"wan_rxpkts": {
		Name:  "wan_rxpkts",
		Value: "262693070",
	},
	"wan_collisions": {
		Name:  "wan_collisions",
		Value: "0",
	},
	"wan_txbs": {
		Name:  "wan_txbs",
		Value: "1399",
	},
	"wan_rxbs": {
		Name:  "wan_rxbs",
		Value: "1399",
	},
	"wan_systime": {
		Name:  "wan_systime",
		Value: "2695",
	},
	"bgn_status": {
		Name:  "bgn_status",
		Value: "400M",
	},
	"bgn_txpkts": {
		Name:  "bgn_txpkts",
		Value: "31060602",
	},
	"bgn_rxpkts": {
		Name:  "bgn_rxpkts",
		Value: "60138",
	},
	"bgn_collisions": {
		Name:  "bgn_collisions",
		Value: "0",
	},
	"bgn_txbs": {
		Name:  "bgn_txbs",
		Value: "32",
	},
	"bgn_rxbs": {
		Name:  "bgn_rxbs",
		Value: "6",
	},
	"bgn_systime": {
		Name:  "bgn_systime",
		Value: "3069110",
	},
	"an_status": {
		Name:  "an_status",
		Value: "866M",
	},
	"an_txpkts": {
		Name:  "an_txpkts",
		Value: "17510305",
	},
	"an_rxpkts": {
		Name:  "an_rxpkts",
		Value: "0",
	},
	"an_collisions": {
		Name:  "an_collisions",
		Value: "0",
	},
	"an_txbs": {
		Name:  "an_txbs",
		Value: "0",
	},
	"an_rxbs": {
		Name:  "an_rxbs",
		Value: "0",
	},
	"an_systime": {
		Name:  "an_systime",
		Value: "3069093",
	},
	"bh_status": {
		Name:  "bh_status",
		Value: "866M",
	},
	"bh_txpkts": {
		Name:  "bh_txpkts",
		Value: "32365100",
	},
	"bh_rxpkts": {
		Name:  "bh_rxpkts",
		Value: "4094",
	},
	"bh_collisions": {
		Name:  "bh_collisions",
		Value: "0",
	},
	"bh_txbs": {
		Name:  "bh_txbs",
		Value: "0",
	},
	"bh_rxbs": {
		Name:  "bh_rxbs",
		Value: "0",
	},
	"bh_systime": {
		Name:  "bh_systime",
		Value: "3069093",
	},
	"wwan0_status": {
		Name:  "wwan0_status",
		Value: "no service",
	},
	"wwan0_txpkts": {
		Name:  "wwan0_txpkts",
		Value: "0",
	},
	"wwan0_rxpkts": {
		Name:  "wwan0_rxpkts",
		Value: "0",
	},
	"wwan0_collisions": {
		Name:  "wwan0_collisions",
		Value: "0",
	},
	"wwan0_txbs": {
		Name:  "wwan0_txbs",
		Value: "0",
	},
	"wwan0_rxbs": {
		Name:  "wwan0_rxbs",
		Value: "0",
	},
	"wwan0_systime": {
		Name:  "wwan0_systime",
		Value: "0",
	},
	"lan_status0": {
		Name:  "lan_status0",
		Value: "1000M/Full",
	},
	"lan0_systime": {
		Name:  "lan0_systime",
		Value: "3069133",
	},
	"lan_status1": {
		Name:  "lan_status1",
		Value: "Link down",
	},
	"lan1_systime": {
		Name:  "lan1_systime",
		Value: "0",
	},
	"lan_status2": {
		Name:  "lan_status2",
		Value: "Link down",
	},
	"lan2_systime": {
		Name:  "lan2_systime",
		Value: "0",
	},
	"lan_status3": {
		Name:  "lan_status3",
		Value: "Link down",
	},
	"lan3_systime": {
		Name:  "lan3_systime",
		Value: "0",
	},
}

var expectedResponse = `
# HELP orbi_collisions_total Collisions
# TYPE orbi_collisions_total counter
orbi_collisions_total{component="LAN",host="http://fakehost.local",original_metric_name="lan_collisions"} 0
orbi_collisions_total{component="WAN",host="http://fakehost.local",original_metric_name="wan_collisions"} 0
orbi_collisions_total{component="backhaul",host="http://fakehost.local",original_metric_name="bh_collisions"} 0
orbi_collisions_total{component="wifi-2.4Ghz",host="http://fakehost.local",original_metric_name="bgn_collisions"} 0
orbi_collisions_total{component="wifi-5GHz",host="http://fakehost.local",original_metric_name="an_collisions"} 0
# HELP orbi_received_bytes_per_second Received Bytes per second
# TYPE orbi_received_bytes_per_second gauge
orbi_received_bytes_per_second{component="LAN",host="http://fakehost.local",original_metric_name="lan_rxbs"} 729
orbi_received_bytes_per_second{component="WAN",host="http://fakehost.local",original_metric_name="wan_rxbs"} 1399
orbi_received_bytes_per_second{component="backhaul",host="http://fakehost.local",original_metric_name="bh_rxbs"} 0
orbi_received_bytes_per_second{component="wifi-2.4Ghz",host="http://fakehost.local",original_metric_name="bgn_rxbs"} 6
orbi_received_bytes_per_second{component="wifi-5GHz",host="http://fakehost.local",original_metric_name="an_rxbs"} 0
# HELP orbi_received_packets_total Packets Received
# TYPE orbi_received_packets_total counter
orbi_received_packets_total{component="LAN",host="http://fakehost.local",original_metric_name="lan_rxpkts"} 4.382645e+07
orbi_received_packets_total{component="WAN",host="http://fakehost.local",original_metric_name="wan_rxpkts"} 2.6269307e+08
orbi_received_packets_total{component="backhaul",host="http://fakehost.local",original_metric_name="bh_rxpkts"} 4094
orbi_received_packets_total{component="wifi-2.4Ghz",host="http://fakehost.local",original_metric_name="bgn_rxpkts"} 60138
orbi_received_packets_total{component="wifi-5GHz",host="http://fakehost.local",original_metric_name="an_rxpkts"} 0
# HELP orbi_system_uptime_seconds_total System Uptime in Seconds
# TYPE orbi_system_uptime_seconds_total counter
orbi_system_uptime_seconds_total{host="http://fakehost.local",original_metric_name="sys_uptime"} 3.069182e+06
# HELP orbi_transmitted_bytes_per_second Transmitted Bytes per second
# TYPE orbi_transmitted_bytes_per_second gauge
orbi_transmitted_bytes_per_second{component="LAN",host="http://fakehost.local",original_metric_name="lan_txbs"} 1345
orbi_transmitted_bytes_per_second{component="WAN",host="http://fakehost.local",original_metric_name="wan_txbs"} 1399
orbi_transmitted_bytes_per_second{component="backhaul",host="http://fakehost.local",original_metric_name="bh_txbs"} 0
orbi_transmitted_bytes_per_second{component="wifi-2.4Ghz",host="http://fakehost.local",original_metric_name="bgn_txbs"} 32
orbi_transmitted_bytes_per_second{component="wifi-5GHz",host="http://fakehost.local",original_metric_name="an_txbs"} 0
# HELP orbi_transmitted_packets_total Packets Transmitted
# TYPE orbi_transmitted_packets_total counter
orbi_transmitted_packets_total{component="LAN",host="http://fakehost.local",original_metric_name="lan_txpkts"} 8.8670935e+07
orbi_transmitted_packets_total{component="WAN",host="http://fakehost.local",original_metric_name="wan_txpkts"} 1.21869567e+08
orbi_transmitted_packets_total{component="backhaul",host="http://fakehost.local",original_metric_name="bh_txpkts"} 3.23651e+07
orbi_transmitted_packets_total{component="wifi-2.4Ghz",host="http://fakehost.local",original_metric_name="bgn_txpkts"} 3.1060602e+07
orbi_transmitted_packets_total{component="wifi-5GHz",host="http://fakehost.local",original_metric_name="an_txpkts"} 1.7510305e+07
# HELP orbi_uptime_seconds_total Component uptime in seconds
# TYPE orbi_uptime_seconds_total counter
orbi_uptime_seconds_total{component="LAN",host="http://fakehost.local",original_metric_name="lan_systime"} 3.069182e+06
orbi_uptime_seconds_total{component="WAN",host="http://fakehost.local",original_metric_name="wan_systime"} 2695
orbi_uptime_seconds_total{component="backhaul",host="http://fakehost.local",original_metric_name="bh_systime"} 3.069093e+06
orbi_uptime_seconds_total{component="wifi-2.4Ghz",host="http://fakehost.local",original_metric_name="bgn_systime"} 3.06911e+06
orbi_uptime_seconds_total{component="wifi-5GHz",host="http://fakehost.local",original_metric_name="an_systime"} 3.069093e+06
 
`

type fakeClient struct {
}

func (f fakeClient) GetDevices() ([]orbi.Device, error) {
	return []orbi.Device{
		{
			IP:                "10.0.0.10",
			MAC:               "AA:BB:CC:00:11:22",
			ConnectionType:    "5G",
			AttachType:        "0",
			DeviceType:        "2",
			DeviceTypeName:    "Laptop",
			Model:             "Generic",
			Name:              "My Laptop",
			ActiveStatus:      "0",
			ConnectedOrbiName: "Orbi Router",
			ConnectedOrbiMAC:  "FF:FF:FF:FF:FF:FF",
			BackhaulStatus:    "Good",
			LedState:          "0",
			LedFunc:           "0",
			SyncBtn:           "0",
			Uprate:            "0.00",
			Downrate:          "0.00",
			ModuleName:        "",
		},
	}, nil
}

func (f fakeClient) GetMetrics() (map[string]*orbi.Metric, error) {
	return cannedResponse, nil
}

func (f fakeClient) Host() string {
	return "http://fakehost.local"
}
