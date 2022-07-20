package exoscale

import (
	"context"

	"github.com/exoscale/egoscale"
)

type ComputeAPI struct {
	cs *egoscale.Client
}

func NewComputeAPI(key string, secret string) *ComputeAPI {
	cs := egoscale.NewClient("https://api.exoscale.com/compute", key, secret, egoscale.WithoutV2Client())
	return &ComputeAPI{
		cs: cs,
	}
}

func (c *ComputeAPI) ListVirtualMachines(ctx context.Context) ([]egoscale.VirtualMachine, error) {
	req := &egoscale.ListVirtualMachines{}
	resp, err := c.cs.RequestWithContext(ctx, req)
	if err != nil {
		return []egoscale.VirtualMachine{}, err
	}
	vms := resp.(*egoscale.ListVirtualMachinesResponse)

	return vms.VirtualMachine, err
}
