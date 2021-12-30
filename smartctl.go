package main

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

// SMARTDevice - short info about device
type SMARTDevice struct {
	device string
	serial string
	family string
	model  string
}

// SMARTctl object
type SMARTctl struct {
	ch     chan<- prometheus.Metric
	json   gjson.Result
	device SMARTDevice
}

// NewSMARTctl is smartctl constructor
func NewSMARTctl(json gjson.Result, ch chan<- prometheus.Metric) SMARTctl {
	smart := SMARTctl{}
	smart.ch = ch
	smart.json = json
	smart.device = SMARTDevice{
		device: strings.TrimSpace(smart.json.Get("device.name").String()),
		serial: strings.TrimSpace(smart.json.Get("serial_number").String()),
		family: strings.TrimSpace(smart.json.Get("model_family").String()),
		model:  strings.TrimSpace(smart.json.Get("model_name").String()),
	}
	return smart
}

// Collect metrics
func (smart *SMARTctl) Collect() {
	logger.Verbose("Collecting metrics from %s: %s, %s", smart.device.device, smart.device.family, smart.device.model)
	smart.mineExitStatus()
	smart.mineDevice()
	smart.mineCapacity()
	smart.mineInterfaceSpeed()
	smart.mineDeviceAttribute()
	smart.minePowerOnSeconds()
	smart.mineRotationRate()
	smart.mineTemperatures()
	smart.minePowerCycleCount()
	smart.mineDeviceSCTStatus()
	smart.mineDeviceStatistics()
	smart.mineNvmeSmartHealthInformationLog()
	smart.mineNvmeSmartStatus()
	smart.mineDeviceStatus()
	smart.mineDeviceErrorLog()
	smart.mineDeviceSelfTestLog()
	smart.mineDeviceERC()
	smart.minePercentageUsed()
	smart.mineAvailableSpare()
	smart.mineAvailableSpareThreshold()
	smart.mineCriticalWarning()
	smart.mineMediaErrors()
	smart.mineNumErrLogEntries()
	smart.mineBytesRead()
	smart.mineBytesWritten()
	smart.mineSmartStatus()

}

func (smart *SMARTctl) mineExitStatus() {
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceExitStatus,
		prometheus.GaugeValue,
		smart.json.Get("smartctl.exit_status").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}
