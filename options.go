package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	yaml "gopkg.in/yaml.v2"
)

var (
	exporterVersion = "0.5"
)

// SMARTOptions is a inner representation of a options
type SMARTOptions struct {
	BindTo           string   `yaml:"bind_to"`
	PushTo           string   `yaml:"push_to"`
	URLPath          string   `yaml:"url_path"`
	FakeJSON         bool     `yaml:"fake_json"`
	SMARTctlLocation string   `yaml:"smartctl_location"`
	PushInterval     string   `yaml:"push_interval"`
	CollectPeriod    string   `yaml:"collect_not_more_than_period"`
	Devices          []string `yaml:"devices"`
}

// Options is a representation of a options
type Options struct {
	SMARTctl              SMARTOptions `yaml:"smartctl_exporter"`
	PushIntervalDuration  time.Duration
	CollectPeriodDuration time.Duration
}

// Parse options from yaml config file
func loadOptions() Options {
	configFile := flag.String("config", "/etc/smartctl_exporter.yaml", "Path to smartctl_exporter config file")
	verbose := flag.Bool("verbose", false, "Verbose log output")
	debug := flag.Bool("debug", false, "Debug log output")
	version := flag.Bool("version", false, "Show application version and exit")
	flag.Parse()

	if *version {
		fmt.Printf("smartctl_exporter version: %s\n", exporterVersion)
		os.Exit(0)
	}

	logger = newLogger(*verbose, *debug)

	logger.Verbose("Read options from %s\n", *configFile)
	yamlFile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		logger.Panic("Failed read %s: %s", configFile, err)
	}

	opts := Options{
		SMARTctl: SMARTOptions{
			BindTo:           "9633",
			URLPath:          "/metrics",
			FakeJSON:         false,
			SMARTctlLocation: "/usr/sbin/smartctl",
			PushInterval:     "30s",
			CollectPeriod:    "10m",
			Devices:          []string{},
		},
	}

	if yaml.Unmarshal(yamlFile, &opts) != nil {
		logger.Panic("Failed parse %s: %s", configFile, err)
	}

	if d, err := time.ParseDuration(opts.SMARTctl.PushInterval); err != nil {
		logger.Panic("Failed read push_interval (%s): %s", opts.SMARTctl.CollectPeriod, err)
	} else {
		opts.PushIntervalDuration = d
	}

	if d, err := time.ParseDuration(opts.SMARTctl.CollectPeriod); err != nil {
		logger.Panic("Failed read collect_not_more_than_period (%s): %s", opts.SMARTctl.CollectPeriod, err)
	} else {
		opts.CollectPeriodDuration = d
	}

	logger.Debug("Parsed options: %s", opts)
	return opts
}
