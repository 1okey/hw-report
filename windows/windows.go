package windows

import (
	"fmt"
	"log"
	"os/exec"
	"strings"

	. "github.com/1okey/hw-report/types"
)


func get_win_cli_output_for_query(cli_command string, query... string) ([]string, error) {
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
			lines = append(lines, strings.Trim(string(output[start:idx]), " "))
			start = idx + 1
		}
	}

	return lines, nil
}

func GetWinCPU(cpu_report []string) CPU {
	return CPU{}
}

func GetWinGPU(cpu_report []string) GPU {
	return GPU{}
}

func GetWinMemory(cpu_report []string) []RAM {
	return []RAM{}
}

func GetWinReport() HwReport {
	win_cpu_report, err := get_win_cli_output_for_query("gwmi", "-Class", "win32_processor", "-Property", "TODO")
	if err != nil {
		log.Fatal(err.Error())
		panic(err)
	}

	cpu := GetWinCPU(win_cpu_report)

	gpu_fields, err := get_win_cli_output_for_query("gwmi", "-Class", "win32_videocontroller", "-Property", "TODO")
	if err != nil {
		log.Fatal(err.Error())
	}

	gpu := GetWinGPU(gpu_fields)

	mem_fields, err := get_win_cli_output_for_query("gwmi", "-Class", "win32_physicalmemory", "-Property", "TODO")

	if err != nil {
		log.Fatal(err.Error())
	}
	mem := GetWinMemory(mem_fields)

	return HwReport{
		Cpu: cpu,
		Gpu: gpu,
		Ram: mem,
	}
}
