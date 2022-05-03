package mms

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/LeonidStefanov/Statuspage/models"
	"github.com/LeonidStefanov/Statuspage/utils"
)

type MMS struct {
	url    string
	client *http.Client
}

func NewMmsReader(url string, timeOut time.Duration) *MMS {
	return &MMS{
		url: url,
		client: &http.Client{
			Timeout: timeOut,
		},
	}
}

func (m *MMS) MMSData() (models.MMSArray, error) {
	req, err := http.NewRequest("GET", m.url+"/mms", nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := m.client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		buf, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		mssData := models.MMSArray{}

		err = json.Unmarshal(buf, &mssData)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		return mssData, nil
	case http.StatusInternalServerError:
		return nil, fmt.Errorf("server error code %v", resp.StatusCode)
	default:
		return nil, fmt.Errorf("unexpected response code %v", resp.StatusCode)
	}

}

func (m *MMS) MMSFilter(mms models.MMSArray) models.MMSArray {
	countries := utils.CountriesList()
	filterData := []models.MMSData{}

	for i := 0; i < len(mms); i++ {
		for j := 0; j < len(countries); j++ {
			if mms[i].Country == countries[j] && mms[i].Provider == utils.ProviderByCountry(countries[j]) {

				filterData = append(filterData, mms[i])
			}
		}

	}

	if len(filterData) != cap(filterData) {
		tmp := make(models.MMSArray, len(filterData))
		copy(tmp, filterData)
		filterData = tmp
	}

	return filterData
}
