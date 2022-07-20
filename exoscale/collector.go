package exoscale

import (
	"context"
	"time"

	"github.com/SquareFactory/exoscale_exporter/log"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"
)

type Collector struct {
	computeAPI *ComputeAPI
}

func NewCollector(computeAPI *ComputeAPI) *Collector {
	if computeAPI == nil {
		log.Logger.Panic("computeAPI is nil")
	}
	return &Collector{
		computeAPI: computeAPI,
	}
}

func (c *Collector) RecordMetrics() {
	for {
		func() {
			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			vms, err := c.computeAPI.ListVirtualMachines(ctx)
			if err != nil {
				log.Logger.Error("ListVirtualMachines thrown an error", zap.Error(err))
				return
			}

			for _, vm := range vms {
				lifetime, err := ComputeLifetime(vm)
				if err != nil {
					log.Logger.Error("ComputeLifetime thrown an error", zap.Error(err))
					return
				}
				lifetimeGauge.With(prometheus.Labels{
					"instance": vm.DisplayName,
					"id":       vm.ID.String(),
				}).Set(lifetime.Seconds())
			}

			serversGauge.Set(float64(len(vms)))
		}()
		time.Sleep(time.Second)
	}
}
