package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDarwinCPU(t *testing.T) {
	var darwin = GetDarwinReport()
	fmt.Printf("%v", darwin.Cpu)
	assert.NotEmpty(t, darwin.Cpu.Vendor.Name)
}
