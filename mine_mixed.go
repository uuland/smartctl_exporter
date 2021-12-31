package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

func (smart *SMARTctl) minePercentageUsed() {
	var used gjson.Result
	if smart.isNVMe() {
		used = smart.json.Get("nvme_smart_health_information_log.percentage_used")
	} else {
		used = getATAStatisticsValue(smart.json, "Solid State Device Statistics", "Percentage Used Endurance Indicator")
	}

	smart.ch <- prometheus.MustNewConstMetric(
		metricDevicePercentageUsed,
		prometheus.CounterValue,
		used.Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineBytesRead() {
	blockSize := smart.json.Get("logical_block_size").Float()
	if blockSize == 0 {
		blockSize = 512 * 1024
	}

	var size gjson.Result
	if smart.isNVMe() {
		size = smart.json.Get("nvme_smart_health_information_log.data_units_read")
	} else {
		size = getATAStatisticsValue(smart.json, "General Statistics", "Logical Sectors Read")
	}

	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceBytesRead,
		prometheus.CounterValue,
		size.Float()*blockSize,
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineBytesWritten() {
	blockSize := smart.json.Get("logical_block_size").Float()
	if blockSize == 0 {
		blockSize = 512 * 1024
	}

	var size gjson.Result
	if smart.isNVMe() {
		size = smart.json.Get("nvme_smart_health_information_log.data_units_written")
	} else {
		size = getATAStatisticsValue(smart.json, "General Statistics", "Logical Sectors Written")
	}

	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceBytesWritten,
		prometheus.CounterValue,
		size.Float()*blockSize,
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineCommandsRead() {
	var count gjson.Result
	if smart.isNVMe() {
		count = smart.json.Get("nvme_smart_health_information_log.host_reads")
	} else {
		count = getATAStatisticsValue(smart.json, "General Statistics", "Number of Read Commands")
	}

	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceCommandsRead,
		prometheus.CounterValue,
		count.Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineCommandsWritten() {
	var count gjson.Result
	if smart.isNVMe() {
		count = smart.json.Get("nvme_smart_health_information_log.host_writes")
	} else {
		count = getATAStatisticsValue(smart.json, "General Statistics", "Number of Write Commands")
	}

	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceCommandsWritten,
		prometheus.CounterValue,
		count.Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}
