package modUtility

import (
	"os"
	"runtime/debug"
	"time"
)

func catchCrash() {
	f, err := os.Create(time.Now().Format("crash20060102150405.log"))
	if err != nil {
		panic(err)
	}
	debug.SetCrashOutput(f, debug.CrashOptions{})
}
