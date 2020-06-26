package main

import (
	"divsperf/command"
	"fmt"
	"runtime"
)

const (
	swname = "divsperf"
	vnumber = "0.0.0"
	lastupdate = "2020/6/22"
)

func main() {
	ShowVersion()
	command.Processing()
}

func ShowVersion() {
	fmt.Printf("%s-%s %s\n", swname, vnumber, lastupdate)
	fmt.Sprintf("%s/%s, %s\n", runtime.GOOS, runtime.GOARCH, runtime.Version())
}