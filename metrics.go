package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	metricSmartctlVersion = prometheus.NewDesc(
		"smartctl_version",
		"smartctl version",
		[]string{
			"json_format_version",
			"smartctl_version",
			"svn_revision",
			"build_info",
		},
		nil,
	)
	metricDeviceExitStatus = prometheus.NewDesc(
		"smartctl_exit_status",
		"Exit status of smartctl on device",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceModel = prometheus.NewDesc(
		"smartctl_device",
		"Device info",
		[]string{
			"device",
			"interface",
			"protocol",
			"model_family",
			"model_name",
			"serial_number",
			"ata_additional_product_id",
			"firmware_version",
			"ata_version",
			"sata_version",
		},
		nil,
	)
	metricDeviceCapacityBlocks = prometheus.NewDesc(
		"smartctl_device_capacity_blocks",
		"Device capacity in blocks",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceCapacityBytes = prometheus.NewDesc(
		"smartctl_device_capacity_bytes",
		"Device capacity in bytes",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceBlockSize = prometheus.NewDesc(
		"smartctl_device_block_size",
		"Device block size",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"blocks_type",
		},
		nil,
	)
	metricDeviceInterfaceSpeed = prometheus.NewDesc(
		"smartctl_device_interface_speed",
		"Device interface speed, bits per second",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"speed_type",
		},
		nil,
	)
	metricDeviceAttribute = prometheus.NewDesc(
		"smartctl_device_attribute",
		"Device attributes",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"attribute_name",
			"attribute_flags_short",
			"attribute_flags_long",
			"attribute_value_type",
			"attribute_id",
		},
		nil,
	)
	metricDeviceStatistics = prometheus.NewDesc(
		"smartctl_device_statistics",
		"Device statistics",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"statistic_table",
			"statistic_name",
			"statistic_flags_short",
			"statistic_flags_long",
		},
		nil,
	)
	metricNVMeHealthLog = prometheus.NewDesc(
		"smartctl_device_nvme_health",
		"NVMe Device health information log",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"health_log_name",
		},
		nil,
	)
	metricDevicePowerOnSeconds = prometheus.NewDesc(
		"smartctl_device_power_on_seconds",
		"Device power on seconds",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceRotationRate = prometheus.NewDesc(
		"smartctl_device_rotation_rate",
		"Device rotation rate",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceTemperature = prometheus.NewDesc(
		"smartctl_device_temperature",
		"Device temperature celsius",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"temperature_type",
		},
		nil,
	)
	metricDevicePowerCycleCount = prometheus.NewDesc(
		"smartctl_device_power_cycle_count",
		"Device power cycle count",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceSmartStatus = prometheus.NewDesc(
		"smartctl_device_smart_status",
		"General smart status",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceState = prometheus.NewDesc(
		"smartctl_device_state",
		"Device state (0=active, 1=standby, 2=sleep, 3=dst, 4=offline, 5=sct)",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceErrorLogCount = prometheus.NewDesc(
		"smartctl_device_error_log_count",
		"Device SMART error log count",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"error_log_type",
		},
		nil,
	)
	metricDeviceSelfTestLogCount = prometheus.NewDesc(
		"smartctl_device_self_test_log_count",
		"Device SMART self test log count",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"self_test_log_type",
		},
		nil,
	)
	metricDeviceSelfTestLogErrorCount = prometheus.NewDesc(
		"smartctl_device_self_test_log_error_count",
		"Device SMART self test log error count",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"self_test_log_type",
		},
		nil,
	)
	metricDeviceERCSeconds = prometheus.NewDesc(
		"smartctl_device_erc_seconds",
		"Device SMART Error Recovery Control Seconds",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
			"op_type",
		},
		nil,
	)
	metricDevicePercentageUsed = prometheus.NewDesc(
		"smartctl_device_percentage_used",
		"Device write percentage used",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceBytesRead = prometheus.NewDesc(
		"smartctl_device_bytes_read",
		"",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceBytesWritten = prometheus.NewDesc(
		"smartctl_device_bytes_written",
		"",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceCommandsRead = prometheus.NewDesc(
		"smartctl_device_commands_read",
		"",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
	metricDeviceCommandsWritten = prometheus.NewDesc(
		"smartctl_device_commands_written",
		"",
		[]string{
			"device",
			"model_family",
			"model_name",
			"serial_number",
		},
		nil,
	)
)
