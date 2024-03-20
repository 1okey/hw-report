package main

import (
	"errors"
	"fmt"
	"runtime"

	. "github.com/1okey/hw-report/types"
	. "github.com/1okey/hw-report/darwin"
	. "github.com/1okey/hw-report/linux"
	. "github.com/1okey/hw-report/windows"
)

func get_report() (HwReport, error) {
	switch runtime.GOOS {
	case "windows":
		return GetWinReport(), nil
	case "linux":
		return GetLinuxReport(), nil
	case "darwin":
		return GetDarwinReport(), nil
	default:
		return HwReport{}, errors.New("unsupported system")
	}
}

func main() {
	var report, err = get_report()
	
	if err != nil {
		panic("Something strange occurred")
	}

	fmt.Print(report)
}
