package collector

import (
	"github.com/LukeWinikates/orbi-exporter/orbi"
	"github.com/prometheus/client_golang/prometheus"
	"html"
	"log"
)

type devicesCollector struct {
	orbiClient orbi.Client
}

func (o *devicesCollector) Describe(descs chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(o, descs)
}

func (o *devicesCollector) Collect(metrics chan<- prometheus.Metric) {
	log.Println("beginning device metrics collection")
	devices, err := o.orbiClient.GetDevices()
	if err != nil {
		log.Printf("error collecting devices: %s\n", err.Error())
		return
	}

	for _, device := range devices {
		metrics <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				"orbi_connected_device_downrate",
				"Unknown",
				[]string{"host"},
				deviceLabels(device),
			),
			prometheus.GaugeValue,
			device.DownrateFloat(),
			o.orbiClient.Host())
		metrics <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				"orbi_connected_device_uprate",
				"Unknown",
				[]string{"host"},
				deviceLabels(device),
			),
			prometheus.GaugeValue,
			device.UprateFloat(),
			o.orbiClient.Host())
		metrics <- prometheus.MustNewConstMetric(
			prometheus.NewDesc(
				"orbi_connected_device_info",
				"is the device active",
				[]string{"host"},
				deviceLabels(device),
			),
			prometheus.GaugeValue,
			1,
			o.orbiClient.Host())
	}
	log.Println("finished collecting devices")
}

func deviceLabels(device orbi.Device) prometheus.Labels {
	return prometheus.Labels{
		"ip":              device.IP,
		"connection_type": device.ConnectionType,
		"device_type":     device.DeviceTypeName,
		"name":            html.UnescapeString(device.Name),
		"connected_orbi":  device.ConnectedOrbiName,
		"model":           html.UnescapeString(device.Model),
	}
}

func NewDevicesCollector(client orbi.Client) prometheus.Collector {
	return &devicesCollector{
		orbiClient: client,
	}
}
