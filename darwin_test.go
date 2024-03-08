package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDarwinCPU(t *testing.T) {
	var darwin = get_darwin_report()
	var darwin_cpu CPU = darwin.GetCPU()
	assert.Equal(t, darwin_cpu.Vendor.Name, "apple")
	assert.Equal(t, darwin_cpu.Cores, uint8(8))
}
