package collector

import (
	"fmt"
	"github.com/LukeWinikates/orbi-exporter/orbi"
	"github.com/prometheus/client_golang/prometheus"
)

type orbiCollector struct {
	orbiClient orbi.Client
}

func (o *orbiCollector) Describe(descs chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(o, descs)
}

func (o *orbiCollector) Collect(metrics chan<- prometheus.Metric) {
	orbiMetrics, err := o.orbiClient.GetMetrics()
	if err != nil {
		fmt.Printf("error collecting metrics: %s\n", err.Error())
	}

	for _, metric := range orbiMetrics {
		if shouldTranslate(*metric) {
			metrics <- translateMetric(*metric, o.orbiClient.Host())
		}
	}

}

func NewCollector(client orbi.Client) prometheus.Collector {
	return &orbiCollector{
		orbiClient: client,
	}
}
