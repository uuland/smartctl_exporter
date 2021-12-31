package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

func (smart *SMARTctl) mineNVMeHealthInformationLog() {
	write := func(name string, value float64) {
		smart.ch <- prometheus.MustNewConstMetric(
			metricNVMeHealthLog,
			prometheus.GaugeValue,
			value,
			smart.device.device,
			smart.device.family,
			smart.device.model,
			smart.device.serial,
			name,
		)
	}

	for name, log := range smart.json.Get("nvme_smart_health_information_log").Map() {
		if log.IsArray() {
			for idx, value := range log.Array() {
				write(fmt.Sprintf("%s_%d", name, idx), value.Float())
			}
		} else {
			write(name, log.Float())
		}
	}
}
