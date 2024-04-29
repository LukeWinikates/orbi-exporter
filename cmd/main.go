package main

import (
	"github.com/LukeWinikates/orbi-exporter/collector"
	"github.com/LukeWinikates/orbi-exporter/orbi"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	orbiClient, err := orbi.NewClient(os.Getenv("ORBI_HOST"), os.Getenv("ORBI_BASICAUTH_USER"), os.Getenv("ORBI_BASICAUTH_PASSWORD"))
	if err != nil {
		log.Fatal(err)
	}
	promauto.NewCounterFunc(prometheus.CounterOpts{
		Name: "system_uptime_seconds",
	}, func() float64 {
		metrics, err := orbiClient.GetMetrics()

		if err != nil {
			return 0
		}
		return collector.Translate(metrics).SystemUptimeSeconds
	})
	http.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe("localhost:6724", nil))
}
