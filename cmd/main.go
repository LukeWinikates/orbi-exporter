package main

import (
	"github.com/LukeWinikates/orbi-exporter/collector"
	"github.com/LukeWinikates/orbi-exporter/orbi"
	"github.com/prometheus/client_golang/prometheus"
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
	reg := prometheus.NewRegistry()
	err = reg.Register(collector.NewCollector(orbiClient))
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

	log.Fatal(http.ListenAndServe("localhost:6724", nil))
}
