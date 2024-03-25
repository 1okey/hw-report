package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"
	"slices"
	"strings"

	. "github.com/1okey/hw-report/darwin"
	. "github.com/1okey/hw-report/linux"
	. "github.com/1okey/hw-report/types"
	. "github.com/1okey/hw-report/windows"
)

func make_report() (HwReport, error) {
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

func save_report(report HwReport, format string) {
	log.Panic("save_report not implemented")
}

func print_report(report HwReport) {
	fmt.Printf("%+v", report)
}

var SUPPORTED_FORMATS []string = []string{"markdown"}

func main() {
	args := os.Args[1:]
	format_arg_idx := -1
	for idx, arg := range args {
		if strings.Index(arg, "format") >= 0 {
			format_arg_idx = idx
			break
		}
	}

	format := ""
	if format_arg_idx >= 0 {
		format_arg := strings.Split(args[format_arg_idx], "=")
		if len(format_arg) < 2 || slices.Index(SUPPORTED_FORMATS, format_arg[1]) < 0 {
			log.Fatal("format type must be specified in the following form `format=type` where type is either `markdown` or `json`")
			os.Exit(1)
		}
		format = format_arg[1]
	}

	var report, err = make_report()
	if err != nil {
		log.Panic("Something strange occurred")
	}

	if len(format) > 0 {
		save_report(report, format)
	} else {
		print_report(report)
	}
}
