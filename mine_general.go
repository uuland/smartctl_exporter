package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/tidwall/gjson"
)

func (smart *SMARTctl) mineDevice() {
	device := smart.json.Get("device")
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceModel,
		prometheus.GaugeValue,
		1,
		smart.device.device,
		device.Get("type").String(),
		device.Get("protocol").String(),
		smart.device.family,
		smart.device.model,
		smart.device.serial,
		getStringIfExists(smart.json, "ata_additional_product_id", "unknown"),
		smart.json.Get("firmware_version").String(),
		smart.json.Get("ata_version.string").String(),
		smart.json.Get("sata_version.string").String(),
	)
}

func (smart *SMARTctl) mineDeviceStatus() {
	status := smart.json.Get("smart_status")
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceStatus,
		prometheus.GaugeValue,
		status.Get("passed").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineSmartStatus() {
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceSmartStatus,
		prometheus.GaugeValue,
		smart.json.Get("smart_status.passed").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) minePowerCycleCount() {
	smart.ch <- prometheus.MustNewConstMetric(
		metricDevicePowerCycleCount,
		prometheus.CounterValue,
		smart.json.Get("power_cycle_count").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) minePowerOnSeconds() {
	pot := smart.json.Get("power_on_time")
	smart.ch <- prometheus.MustNewConstMetric(
		metricDevicePowerOnSeconds,
		prometheus.CounterValue,
		getFloatIfExists(pot, "hours", 0)*60*60+getFloatIfExists(pot, "minutes", 0)*60,
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
}

func (smart *SMARTctl) mineCapacity() {
	capacity := smart.json.Get("user_capacity")
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceCapacityBlocks,
		prometheus.GaugeValue,
		capacity.Get("blocks").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
	smart.ch <- prometheus.MustNewConstMetric(
		metricDeviceCapacityBytes,
		prometheus.GaugeValue,
		capacity.Get("bytes").Float(),
		smart.device.device,
		smart.device.family,
		smart.device.model,
		smart.device.serial,
	)
	for _, blockType := range []string{"logical", "physical"} {
		smart.ch <- prometheus.MustNewConstMetric(
			metricDeviceBlockSize,
			prometheus.GaugeValue,
			smart.json.Get(fmt.Sprintf("%s_block_size", blockType)).Float(),
			smart.device.device,
			smart.device.family,
			smart.device.model,
			smart.device.serial,
			blockType,
		)
	}
}

func (smart *SMARTctl) mineInterfaceSpeed() {
	iSpeed := smart.json.Get("interface_speed")
	for _, speedType := range []string{"max", "current"} {
		tSpeed := iSpeed.Get(speedType)
		smart.ch <- prometheus.MustNewConstMetric(
			metricDeviceInterfaceSpeed,
			prometheus.GaugeValue,
			tSpeed.Get("units_per_second").Float()*tSpeed.Get("bits_per_unit").Float(),
			smart.device.device,
			smart.device.family,
			smart.device.model,
			smart.device.serial,
			speedType,
		)
	}
}

func (smart *SMARTctl) mineTemperatures() {
	temperatures := smart.json.Get("temperature")
	if temperatures.Exists() {
		temperatures.ForEach(func(key, value gjson.Result) bool {
			smart.ch <- prometheus.MustNewConstMetric(
				metricDeviceTemperature,
				prometheus.GaugeValue,
				value.Float(),
				smart.device.device,
				smart.device.family,
				smart.device.model,
				smart.device.serial,
				key.String(),
			)
			return true
		})
	}
}

func (smart *SMARTctl) mineRotationRate() {
	rRate := getFloatIfExists(smart.json, "rotation_rate", 0)
	if rRate > 0 {
		smart.ch <- prometheus.MustNewConstMetric(
			metricDeviceRotationRate,
			prometheus.GaugeValue,
			rRate,
			smart.device.device,
			smart.device.family,
			smart.device.model,
			smart.device.serial,
		)
	}
}
