package exoscale

import (
	"time"

	"github.com/exoscale/egoscale"
)

const exoscaleTimeFormat = "2006-01-02T15:04:05-0700"

func ComputeLifetime(vm egoscale.VirtualMachine) (time.Duration, error) {
	created, err := time.Parse(exoscaleTimeFormat, vm.Created)
	if err != nil {
		return time.Duration(0), err
	}
	return time.Since(created), nil
}

func MapToTags(m map[string]string) []egoscale.ResourceTag {
	tags := make([]egoscale.ResourceTag, 0, len(m))
	for key, value := range m {
		tags = append(tags, egoscale.ResourceTag{
			Key:   key,
			Value: value,
		})
	}
	return tags
}
