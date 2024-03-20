package darwin

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

func TestValidMemInput(t *testing.T) {
	mem := GetDarwinMemory([]string {
	"BANK 0/DIMM0:",
	"",
    "Size: 8 GB",
    "Type: DDR4 SO-DIMM",
    "Speed: 2400 MHz",
    "Status: OK",
    "Manufacturer: 0x802C",
    "Part Number: 0x3841544631473634485A2D324733453220202020",
    "Serial Number: 0x25250504",
	"",
    "BANK 0/DIMM1:",
	"",
    "Size: 8 GB",
    "Type: DDR4 SO-DIMM",
    "Speed: 2400 MHz",
    "Status: OK",
    "Manufacturer: 0x802C",
    "Part Number: 0x3841544631473634485A2D324733453220202020",
    "Serial Number: 0x25250417",
	"",
	})

	assert.Len(t, mem, 2)
}
