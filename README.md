# orbi-exporter
Prometheus Exporter for Orbi Metrics


## Origin Story

My Orbi Router doesn't expose metrics in Prometheus format, but it does expose some network metrics via its control panel.
The control panel is a web application running on the router protected with HTTP Basic Auth. Although it responds in HTML,
the embedded javascript variables are easy to parse.

this is built to run via the container service on my synology NAS

## Configuration

The exporter listens on port `6724` (on a telephone keypad, that's `O:6 R:7 B:2 I:4`)

The exporter expects the following environment variables:

```
ORBI_HOST=http://192.168.1.1
ORBI_BASICAUTH_USER=admin
ORBI_BASICAUTH_PASSWORD=<your password>
```

## Compatibility

Tested against an Orbi `RBR20`. Router Firmware Version `V2.7.4.24`


## Metrics Provided

Beyond parsing the metrics, this exporter also adapts their structure to reflect Prometheus [naming conventions and recommended patterns](https://prometheus.io/docs/practices/naming/) for metrics and labels:

* including the unit in the metric name (sys_uptime -> "system_uptime_seconds")
* "received_packets" using labels -> "lan_rxpkts" "wan_rxpkts", etc are all "received_packets" and can be aggregated easily

`received_bytes_per_second`
`transmitted_bytes_per_second`
`received_packets`
`transmitted_packets`
`collisions`
`uptime_seconds`
`system_uptime_seconds`

## Labels

* `component`: the network within the router for which the metric applies
* values: "LAN", "WAN", "backhaul", "wifi-2.4Ghz", "wifi-5Ghz"
* `original_metric_name`: the name of the JavaScript variable used for this metric
* `host`: the ip address of the router (the same value as `ORBI_HOST`)
