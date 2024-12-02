package models

import (
	"encoding/json"
	"log"
)

type logType string

const (
	INFO  logType = "INFO"
	warn  logType = "WARN"
	error logType = "ERROR"
)

type Log interface {
	ToJson() string
}

type ApiInfoLog struct {
	DateTime      string
	LogType       logType
	Address       string
	Message       string
	Route         string
	ExecutionTime string
}

func (l ApiInfoLog) ToJson() string {
	json, err := json.Marshal(l)
	if err != nil {
		log.Fatal(err)
	}
	return string(json)
}
