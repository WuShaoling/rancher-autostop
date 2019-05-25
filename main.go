package main

import (
	"encoding/json"
	"github.com/rancher-delete/config"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"time"
)

func main() {
	config.InitEnvs()
	t := time.NewTicker(120 * time.Second)
	for {
		select {
		case <-t.C:
			doStop()
		}
	}
}

func checkDate(d string) bool {
	stamp, err := time.ParseInLocation("2006-01-02T15:04:05Z", d, time.Local)
	if err != nil {
		return false
	}
	if time.Now().Unix()-stamp.Unix() >= 7200 {
		return true
	}
	return false
}

func doStop() {
	res, err := getAllCloudwares()
	if err == nil {
		for _, v := range res {
			tt := v.(map[string]interface{})
			if checkDate(tt["created"].(string)) && tt["state"] == "active" {
				stopContainer(tt["id"].(string))
			}
		}
	}
}

func stopContainer(sid string) {
	req, err := http.NewRequest(http.MethodDelete,
		config.RancherBaseUrl+path.Join(config.RancherEnvID, "services", sid),
		nil)
	if err != nil {
		log.Print("stopContainer http.NewRequest: ", err)
		return
	} else {
		log.Print("--->11")
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("authorization", config.RancherBasicAuth)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("stopContainer http.DefaultClient.Do: ", err)
		return
	} else {
		log.Print("--->12")
	}

	if resp.StatusCode != 200 {
		log.Print("停止失败: ", err)
	}
}

func getAllCloudwares() ([]interface{}, error) {

	req, err := http.NewRequest(http.MethodGet,
		config.RancherBaseUrl+path.Join(config.RancherEnvID, "stacks", config.RancherStackID, "services"), nil)
	if err != nil {
		log.Print("getAllCloudwares http.NewRequest: ", err)
		return nil, err
	} else {
		log.Print("--->1")
	}

	req.Header.Add("authorization", config.RancherBasicAuth)
	req.Header.Add("accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("getAllCloudwares http.DefaultClient.Do: ", err)
		return nil, err
	} else {
		log.Print("--->2")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("getAllCloudwares ioutil.ReadAll: ", err)
		return nil, err
	} else {
		log.Print("--->3")
	}

	var temp map[string]interface{}
	if err := json.Unmarshal([]byte(body), &temp); err != nil {
		log.Print("getAllCloudwares json.Unmarshal", err)
		return nil, err
	} else {
		log.Print("--->4")
	}

	return temp["data"].([]interface{}), nil
}
