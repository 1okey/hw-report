package main

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
