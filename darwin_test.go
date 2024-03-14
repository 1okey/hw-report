package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidGpuInput(t *testing.T) {
	cpu := GetDarwinCPU([]string {
		"core_count: 4",
	})

	assert.Equal(t, uint8(0), cpu.Cores)
}

func TestValidGpuInput(t *testing.T) {
	cpu := GetDarwinCPU([]string {
		cpu_prefix + "core_count: 4",
	})

	assert.Equal(t, uint8(4), cpu.Cores)
}


func TestInvalidCpuInput(t *testing.T) {
	gpu := GetDarwinGPU([]string {
		"VRAMM: 8 GB",
	})

	assert.Equal(t, "", gpu.Memory)
}

func TestValidCpuInput(t *testing.T) {
	gpu := GetDarwinGPU([]string {
		"VRAM: 8 GB",
	})

	assert.Equal(t, "8 GB", gpu.Memory)
}
