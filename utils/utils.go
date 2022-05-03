package utils

func CountriesList() []string {
	return []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
}
func ProviderByCountry(country string) string {
	ProviderMap := map[string]string{
		"RU": "Topolo",
		"US": "Rond",
		"GB": "Topolo",
		"FR": "Topolo",
		"BL": "Kildy",
		"AT": "Topolo",
		"BG": "Rond",
		"DK": "Topolo",
		"CA": "Rond",
		"ES": "Topolo",
		"CH": "Topolo",
		"TR": "Rond",
		"PE": "Topolo",
		"NZ": "Kildy",
		"MC": "Kildy",
	}

	return ProviderMap[country]
}

func VoiceCallProviderByCountry(country string) string {
	voiceProviderMap := map[string]string{
		"RU": "TransparentCalls",
		"US": "E-Voice",
		"GB": "TransparentCalls",
		"FR": "TransparentCalls",
		"BL": "E-Voice",
		"AT": "TransparentCalls",
		"BG": "E-Voice",
		"DK": "JustPhone",
		"CA": "JustPhone",
		"ES": "E-Voice",
		"CH": "JustPhone",
		"TR": "TransparentCalls",
		"PE": "JustPhone",
		"NZ": "JustPhone",
		"MC": "E-Voice",
	}

	return voiceProviderMap[country]
}
