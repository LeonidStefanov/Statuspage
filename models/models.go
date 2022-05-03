package models

import "fmt"

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type VoiceCallData struct {
	Country             string
	Bandwidth           string
	ResponseTime        string
	Provider            string
	ConnectionStability float32
	TTFB                int
	VoicePurity         int
	MedianOfCallsTime   int
}

type SMSArray []SMSData
type MMSArray []MMSData
type VoiceCallArray []VoiceCallData

func (m MMSArray) Print() {
	for i := 0; i < len(m); i++ {
		fmt.Println(m[i])
	}
}

func (s SMSArray) Print() {
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func (v VoiceCallArray) Print() {
	for i := 0; i < len(v); i++ {
		fmt.Println(v[i])
	}
}
