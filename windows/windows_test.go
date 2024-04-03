package windows

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvalidCpuInput(t *testing.T) {
	cpu_data := filter_output_with_map([]string{
		"Availability                            : 3",
		"CpuStatus                               : 1",
		"CurrentVoltage                          : 7",
		"DeviceID                                : CPU0",
		"ErrorCleared                            :",
		"ErrorDescription                        :",
		"LastErrorCode                           :",
		"LoadPercentage                          : 3",
		"Status                                  : OK",
		"StatusInfo                              : 3",
		"AddressWidth                            : 64",
		"DataWidth                               : 64",
		"ExtClock                                : 100",
		"L2CacheSize                             : 2048",
		"L2CacheSpeed                            :",
		"MaxClockSpeed                           : 1190",
		"PowerManagementSupported                : False",
		"ProcessorType                           : 3",
		"Revision                                : 32261",
		"SocketDesignation                       : U3E1",
		"Version                                 :",
		"VoltageCaps                             :",
		"AssetTag                                : To Be Filled By O.E.M.",
		"Caption                                 : Intel64 Family 6 Model 126 Stepping 5",
		"Characteristics                         : 252",
		"ConfigManagerErrorCode                  :",
		"ConfigManagerUserConfig                 :",
		"CreationClassName                       : Win32_Processor",
		"CurrentClockSpeed                       : 991",
		"Description                             : Intel64 Family 6 Model 126 Stepping 5",
		"Family                                  : 205",
		"InstallDate                             :",
		"L3CacheSize                             : 6144",
		"L3CacheSpeed                            : 0",
		"Level                                   : 6",
		"NumberOfEnabledCore                     : 4",
		"OtherFamilyDescription                  :",
		"PartNumber                              : To Be Filled By O.E.M.",
		"PNPDeviceID                             :",
		"PowerManagementCapabilities             :",
		"Role                                    : CPU",
		"SecondLevelAddressTranslationExtensions : False",
		"SerialNumber                            : To Be Filled By O.E.M.",
		"Stepping                                :",
		"SystemCreationClassName                 : Win32_ComputerSystem",
		"ThreadCount                             : 8",
		"UniqueId                                :",
		"UpgradeMethod                           : 1",
		"VirtualizationFirmwareEnabled           : False",
		"VMMonitorModeExtensions                 : False",
	}, cpu_fields_map)

	assert.Len(t, cpu_data, 0)
}

func TestValidCpuInput(t *testing.T) {
	cpu_data := filter_output_with_map([]string{
		"Manufacturer                            : GenuineIntel",
		"Name                                    : Intel(R) Core(TM) i5-1035G1 CPU @ 1.00GHz",
		"NumberOfCores                           : 4",
		"NumberOfEnabledCore                     : 4",
		"NumberOfLogicalProcessors               : 8",
		"ThreadCount                             : 8",
	}, cpu_fields_map)

	assert.Len(t, cpu_data, 1)
	assert.Len(t, cpu_data[0], 4)

	cpu := GetWinCPU(cpu_data[0])

	assert.Equal(t, uint8(4), cpu.Cores)
	assert.Equal(t, "GenuineIntel", cpu.Vendor.Name)
}
