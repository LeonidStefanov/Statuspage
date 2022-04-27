package utils

func GetCountriesList() []string {
	return []string{"RU", "US", "GB", "FR", "BL", "AT", "BG", "DK", "CA", "ES", "CH", "TR", "PE", "NZ", "MC"}
}
func GetProviderByCountry(country string) string {
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
