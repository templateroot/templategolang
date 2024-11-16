package main

import (
	"fmt"
	"os"
	"runtime/debug"
	"template1/modHttpServer"
	"template1/modUtility"
	"time"
)

func main() {

	catchCrash()
	err := modUtility.Utility_Initialize()
	if err != nil {
		fmt.Println("utility initialize failed: ", err)
		return
	}

	modUtility.Utility_writeStartLog()

	err = modHttpServer.Http_Initialize()
	if err != nil {
		modUtility.LogError("http init failed: " + err.Error())
		return
	}

	err = modHttpServer.Http_Start()
	if err != nil {
		modUtility.LogError("http start failed: " + err.Error())
		return
	}

}

func catchCrash() {
	f, err := os.Create(time.Now().Format("crash20060102150405.log"))
	if err != nil {
		panic(err)
	}
	debug.SetCrashOutput(f, debug.CrashOptions{})
}
