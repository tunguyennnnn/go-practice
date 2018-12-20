package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func checkVersion() {
	goVersion := runtime.Version()
	major := strings.Split(goVersion, ".")[0][2]
	minor := strings.Split(goVersion, ".")[1]
	m1, _ := strconv.Atoi(string(major))
	m2, _ := strconv.Atoi(string(minor))
	if m1 == 1 && m2 < 12 {
		fmt.Println("Need version greater than 1.8")
		panic("Invalid version")
	}
}

func main() {
	fmt.Println("Youre using ", runtime.Compiler)
	fmt.Println("on a ", runtime.GOARCH, "machine")
	fmt.Println("using Go version ", runtime.Version())
	fmt.Println("Number of CPUs ", runtime.NumCPU())
	fmt.Println("Number of Go Routine", runtime.NumGoroutine())
	checkVersion()
}
