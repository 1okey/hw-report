package main

import (
	"errors"
	"fmt"
	"runtime"
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
