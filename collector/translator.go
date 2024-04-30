package collector

import (
	"fmt"
	"github.com/LukeWinikates/orbi-exporter/orbi"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"strings"
)

type OrbiMetricTranslation struct {
	Name  string
	Value float64
}

func toM(num float64) float64 {
	return num * 1000_000
}

func toNumber(s string) float64 {
	if s == "1000M/Full" {
		return toM(1000)
	}

	if s == "Link up" {
		return 1
	}

	if s == "Link down" {
		return 0
	}

	if strings.HasSuffix(s, "M") {
		return toM(toNumber(strings.TrimSuffix(s, "M")))
	}

	float, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Printf("unexpected number format: %s\n", s)
	}
	return float
}

type prefixContext struct {
	For   string
	Label string
}

var prefixContexts = map[string]prefixContext{
	"lan": {
		Label: "LAN",
		For:   "wired LAN",
	},
	"wan": {
		Label: "WAN",
		For:   "WAN",
	},
	"bgn": {
		Label: "wifi-2.4Ghz",
		For:   "WIFI 802.11b/g/n 2.4GHz",
	},
	"an": {
		Label: "wifi-5GHz",
		For:   "WIFI 802.11a/n 5GHz",
	},
	"bh": {
		Label: "backhaul",
		For:   "Backhaul",
	},
}

type metricTranslation struct {
	Help      string
	Name      string
	ValueType prometheus.ValueType
}

var metricTranslations = map[string]metricTranslation{
	"txpkts": {
		Name:      "transmitted_packets",
		Help:      "Packets Transmitted",
		ValueType: prometheus.CounterValue,
	},
	"rxpkts": {
		Name:      "received_packets",
		Help:      "Packets Received",
		ValueType: prometheus.CounterValue,
	},
	"collisions": {
		Name:      "collisions",
		Help:      "Collisions",
		ValueType: prometheus.CounterValue,
	},
	"txbs": {
		Name:      "transmitted_bytes_per_second",
		Help:      "Transmitted Bytes per second",
		ValueType: prometheus.GaugeValue,
	},
	"rxbs": {
		Name:      "received_bytes_per_second",
		Help:      "Received Bytes per second",
		ValueType: prometheus.GaugeValue,
	},
	"systime": {
		Name:      "uptime_seconds",
		Help:      "Component uptime in seconds",
		ValueType: prometheus.CounterValue,
	},
}

func prefixFromMetricName(metricName string) prefixContext {
	prefix := strings.Split(metricName, "_")[0]
	return prefixContexts[prefix]
}

func translationFromMetricName(metricName string) metricTranslation {
	metric := strings.Split(metricName, "_")[1]
	return metricTranslations[metric]
}

func shouldTranslate(metric orbi.Metric) bool {
	if strings.HasSuffix(metric.Name, "status") {
		return false
	}
	if metric.Name == "sys_uptime" {
		return true
	}

	splits := strings.Split(metric.Name, "_")
	_, prefixExists := prefixContexts[splits[0]]
	_, metricNameExists := metricTranslations[splits[1]]

	return prefixExists && metricNameExists
}

func translateMetric(metric orbi.Metric, host string) prometheus.Metric {
	if metric.Name == "sys_uptime" {
		return prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				"system_uptime_seconds",
				"System Uptime in Seconds",
				[]string{"host"},
				prometheus.Labels{
					"original_metric_name": metric.Name,
				},
			),
			prometheus.CounterValue,
			toNumber(metric.Value),
			host)
	}

	prefix := prefixFromMetricName(metric.Name)
	translation := translationFromMetricName(metric.Name)

	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			translation.Name,
			translation.Help,
			[]string{"host", "component"},
			prometheus.Labels{
				"original_metric_name": metric.Name,
			},
		),
		translation.ValueType,
		toNumber(metric.Value),
		host, prefix.Label)
}
