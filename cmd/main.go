package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LeonidStefanov/Statuspage/mms"
	"github.com/LeonidStefanov/Statuspage/sms"
)

func main() {

	// // SMS Data

	smsReader := sms.NewSmsReader("./simulator/sms.data")

	smsData, err := smsReader.SMSData()
	if err != nil {
		log.Println(err)
		return
	}

	filterSMS := smsReader.SMSFilter(smsData)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("SMS:")

	filterSMS.Print()

	// // MMS Data

	mmsReader := mms.NewMmsReader("http://127.0.0.1:8383", time.Second*15)

	mmsData, err := mmsReader.MMSData()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("MMS:")

	filterMMS := mmsReader.MMSFilter(mmsData)
	if err != nil {
		log.Println(err)
	}

	filterMMS.Print()

	// Voice Call Data

}
