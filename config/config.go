package config

import (
	"log"
	"os"
)

var (
	RancherEnvID     = "1a51"
	RancherStackID   = "1st16"
	RancherBaseUrl   = "http://10.2.253.121:8080/v2-beta/projects/"
	RancherBasicAuth = "Basic ODc2MkNBOUExMzM3NTYyNjIwODU6YzdBbWZrZHFSem9ONzE5Y29jeTFzdEw1cmk4YmloaXlVaVhqaEM1Vg=="
)

func InitEnvs() {

	if t := os.Getenv("RancherBaseUrl"); t != "" {
		RancherBaseUrl = t
	}
	if t := os.Getenv("RancherEnvID"); t != "" {
		RancherEnvID = t
	}
	if t := os.Getenv("RancherStackID"); t != "" {
		RancherStackID = t
	}
	if t := os.Getenv("RancherBasicAuth"); t != "" {
		RancherBasicAuth = t
	}

	log.Print("RancherBaseUrl:      ", RancherBaseUrl)
	log.Print("RancherEnvID:        ", RancherEnvID)
	log.Print("RancherStackID:      ", RancherStackID)
	log.Print("RancherBasicAuth:    ", RancherBasicAuth)
}
