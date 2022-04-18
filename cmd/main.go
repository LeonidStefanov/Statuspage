package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

// Cr := map[string]string{
// 	"RU": "Russian Federation",
// 	"CA": "Canada",
// 	"BG": "Great Britain",
// }

func main() {

	file, err := ioutil.ReadFile("sms.data")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(file))

	smsStr := strings.Split(string(file), ";")

	fmt.Println(smsStr)

}
