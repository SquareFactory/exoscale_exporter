package exoscale

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var ns = "exoscale"

var (
	lifetimeGauge = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: ns,
		Name:      "lifetime_seconds",
		Subsystem: "compute",
		Help:      "Instance lifetime in seconds",
	}, []string{"id", "instance"})
	serversGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: ns,
		Name:      "total",
		Subsystem: "compute",
		Help:      "Number of instances",
	})
)
