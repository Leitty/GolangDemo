package eureka

import (
	"Gin/learnGin/golangDemo/pkg/logging"
	"Gin/learnGin/golangDemo/pkg/setting"
	"bytes"
	"encoding/json"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	"strconv"
	"time"
)

type EurekaInstance struct {
	Instance Eureka `json:"instance"`
}

type Eureka struct {
	Hostname string `json:"hostName"`
	App string `json:"app"`
	IpAddr string `json:"ipAddr"`
	VipAddress string `json:"vipAddress"`
	Status string `json:"status"`
	Port int `json:"port"`
	SecurePort string `json:"securePort"`
	HomePageUrl string `json:"homePageUrl"`
	StatusPageUrl string `json:"statusPageUrl"`
	HealthCheckUrl string `json:"healthCheckUrl"`
	DataCenterInfo DataCenterInfo `json:"dataCenterInfo"`
	Metadata  Metadata  `json:"metadata"`
}

type DataCenterInfo struct {
	Name string `json:"name"`
}

type Metadata struct {
	InstanceId string `json:"instanceId"`
}

var instanceId string
var eurekaHealthUrl string

func Register() {
	//获取eureka的参数
	app := setting.EurekaSetting.AppName
	u1,_ := uuid.NewV4()
	instanceId = app + "-" + u1.String()
	port := setting.ServerSetting.HttpPort

	ipAddress,err := GetLocalIP()
	if err !=nil {
		ipAddress = "localhost"
	}
	homePageUrl := "http://" + ipAddress + ":" + strconv.Itoa(port)
	statusPageUrl := homePageUrl + "/" + setting.EurekaSetting.StatusUrl
	healthPageUrl := homePageUrl + "/" + setting.EurekaSetting.HealthUrl
	dataCenter := setting.EurekaSetting.DataCenterInfo
	securePort := setting.EurekaSetting.SecurePort

	eurekaHealthUrl = setting.EurekaHomeUrl + "/" + ipAddress+":" + instanceId

	eka := Eureka{
		Hostname:ipAddress,
		App:app,
		IpAddr:ipAddress,
		VipAddress:app,
		Status:"UP",
		Port:port,
		SecurePort: securePort,
		HomePageUrl: homePageUrl,
		StatusPageUrl: statusPageUrl,
		HealthCheckUrl: healthPageUrl,
		DataCenterInfo: DataCenterInfo{
			Name: dataCenter,
		},
		Metadata: Metadata{
			InstanceId:instanceId,
		},
	}
	instance := EurekaInstance{eka}

	registerData, err := json.Marshal(instance)
	if err != nil {
		log.Fatalf("JSON marshal failed: %s", err)
	}
	var result bool
	for {
		req := httpRequest(registerData,"POST", setting.EurekaHomeUrl)
		result = httpDo(req)
		if result {
			break
		} else {
			time.Sleep(time.Second * 5)
		}
	}
}

func StartHeartbeat() {
	for {
		time.Sleep(time.Second * 30)
		heartbeat(eurekaHealthUrl)
	}
}

func heartbeat(url string) {
	req := httpRequest(nil, "PUT", url)
	httpDo(req)
}


// 组装request
func httpRequest(byte []byte, method, url string) *http.Request{
	reader := bytes.NewBuffer(byte)
	req, err := http.NewRequest(method, url, reader)
	if err != nil {
		logging.Fatalf("Fail to translate data to NewRequest: %s", err)
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	return req
}

func httpDo(req *http.Request) bool{
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		logging.Warn("HTTP request failed: %s", err)
		return false
	}

	return true
}


func DeRegister(){
	req := httpRequest(nil, "DELETE", eurekaHealthUrl)
	httpDo(req)
}