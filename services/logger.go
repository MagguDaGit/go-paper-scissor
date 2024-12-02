package services

import (
	"log"
	"rubrumcreation.com/go-paper-scissor/models"
	"rubrumcreation.com/go-paper-scissor/util"
	"time"
)

func LogApiInfo(route string, address string, info string, executionTime int64) {
	duration := util.DurationToString(time.Duration(executionTime))
	logEntry := &models.ApiInfoLog{DateTime: util.DateTimeNowString(), LogType: models.INFO, Address: address, Message: info, Route: route, ExecutionTime: duration}
	log.Default().Println(logEntry.ToJson())
}
