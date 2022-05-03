package voicecall

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/LeonidStefanov/Statuspage/models"
	"github.com/LeonidStefanov/Statuspage/utils"
)

type VoiceCall struct {
	path string
}

func NewVoiceCallReader(path string) *VoiceCall {
	return &VoiceCall{
		path: path,
	}
}

func (v *VoiceCall) VoiceCallData() (models.VoiceCallArray, error) {

	file, err := ioutil.ReadFile(v.path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	lines := strings.Split(string(file), "\n")

	voiceData := make(models.VoiceCallArray, 0, len(lines))

	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], ";")

		if len(data) == 8 {
			voiceDataItem, err := parse(data)
			if err != nil {
				return nil, err
			}
			voiceData = append(voiceData, *voiceDataItem)
		}

	}

	return voiceData, nil
}

func (v *VoiceCall) VoiceCallFilter(voice models.VoiceCallArray) models.VoiceCallArray {
	countries := utils.CountriesList()
	filterData := models.VoiceCallArray{}

	for i := 0; i < len(voice); i++ {
		for j := 0; j < len(countries); j++ {
			if voice[i].Country == countries[j] && voice[i].Provider == utils.ProviderByCountry(countries[j]) {

				filterData = append(filterData, voice[i])
			}
		}

	}

	if len(filterData) != cap(filterData) {
		tmp := make(models.VoiceCallArray, len(filterData))
		copy(tmp, filterData)
		filterData = tmp
	}

	return filterData
}

func parse(data []string) (*models.VoiceCallData, error) {
	var voiceCallInfo models.VoiceCallData

	dig1, err := strconv.ParseFloat(data[4], 32)
	if err != nil {
		return nil, err
	}

	dig2, err := strconv.Atoi(data[5])
	if err != nil {
		return nil, err
	}

	dig3, err := strconv.Atoi(data[6])
	if err != nil {
		return nil, err
	}

	dig4, err := strconv.Atoi(data[7])
	if err != nil {
		return nil, err
	}

	voiceCallInfo = models.VoiceCallData{
		Country:             data[0],
		Bandwidth:           data[1],
		ResponseTime:        data[2],
		Provider:            data[3],
		ConnectionStability: float32(dig1),
		TTFB:                dig2,
		VoicePurity:         dig3,
		MedianOfCallsTime:   dig4,
	}

	return &voiceCallInfo, nil
}
