package types

import (
	"log"
	"os"
	"strconv"
	"strings"

	. "github.com/1okey/hw-report/markdown"
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
	Memory        string
	DriverVersion string
}

type RAM struct {
	Vendor
	Size  string
	Speed string
	Type  string
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

type HwReport struct {
	Cpu CPU
	Gpu GPU
	Ram []RAM
}

func getReportMarkdown(hw *HwReport) *Markdown {
	md :=  Markdown{}

	cpu := []Serializable{
		Heading{Text: Paragraph{TextElements: []Text{{Content: "CPU", Style: Normal}}}, Size: uint8(3)},
		Paragraph{TextElements: []Text{
			{Content: "Vendor:", Style: Bold},
			{Content: hw.Cpu.Vendor.Name, Style: Normal},
		}},
		Paragraph{TextElements: []Text{
			{Content: "Name:", Style: Bold},
			{Content: hw.Cpu.Vendor.Model, Style: Normal},
		}},
		Paragraph{TextElements: []Text{
			{Content: "Cores:", Style: Bold},
			{Content: strconv.Itoa(int(hw.Cpu.Cores)), Style: Normal},
		}},
		Paragraph{TextElements: []Text{
			{Content: "Features:", Style: Bold},
			{Content: "[" + strings.Join(hw.Cpu.Features, ", ") + "]", Style: Normal},
		}},
	}	
	md.Add(cpu...)	

	gpu := []Serializable{
		Heading{Text: Paragraph{TextElements: []Text{{Content: "GPU", Style: Normal}}}, Size: uint8(3)},
		Paragraph{TextElements: []Text{
			{Content: "Vendor:", Style: Bold},
			{Content: hw.Gpu.Vendor.Name, Style: Normal},
		}},
		Paragraph{TextElements: []Text{
			{Content: "Name:", Style: Bold},
			{Content: hw.Gpu.Vendor.Model, Style: Normal},
		}},
		Paragraph{TextElements: []Text{
			{Content: "Memory:", Style: Bold},
			{Content: hw.Gpu.Memory, Style: Normal},
		}},
		Paragraph{TextElements: []Text{
			{Content: "Driver Version:", Style: Bold},
			{Content: hw.Gpu.DriverVersion, Style: Normal},
		}},
	}	
	md.Add(gpu...)	


	md.Add(Heading{Text: Paragraph{TextElements: []Text{{Content: "RAM", Style: Normal}}}, Size: uint8(3)})

	for _, mem := range hw.Ram {
		md.Add([]Serializable{
			Heading{Text: Paragraph{TextElements: []Text{
				{Content: "Name:", Style: Bold},
				{Content: mem.Vendor.Model, Style: Normal},
			}}, Size: uint8(5)},
			Paragraph{TextElements: []Text{
				{Content: "Vendor:", Style: Bold},
				{Content: mem.Vendor.Name, Style: Normal},
			}},
			Paragraph{TextElements: []Text{
				{Content: "Size:", Style: Bold},
				{Content: mem.Size, Style: Normal},
			}},
			Paragraph{TextElements: []Text{
				{Content: "Speed:", Style: Bold},
				{Content: mem.Speed, Style: Normal},
			}},
		}...)	

	}

	return &md
}


func (hw *HwReport) Print() {
	md := getReportMarkdown(hw)
	_, err := os.Stdout.Write(md.Bytes())

	if err != nil {
		log.Panic("Failed saving report to stdout")
	}
}

func (hw *HwReport) Save() {
	file, err := os.Create("hwreport.md")
	if err != nil {
		log.Panic("Failed to create file for report")
	}

	defer file.Close();

	md := getReportMarkdown(hw)
	_, err = file.Write(md.Bytes())

	if err != nil {
		file.Close()
		log.Panic("Failed saving report into markdown file")
	}
}
