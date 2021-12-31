package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/client_golang/prometheus/push"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	options Options
	logger  Logger
)

// SMARTctlManagerCollector implements the Collector interface.
type SMARTctlManagerCollector struct {
}

// Describe sends the super-set of all possible descriptors of metrics
func (i SMARTctlManagerCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(i, ch)
}

// Collect is called by the Prometheus registry when collecting metrics.
func (i SMARTctlManagerCollector) Collect(ch chan<- prometheus.Metric) {
	info := NewSMARTctlInfo(ch)
	success := 0

	for _, device := range options.SMARTctl.Devices {
		if json, err := readData(device); err == nil {
			info.SetJSON(json)
			smart := NewSMARTctl(json, ch)
			smart.Collect()
			success++
		} else {
			logger.Error(err.Error())
		}
	}

	if success > 0 {
		info.Collect()
	}
}

func init() {
	options = loadOptions()

	if len(options.SMARTctl.Devices) == 0 {
		logger.Debug("No devices specified, trying to load them automatically")
		json, err := readSMARTctlDevices()
		if err != nil {
			logger.Panic("Can't find any devices")
			os.Exit(1)
		}
		devices := json.Get("devices").Array()
		for _, d := range devices {
			device := d.Get("name").String()
			logger.Debug("Found device: %s", device)
			options.SMARTctl.Devices = append(options.SMARTctl.Devices, device)
		}
	}
}

func main() {
	reg := prometheus.NewPedanticRegistry()
	reg.MustRegister(
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
		prometheus.NewGoCollector(),
	)

	prometheus.WrapRegistererWithPrefix("", reg).MustRegister(SMARTctlManagerCollector{})

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	if options.SMARTctl.PushTo == "" {
		/**
		metrics http server
		*/

		logger.Info("Starting on %s%s", options.SMARTctl.BindTo, options.SMARTctl.URLPath)
		http.Handle(options.SMARTctl.URLPath, promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

		srv := &http.Server{Addr: options.SMARTctl.BindTo}
		go func() {
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				logger.Error("Listen failed: %s", err)
				os.Exit(1)
			}
		}()

		logger.Info("Metrics server started")

		<-done
		if err := srv.Shutdown(context.TODO()); err != nil {
			logger.Error("Shutdown failed: %s", err)
			os.Exit(1)
		}
	} else {
		/**
		metrics push agent
		*/

		pusher := push.New(options.SMARTctl.PushTo, "smartctl_exporter").Gatherer(reg)

		doPush := func() {
			if err := pusher.Push(); err != nil {
				logger.Warning("Push failed: %s", err)
			} else {
				logger.Verbose("Push success")
			}
		}

		ticker := time.NewTicker(options.PushIntervalDuration)
		stopped := make(chan struct{}, 1)
		go func() {
			doPush()
			for {
				select {
				case <-stopped:
					return
				case <-ticker.C:
					doPush()
				}
			}
		}()

		logger.Info("Metrics push to -> %s", options.SMARTctl.PushTo)

		<-done
		ticker.Stop()
		close(stopped)
	}
}
