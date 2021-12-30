package main

import "github.com/prometheus/client_golang/prometheus"

func (smart *SMARTctl) mineNvmeSmartStatus() {
	iStatus := smart.json.Get("smart_status")
	smart.ch <- prometheus.MustNewConstMetric(
		metricSmartStatus,
		prometheus.GaugeValue,
		iStatus.Get("passed").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineNvmeSmartHealthInformationLog() {
	iHealth := smart.json.Get("nvme_smart_health_information_log")
	smart.ch <- prometheus.MustNewConstMetric(
		metricCriticalWarning,
		prometheus.GaugeValue,
		iHealth.Get("critical_warning").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
	smart.ch <- prometheus.MustNewConstMetric(
		metricAvailableSpare,
		prometheus.GaugeValue,
		iHealth.Get("available_spare").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
	smart.ch <- prometheus.MustNewConstMetric(
		metricMediaErrors,
		prometheus.GaugeValue,
		iHealth.Get("media_errors").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) minePercentageUsed() {
	smart.ch <- prometheus.MustNewConstMetric(
		metricDevicePercentageUsed,
		prometheus.CounterValue,
		smart.json.Get("nvme_smart_health_information_log.percentage_used").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineAvailableSpare() {
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceAvailableSpare,
		prometheus.CounterValue,
		smart.json.Get("nvme_smart_health_information_log.available_spare").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineAvailableSpareThreshold() {
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceAvailableSpareThreshold,
		prometheus.CounterValue,
		smart.json.Get("nvme_smart_health_information_log.available_spare_threshold").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineCriticalWarning() {
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceCriticalWarning,
		prometheus.CounterValue,
		smart.json.Get("nvme_smart_health_information_log.critical_warning").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineMediaErrors() {
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceMediaErrors,
		prometheus.CounterValue,
		smart.json.Get("nvme_smart_health_information_log.media_errors").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineNumErrLogEntries() {
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceNumErrLogEntries,
		prometheus.CounterValue,
		smart.json.Get("nvme_smart_health_information_log.num_err_log_entries").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineBytesRead() {
	blockSize := smart.json.Get("logical_block_size").Float() * 1024
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceBytesRead,
		prometheus.CounterValue,
		smart.json.Get("nvme_smart_health_information_log.data_units_read").Float()*blockSize,
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineBytesWritten() {
	blockSize := smart.json.Get("logical_block_size").Float() * 1024
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceBytesWritten,
		prometheus.CounterValue,
		smart.json.Get("nvme_smart_health_information_log.data_units_written").Float()*blockSize,
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}
