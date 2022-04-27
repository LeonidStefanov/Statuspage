package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LeonidStefanov/Statuspage/mms"
	"github.com/LeonidStefanov/Statuspage/sms"
)

func main() {

	// SMS Data

	smsReader := sms.NewSmsReader("./simulator/sms.data")

	data, err := smsReader.SMSData()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("SMS")

	data.Print()

	// MMS Data

	mmsReader := mms.NewMmsReader("http://127.0.0.1:8383", time.Second*15)

	mmsData, err := mmsReader.MMSData()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("MMS")

	filteredData := mmsReader.MMSFilter(mmsData)
	if err != nil {
		log.Println(err)
	}

	filteredData.Print()

}
