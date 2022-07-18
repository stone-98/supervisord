//go:build windows
// +build windows

package main_

func Daemonize(logfile string, proc func()) {
	proc()
}
