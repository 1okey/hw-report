package main

import (
	"encoding/json"
	"log"
	"os/exec"
	"runtime"
	"strings"
)

const os string = runtime.GOOS

func main() {

	switch {
	case strings.Index(MAC, os) > -1:
		gpuCmd := exec.Command("system_profiler", "SPDisplaysDataType", "SPHardwareDataType", "-json")
		out, err := gpuCmd.Output()
		if err != nil {
			log.Fatal(err.Error())
		}

		var decodedOutput DarwinProfilerOutput
		err = json.Unmarshal(out, &decodedOutput)
		if err != nil {
			log.Fatal(err.Error())
		}

		for _, gpu := range decodedOutput.Gpus {
			log.Print(gpu)
		}

		for _, cpu := range decodedOutput.Cpus {
			log.Print(cpu)
		}

	case strings.Index(WIN32, os) > -1 || strings.Index(WIN64, os) > -1:
		//gpuCmd := exec.Command(
		//	"nvidia-smi",
		//	"--query-gpu=index,uuid,driver_version,name,gpu_serial,display_active,display_mode,utilization.gpu,utilization.memory,memory.total,memory.used,memory.free,temperature.memory,temperature.gpu",
		//	"--format=csv",
		//	"--nounits")
		//out, err := gpuCmd.Output()
		//if err != nil {
		//	log.Fatal(err.Error())
		//}
		//
		//var decodedOutput DarwinProfilerOutput
		//err = json.Unmarshal(out, &decodedOutput)
		//if err != nil {
		//	log.Fatal(err.Error())
		//}
	}
}
