package main

type DarwinGpu struct {
	Name        string                 `json:"_name"`
	VramShared  string                 `json:"spdisplays_vram_shared"`
	Vram        string                 `json:"spdisplays_vram"`
	Model       string                 `json:"sppci_model"`
	X           map[string]interface{} `json:"-"` // Rest of the fields should go here.
}

type DarwinCpu struct {
	Model string                 `json:"cpu_type"`
	Cores int                    `json:"number_processors"`
	Ram   string                 `json:"physical_memory"`
	Speed string                 `json:"current_processor_speed"`
	X     map[string]interface{} `json:"-"` // Rest of the fields should go here.
}

type DarwinProfilerOutput struct {
	Gpus []DarwinGpu `json:"SPDisplaysDataType"`
	Cpus []DarwinCpu `json:"SPHardwareDataType"`
}
