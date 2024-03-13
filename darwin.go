package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

const cpu_prefix = "machdep.cpu."

var fields_map = map[string]string{
	"features": "features",
	"cores": "core_count",
	"threads": "thread_count",
	"vendor": "vendor",
	"brand": "brand_string",
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
			if string(output[start:start + cpu_prefix_len]) == cpu_prefix {
				lines = append(lines, string(output[start:idx]))
			}
			start = idx+1
		}
	}

	return lines, nil
 }

func filter_necessary_cpu_fields(cpu_lines []string) map[string]string {
	cpu_map := make(map[string]string)

	var cpu_prefix_len = len(cpu_prefix)
	for _, line := range cpu_lines {
		sep_idx := strings.Index(line, ":")

		for k, v := range fields_map {
			if line[cpu_prefix_len:sep_idx] == v {
				cpu_map[k] = strings.Trim(line[sep_idx + 1:], " ")
			}
		}
	}

	return cpu_map
}

func GetDarwinReport() HwReport {
	darwin_cpu_report, err := get_darwin_all_cpu_fields()
	if err != nil { 
		log.Fatal(err.Error())
		panic(err)
	}

	filtered_fields := filter_necessary_cpu_fields(darwin_cpu_report)
	cores, err := strconv.ParseInt(filtered_fields["cores"], 10, 8)

	return HwReport{
		Cpu: CPU{
			Vendor: Vendor{
				Name:  filtered_fields["vendor"],
				Model: filtered_fields["brand"],
			},
			Cores: uint8(cores),
			Features: strings.Split(filtered_fields["features"], " "),
		},
	}
}
