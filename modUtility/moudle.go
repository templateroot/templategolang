package modUtility

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func Utility_Initialize() error {
	catchCrash()
	err := config_initialize()
	if err != nil {
		return err
	}

	urlLog := Config_Read(C_Key_LogUrl)
	if len(urlLog) < 10 {
		return errors.New("url of log too short")
	}
	tokenLog := Config_Read(C_Key_LogToken)
	err = log_initialize(urlLog, tokenLog)
	if err != nil {
		return err
	}
	utility_checkInstID()

	err = httpkv_Initialize()
	if err != nil {
		LogError("http kv utility initialize error")
		return err
	}

	UpdateTimeNew()

	return nil
}

func Utility_writeStartLog() error {
	strTimeNow := time.Now().Format("2006-01-02 15:04:05")

	strLog := fmt.Sprintf("[%s-%s] start, local time: %s", C_APPID, G_AppToken, strTimeNow)

	return LogInfo(strLog)
}

func utility_checkInstID() {
	G_AppToken = Config_Read(C_Key_AppToken) // this key defined in libgatlingconfig
	if G_AppToken == "" {
		G_AppToken = strconv.Itoa(int(time.Now().Unix()))
	}
}
