package model

import (
	"sync"
)

//OutputLog define log output
type OutputLog struct {
	Mu              sync.Mutex
	Date            string
	WeekDay         string
	ElementInfo     map[string]int
	UserInfo        map[string]int
	MinuteInfo      map[string]int
	QuoteCodeInfo   map[string]int
	SyubetuInfo     map[string]int
	KubunHassinInfo map[string]int
}

//KubunHassinInfo define log kubun-hasshin pair
type KubunHassinInfo struct {
	Kubun  string
	Hassin string
	Count  int
}

//InputLog define input log
type InputLog struct {
	LogType                 string `json:"LogType"`
	Level                   string `json:"Level"`
	Timestamp               string `json:"Timestamp"`
	Port                    string `json:"Port"`
	ReceiveHeader           string `json:"ReceiveHeader"`
	ResponseElement         string `json:"ResponseElement"`
	LogReplenishmentMessage string `json:"LogReplenishmentMessage"`
}
