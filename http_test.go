package remote

import (
	"fmt"
	"testing"
	// "time"
)

type mfbResStruct struct {
	Errno   int         `json:"errno"`
	Errmsg  string      `json:"errmsg"`
	RetData interface{} `json:"retData"`
}

func Test_post(t *testing.T) {
	hc := GetDefaultClient()
	myreq := &RequestStruct{Uri: "http://localhost:8020/worker/task"}
	params := make(map[string]string)
	params["command"] = `{"name":"diff", "params":{"cc":"dd"}}`
	params["id"] = "123"
	request, err := myreq.PostRequest(params)
	if err != nil {
		t.Errorf("%s", "获取PostRequest Failed")
	}
	resp, err := hc.Do(request)
	if err != nil {
		t.Errorf("%s: %s", "请求失败", err.Error())
	} else {
		var retData mfbResStruct
		if err := resp.EchoBody.ToJson(&retData); err != nil {
			t.Errorf("%s: %s", "返回值格式不合法", err.Error())
		}
		fmt.Printf("%v", retData)
		if retData.Errno != 0 {
			t.Errorf("%d: %v", retData.Errno, retData)
		}
		t.Logf("test post request OK")
	}
}

// func Test_GetMethod(t *testing.T) {
// 	hc := GetDefaultClient()
// 	myreq := &RequestStruct{Uri: "http://cq01-testing-zfqa29.cq01.baidu.com:8087/pms/v1/appkey/appkey/pcode"}
// 	request, err := myreq.GetMethodRequest()
// 	if err != nil {
// 		t.Errorf("%s", "获取GetMethodRequest Failed")
// 	}
// 	resp, err := hc.Do(request)
// 	if err != nil {
// 		t.Errorf("%s: %s", "请求失败", err.Error())
// 	} else {
// 		retData := &mfbResStruct{}
// 		err = resp.EchoBody.ToJson(retData)
// 		if err != nil {
// 			t.Errorf("返回值格式不正确")
// 		}
// 		if retData.Errno != 0 {
// 			t.Errorf("返回值格式不正确: %v", retData)
// 		}
// 		t.Logf("%s", "GetMethodRequest test OK")
// 	}
// }

// func Test_post(t *testing.T) {
// 	hc := GetDefaultClient()
// 	myreq := &RequestStruct{Uri: "http://cq01-testing-zfqa29.cq01.baidu.com:8087/pms/v1/appkey/appkey/pcode"}
// 	params := make(map[string]string)
// 	params["pcodeName"] = time.Now().Format("20060102150405")
// 	params["versionName"] = "1.2.3"
// 	params["versionCode"] = "123"
// 	params["packageName"] = "com.baidu.test"
// 	request, err := myreq.PostRequest(params)
// 	if err != nil {
// 		t.Errorf("%s", "获取PostRequest Failed")
// 	}
// 	resp, err := hc.Do(request)
// 	if err != nil {
// 		t.Errorf("%s: %s", "请求失败", err.Error())
// 	} else {
// 		var retData mfbResStruct
// 		if err := resp.EchoBody.ToJson(&retData); err != nil {
// 			t.Errorf("%s: %s", "返回值格式不合法", err.Error())
// 		}
// 		if retData.Errno != 0 {
// 			t.Errorf("%d: %v", retData.Errno, retData)
// 		}
// 		t.Logf("test post request OK")
// 	}
// }

// func Test_postFile(t *testing.T) {
// 	hc := GetDefaultClient()
// 	myreq := &RequestStruct{Uri: "http://szwg-qatest-dpf038.szwg01.baidu.com:8061/apk/upload"}
// 	params := make(map[string]string)
// 	params["appid"] = "123"
// 	fpath := "D:\\test\\KirinDemo.apk"
// 	request, err := myreq.PostFileRequest(params, fpath, "upfile")
// 	if err != nil {
// 		t.Errorf("%s", "获取PostFileRequest失败")
// 	}
// 	res, err := hc.Do(request)
// 	echoStr, err := res.EchoBody.ToString()
// 	if err != nil {
// 		fmt.Println(echoStr)
// 		t.Errorf("%s: %s", "请求失败", err.Error())
// 	} else {
// 		t.Logf("%s", res)
// 	}
// }

// func Test_put(t *testing.T) {
// 	hc := GetDefaultClient()
// 	myreq := &RequestStruct{Uri: "http://cq01-testing-zfqa29.cq01.baidu.com:8087/pms/v1/appkey/appkey/pcode/pcodeName3/release"}
// 	request, err := myreq.PutMethodRequest(nil)
// 	if err != nil {
// 		t.Errorf("%s", "获取PostFileRequest失败")
// 	}
// 	res, err := hc.Do(request)
// 	echoStr, err := res.EchoBody.ToString()
// 	if err != nil {
// 		fmt.Println(echoStr)
// 		t.Errorf("%s: %s", "请求失败", err.Error())
// 	} else {
// 		t.Logf("%s", res)
// 	}
// }
