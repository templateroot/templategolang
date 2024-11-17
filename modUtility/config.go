/*
	APPID: defined when this program start writing;
	AppToken: instance id; define in env or app config file, if not define(usually),  make a random one;

*/

package modUtility

import (
	"strconv"
	"time"

	"github.com/gatlinglab/libgatlingconfig"
)

const C_APPID = "WJGOTEMPLATE1"
const C_Key_AppToken = "APPTOKEN" /// define in library config too..;

const C_Key_LogUrl = "LOGURL"
const C_Key_LogToken = "LOGTOKEN"

var G_AppToken = ""
var G_HttpPort = 10000
var G_CurrentTime = ""

func GetSystemID() string {
	return ""
}

func config_initialize() error {
	err := libgatlingconfig.GetSingleGatlingConfig().Initialize(C_APPID)
	if err != nil {
		return err
	}

	strTmp := Config_Read("HTTPPORT")
	if strTmp != "" {
		iRet, err := strconv.Atoi(strTmp)
		if err == nil && iRet > 100 && iRet < 65535 {
			G_HttpPort = iRet
		}
	}

	return nil
}

func Config_Read(key string) string {
	return libgatlingconfig.GetSingleGatlingConfig().Get(key)
}

func UpdateTimeNew() {
	var cstZone = time.FixedZone("GMT", 8*3600) // 东八
	G_CurrentTime = time.Now().In(cstZone).Format("20060102+15:04:05")
}
