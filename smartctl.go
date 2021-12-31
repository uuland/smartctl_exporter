package main

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

// SMARTDevice - short info about device
type SMARTDevice struct {
	device string
	proto  string
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
		proto:  strings.TrimSpace(smart.json.Get("device.protocol").String()),
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
	smart.mineSmartStatus()
	smart.minePowerCycleCount()
	smart.minePowerOnSeconds()
	smart.mineTemperatures()

	smart.minePercentageUsed()
	smart.mineBytesRead()
	smart.mineCommandsRead()
	smart.mineBytesWritten()
	smart.mineCommandsWritten()

	if smart.isNVMe() {
		smart.mineNVMeHealthInformationLog()
	} else {
		smart.mineRotationRate()
		smart.mineInterfaceSpeed()
		smart.mineDeviceAttribute()
		smart.mineDeviceSCTStatus()
		smart.mineDeviceStatistics()
		smart.mineDeviceErrorLog()
		smart.mineDeviceSelfTestLog()
		smart.mineDeviceERC()
	}

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

func (smart *SMARTctl) isNVMe() bool {
	return smart.device.proto == "NVMe"
}
