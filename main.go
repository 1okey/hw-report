package main

import (
	"errors"
	"fmt"
	"runtime"
)

type Vendor struct {
	Name  string
	Model string
}

type CPU struct {
	Vendor
	Cores    uint8
	Features []string
}

type GPU struct {
	Vendor
	Features []string
}

type RAM struct {
	Vendor
	Size uint64
}

type Storage struct {
	Vendor
	Type string
	Size uint64
}

type OSInfo interface {
	GetMarkdown()
	GetCPU() CPU
	GetGPU() GPU
	GetStorage() []Storage
	GetRAM() RAM
}

type HwReport struct{
	Cpu CPU
}

func (report HwReport) GetCPU() CPU {
	return CPU{
		Vendor: Vendor{
			Name:  "apple",
			Model: "m1",
		},
		Cores:    8,
		Features: []string{},
	}
}

func get_ms_report() HwReport {
	return HwReport{}
}

func get_linux_report() HwReport {
	return HwReport{}
}

func get_darwin_report() HwReport {
	return HwReport{
		Cpu: CPU{
			Vendor: Vendor{
				Name:  "apple",
				Model: "m1",
			},
			Cores:    8,
			Features: []string{},
		},
	}
}

func get_report() (HwReport, error) {
	switch runtime.GOOS {
	case "windows":
		return get_ms_report(), nil
	case "linux":
		return get_linux_report(), nil

	case "darwin":
		return get_darwin_report(), nil
	default:
		return HwReport{}, errors.New("unsupported system")
	}
}

func main() {
	var report, err = get_report()
	
	if err != nil {
		panic("Something strange occured")
	}

	fmt.Print(report)
}
