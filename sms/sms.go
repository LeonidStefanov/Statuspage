package sms

import (
	"io/ioutil"
	"log"
	"strings"

	"github.com/LeonidStefanov/Statuspage/models"
	"github.com/LeonidStefanov/Statuspage/utils"
)

type SMS struct {
	path string
}

func NewSmsReader(path string) *SMS {
	return &SMS{
		path: path,
	}
}

func (s *SMS) SMSData() (models.SMSArray, error) {

	var smsInfo models.SMSData
	file, err := ioutil.ReadFile(s.path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	lines := strings.Split(string(file), "\n")

	smsData := make([]models.SMSData, 0, len(lines))

	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], ";")

		if len(data) == 4 {

			smsInfo = models.SMSData{
				Country:      data[0],
				Bandwidth:    data[1],
				ResponseTime: data[2],
				Provider:     data[3],
			}
			smsData = append(smsData, smsInfo)

		}

	}

	return smsData, nil
}
func (s *SMS) SMSFilter(sms models.SMSArray) models.SMSArray {
	countries := utils.CountriesList()
	filterData := models.SMSArray{}

	for i := 0; i < len(sms); i++ {
		for j := 0; j < len(countries); j++ {
			if sms[i].Country == countries[j] && sms[i].Provider == utils.ProviderByCountry(countries[j]) {

				filterData = append(filterData, sms[i])
			}
		}

	}

	if len(filterData) != cap(filterData) {
		tmp := make(models.SMSArray, len(filterData))
		copy(tmp, filterData)
		filterData = tmp
	}

	return filterData
}
