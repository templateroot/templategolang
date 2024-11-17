package modUtility

import (
	"bytes"
	"io"
	"net/http"
)

const C_Key_DBListUrl = "MAINKVURL"
const C_Key_DBListToken = "MAINKVTOKEN"
const C_Key_HttpPort = "HTTPPORT"
const C_Key_MessageUrl = "MESSAGEURL"
const C_Key_MessageToken = "MESSAGETOKEN"

var G_MainKVUrl string = ""
var G_MainKVToken string = ""
var G_MessageUrl string = ""
var G_MessageToken string = ""

func httpkv_Initialize() error {
	G_MainKVUrl = Config_Read(C_Key_DBListUrl)
	if G_MainKVUrl != "" && G_MainKVUrl[len(G_MainKVUrl)-1] != '/' {
		G_MainKVUrl += "/"
	}
	G_MainKVToken = Config_Read(C_Key_DBListToken)
	G_MessageUrl = Config_Read(C_Key_MessageUrl)
	if G_MessageUrl != "" && G_MessageUrl[len(G_MessageUrl)-1] != '/' {
		G_MessageUrl += "/"
	}
	G_MessageToken = Config_Read(C_Key_MessageToken)

	return nil
}

func HttpKVGet(id string) string {
	url1 := G_MainKVUrl + id
	req, err := http.NewRequest(http.MethodGet, url1, nil)
	if err != nil {
		return ""
	}

	req.Header.Set("user-agent", "WJBK 1.0")
	req.Header.Set("X-API-KEY", G_MainKVToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		//fmt.Println("success")
	} else {
		//fmt.Println("failed: ", resp)
		return ""
	}
	data1, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(data1)
}

// ture or false
func HttpKVSet(key, value string) bool {
	url1 := G_MainKVUrl + key
	req, err := http.NewRequest(http.MethodPost, url1, bytes.NewBuffer([]byte(value)))
	if err != nil {
		return false
	}

	req.Header.Set("user-agent", "WJBK 1.0")
	req.Header.Set("X-API-KEY", G_MainKVToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		//fmt.Println("success")
	} else {
		//fmt.Println("failed: ", resp)
		return false
	}
	data1, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	if string(data1) == "OK" {
		return true
	}
	return false
}
func HttpSystemPostMessage(key, value string) string {
	url1 := G_MessageUrl + key
	req, err := http.NewRequest(http.MethodPost, url1, bytes.NewBuffer([]byte(value)))
	if err != nil {
		return ""
	}

	req.Header.Set("user-agent", "WJBK 1.0")
	req.Header.Set("X-API-KEY", G_MessageToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		//fmt.Println("success")
	} else {
		//fmt.Println("failed: ", resp)
		return ""
	}
	data1, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	return string(data1)
}
