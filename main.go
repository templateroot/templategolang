package main

import (
	"fmt"
	"wjgotemplate1/modHttpServer"
	"wjgotemplate1/modUtility"
)

func main() {

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
