//go:build !windows
// +build !windows

package main_

import (
	"github.com/ochinchina/go-reaper"
)

// ReapZombie reap the zombie child process
func ReapZombie() {
	go reaper.Reap()
}
