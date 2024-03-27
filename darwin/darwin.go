package darwin

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"

	. "github.com/1okey/hw-report/types"
)

const cpu_prefix = "machdep.cpu."

var fields_map = map[string]string{
	"features": "features",
	"cores":    "core_count",
	"threads":  "thread_count",
	"vendor":   "vendor",
	"brand":    "brand_string",
}

var gpu_fields_map = map[string]string{
	"driver_version": "EFI Driver Version",
	"memory":         "VRAM (Total)",
	"vendor":         "Vendor",
	"brand":          "Chipset Model",
}

var mem_fields_map = map[string]string{
	"size":  "Size",
	"speed": "Speed",
	"type":  "Type",
	"brand": "Manufacturer",
}

func get_darwin_all_cpu_fields() ([]string, error) {
	cmd := exec.Command("sysctl", "-a")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return []string{}, err
	}

	var lines = make([]string, 0)
	var start = 0
	var cpu_prefix_len = len(cpu_prefix)
	for idx, ch := range output {
		if ch == byte('\n') {
			if string(output[start:start+cpu_prefix_len]) == cpu_prefix {
				lines = append(lines, string(output[start:idx]))
			}
			start = idx + 1
		}
	}

	return lines, nil
}

func filter_necessary_cpu_fields(cpu_lines []string) map[string]string {
	cpu_map := make(map[string]string)

	var cpu_prefix_len = len(cpu_prefix)
	for _, line := range cpu_lines {
		if sep_idx := strings.Index(line, ":"); sep_idx >= cpu_prefix_len {
			for k, v := range fields_map {
				if line[cpu_prefix_len:sep_idx] == v {
					cpu_map[k] = strings.Trim(line[sep_idx+1:], " ")
				}
			}
		}
	}

	return cpu_map
}

func GetDarwinCPU(cpu_fields []string) CPU {
	filtered_fields := filter_necessary_cpu_fields(cpu_fields)

	cores, _ := strconv.ParseInt(filtered_fields["cores"], 10, 8)

	return CPU{
		Vendor: Vendor{
			Name:  filtered_fields["vendor"],
			Model: filtered_fields["brand"],
		},
		Cores:    uint8(cores),
		Features: strings.Split(filtered_fields["features"], " "),
	}
}

func get_darwin_gpu_report() ([]string, error) {
	cmd := exec.Command("system_profiler", "SPDisplaysDataType")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return []string{}, err
	}

	var lines = make([]string, 0)
	var start = 0
	for idx, ch := range output {
		if ch == byte('\n') {
			lines = append(lines, strings.Trim(string(output[start:idx]), " "))
			start = idx + 1
		}
	}

	return lines, nil
}

func filter_necessary_gpu_fields(gpu_lines []string) map[string]string {
	gpu_map := make(map[string]string)
	for _, line := range gpu_lines {
		sep_idx := strings.Index(line, ":")

		if sep_idx >= 0 {
			for k, v := range gpu_fields_map {
				if line[:sep_idx] == v {
					gpu_map[k] = strings.Trim(line[sep_idx+1:], " ")
				}
			}
		}
	}

	return gpu_map
}

func GetDarwinGPU(gpu_fields []string) GPU {
	filtered_fields := filter_necessary_gpu_fields(gpu_fields)
	return GPU{
		Vendor: Vendor{
			Name:  filtered_fields["vendor"],
			Model: filtered_fields["brand"],
		},
		Memory:        filtered_fields["memory"],
		DriverVersion: filtered_fields["driver_version"],
	}
}

func get_darwin_mem_report() ([]string, error) {
	cmd := exec.Command("system_profiler", "SPMemoryDataType")
	output, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return []string{}, err
	}

	var lines = make([]string, 0)
	var start = 0
	for idx, ch := range output {
		if ch == byte('\n') {
			lines = append(lines, strings.Trim(string(output[start:idx]), " "))
			start = idx + 1
		}
	}

	return lines, nil
}

func parse_mem_from_lines(lines []string) RAM {
	mem_map := make(map[string]string)
	for _, line := range lines {
		sep_idx := strings.Index(line, ":")

		if sep_idx >= 0 {
			for k, v := range mem_fields_map {
				if line[:sep_idx] == v {
					mem_map[k] = strings.Trim(line[sep_idx+1:], " ")
				}
			}
		}
	}
	return RAM{
		Vendor: Vendor{
			Model: "-",
			Name:  mem_map["brand"],
		},
		Speed: mem_map["speed"],
		Size:  mem_map["size"],
		Type:  mem_map["type"],
	}
}

func GetDarwinMemory(mem_lines []string) []RAM {
	mem_banks := make([]RAM, 0)

	for idx := 0; idx < len(mem_lines); {
		if strings.Index(mem_lines[idx], "BANK") >= 0 && idx+9 < len(mem_lines) {
			mem_banks = append(mem_banks, parse_mem_from_lines(mem_lines[idx+2:idx+9]))
			idx += 9
		} else {
			idx += 1
		}
	}

	return mem_banks
}

func GetDarwinReport() HwReport {
	darwin_cpu_report, err := get_darwin_all_cpu_fields()
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	cpu := GetDarwinCPU(darwin_cpu_report)

	gpu_fields, err := get_darwin_gpu_report()
	if err != nil {
		log.Fatal(err.Error())
	}

	gpu := GetDarwinGPU(gpu_fields)

	mem_fields, err := get_darwin_mem_report()
	if err != nil {
		log.Fatal(err.Error())
	}
	mem := GetDarwinMemory(mem_fields)

	return HwReport{
		Cpu: cpu,
		Gpu: gpu,
		Ram: mem,
	}
}
