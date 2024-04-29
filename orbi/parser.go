package orbi

import (
	"bufio"
	"io"
	"regexp"
	"strings"
)

var varMatcher = regexp.MustCompile(`\s?var\s+([a-z0-9_]+)\s?=\s?"(.*)"`)
var allowList = map[string]bool{
	"sys_uptime":       true,
	"lan_status":       true,
	"lan_txpkts":       true,
	"lan_rxpkts":       true,
	"lan_collisions":   true,
	"lan_txbs":         true,
	"lan_rxbs":         true,
	"lan_systime":      true,
	"wan_status":       true,
	"wan_txpkts":       true,
	"wan_rxpkts":       true,
	"wan_collisions":   true,
	"wan_txbs":         true,
	"wan_rxbs":         true,
	"wan_systime":      true,
	"bgn_status":       true,
	"bgn_txpkts":       true,
	"bgn_rxpkts":       true,
	"bgn_collisions":   true,
	"bgn_txbs":         true,
	"bgn_rxbs":         true,
	"bgn_systime":      true,
	"an_status":        true,
	"an_txpkts":        true,
	"an_rxpkts":        true,
	"an_collisions":    true,
	"an_txbs":          true,
	"an_rxbs":          true,
	"an_systime":       true,
	"bh_status":        true,
	"bh_txpkts":        true,
	"bh_rxpkts":        true,
	"bh_collisions":    true,
	"bh_txbs":          true,
	"bh_rxbs":          true,
	"bh_systime":       true,
	"wwan0_status":     true,
	"wwan0_txpkts":     true,
	"wwan0_rxpkts":     true,
	"wwan0_collisions": true,
	"wwan0_txbs":       true,
	"wwan0_rxbs":       true,
	"wwan0_systime":    true,
	"lan_status0":      true,
	"lan0_systime":     true,
	"lan_status1":      true,
	"lan1_systime":     true,
	"lan_status2":      true,
	"lan2_systime":     true,
	"lan_status3":      true,
	"lan3_systime":     true,
}

func parse(body io.Reader) (map[string]Metric, error) {
	metrics := make(map[string]Metric)
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		if metric, ok := readMetricLine(scanner.Text()); ok {
			if isInAllowList(metric) {
				metrics[metric.Name] = *metric
			}
		}

		if shouldEnd(scanner.Text()) {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return metrics, nil
}

func shouldEnd(text string) bool {
	return strings.Contains(text, "timereset")
}

func isInAllowList(metric *Metric) bool {
	return allowList[metric.Name]
}

func readMetricLine(text string) (*Metric, bool) {
	matches := varMatcher.FindStringSubmatch(text)
	if len(matches) == 0 {
		return nil, false
	}

	return &Metric{
		Name:  matches[1],
		Value: matches[2],
	}, true
}
