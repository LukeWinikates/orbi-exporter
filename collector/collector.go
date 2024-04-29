package collector

import (
	"fmt"
	"github.com/LukeWinikates/orbi-exporter/orbi"
	"strconv"
	"strings"
)

type OrbiMetricTranslation struct {
	SystemUptimeSeconds float64
	LanCollisions       float64
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

func Translate(metrics map[string]orbi.Metric) OrbiMetricTranslation {
	return OrbiMetricTranslation{
		SystemUptimeSeconds: toNumber(metrics["sys_uptime"].Value),
		LanCollisions:       toNumber(metrics["lan_collisions"].Value),
	}
}
