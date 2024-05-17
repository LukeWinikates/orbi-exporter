package collector

import (
	"github.com/LukeWinikates/orbi-exporter/orbi"
	"github.com/prometheus/client_golang/prometheus"
	"log"
)

type orbiCollector struct {
	orbiClient orbi.Client
}

func (o *orbiCollector) Describe(descs chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(o, descs)
}

func (o *orbiCollector) Collect(metrics chan<- prometheus.Metric) {
	log.Println("beginning metrics collection")
	orbiMetrics, err := o.orbiClient.GetMetrics()
	if err != nil {
		log.Printf("error collecting metrics: %s\n", err.Error())
		return
	}

	for _, metric := range orbiMetrics {
		if shouldTranslate(*metric) {
			metrics <- translateMetric(*metric, o.orbiClient.Host())
		}
	}
	log.Println("finished collecting metrics")
}

func NewStatisticsCollector(client orbi.Client) prometheus.Collector {
	return &orbiCollector{
		orbiClient: client,
	}
}
