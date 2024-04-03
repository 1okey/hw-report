package windows

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	. "github.com/1okey/hw-report/types"
)

var cpu_fields_map = map[string]string{
	"NumberOfLogicalProcessors": "threads",
	"NumberOfCores":             "cores",
	"Manufacturer":              "vendor",
	"Name":                      "brand",
}

var gpu_fields_map = map[string]string{
	"EFI Driver Version": "driver_version",
	"VRAM":               "memory",
	"Vendor":             "vendor",
	"Chipset Model":      "brand",
}

var mem_fields_map = map[string]string{
	"Size":         "size",
	"Speed":        "speed",
	"Type":         "type",
	"Manufacturer": "brand",
}

func get_win_cli_output_for_query(cli_command string, query ...string) ([]string, error) {
	cmd := exec.Command(cli_command, query...)
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return []string{}, err
	}

	var lines = make([]string, 0)
	var start = 0
	for idx, ch := range output {
		if ch == byte('\n') {
			line := strings.Trim(string(output[start:idx]), " ")

			if len(line) > 0 {
				lines = append(lines, line)
			}

			start = idx + 1
		}
	}

	return lines, nil
}

func filter_output_with_map(output []string, fields_map map[string]string) []map[string]string {
	sections := make([]map[string]string, 0)
	section := make(map[string]string, 0)

	for _, line := range output {
		parts := strings.Split(line, ":")

		if len(parts) != 2 {
			continue
		}

		key := strings.Trim(parts[0], " ")
		if key == "__GENUS" && len(section) > 0 {
			sections = append(sections, section)
			for k := range section {
				delete(section, k)
			}
		}

		if val, ok := fields_map[key]; ok {
			section[val] = strings.Trim(parts[1], " ")
		}

	}

	if len(section) > 0 {
		sections = append(sections, section)
	}

	return sections
}

func GetWinCPU(cpu_report map[string]string) CPU {
	cores, _ := strconv.ParseInt(cpu_report["cores"], 10, 8)

	return CPU{
		Vendor: Vendor{
			Name:  cpu_report["vendor"],
			Model: cpu_report["name"],
		},
		Cores:    uint8(cores),
		Features: []string{},
	}
}

func GetWinGPU(cpu_report map[string]string) GPU {
	return GPU{}
}

func GetWinMemory(cpu_report map[string]string) RAM {
	return RAM{}
}

func GetWinReport() HwReport {
	win_cpu_report, err := get_win_cli_output_for_query("gwmi", "-Class", "win32_processor", "-Property", "Name, Manufacturer, NumberOfCores, NumberOfLogicalProcessors")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	cpu := GetWinCPU(filter_output_with_map(win_cpu_report, cpu_fields_map)[0])

	gpu_fields, err := get_win_cli_output_for_query("gwmi", "-Class", "win32_videocontroller", "-Property", "Name, AdapterCompatibility, DriverVersion, AdapterRAM")
	if err != nil {
		log.Fatal(err.Error())
	}

	gpus := make([]GPU, 0)
	for _, section := range filter_output_with_map(gpu_fields, gpu_fields_map) {
		gpus = append(gpus, GetWinGPU(section))
	}

	mem_fields, err := get_win_cli_output_for_query("gwmi", "-Class", "win32_physicalmemory", "-Property", "Manufacturer, Capacity, Speed")

	if err != nil {
		log.Fatal(err.Error())
	}

	mem_banks := make([]RAM, 0)
	for _, section := range filter_output_with_map(mem_fields, mem_fields_map) {
		mem_banks = append(mem_banks, GetWinMemory(section))
	}

	return HwReport{
		Cpu: cpu,
		Gpu: gpus[0],
		Ram: mem_banks,
	}
}
