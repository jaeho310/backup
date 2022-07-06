package service

import (
	"encoding/json"
	"errors"
	"log"
	"strings"
)

type LogProcessor struct{}

func (LogProcessor) ExecuteLogProcessor(logByteData []byte) (string, error) {
	logMessage, err := getZLoggerMessage(logByteData)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return logMessage, nil
}

//func processLogMessage([]string) {
//
//}

func getZLoggerMessage(logByteData []byte) (string, error) {
	logMap := make(map[string]string)
	err := json.Unmarshal(logByteData, &logMap)
	if err != nil {
		log.Println(err)
		return "", err
	}
	for k, v := range logMap {
		if k == "log" {
			i := strings.Index(v, "[ZzL]")
			if i > -1 {
				return v[i+5:], nil
			} else {
				return "", errors.New("not zlogger")
			}
		}
	}
	return "", errors.New("fail")
}
