package api

func decodeCity(city string) (cityCode string) {

	switch city {
	case "МОСКВА":
		city = "2000000"
		return city

	case "САНКТ-ПЕТЕРБУРГ":
		city = "2004001"
		return city

	case "ОРСК":
		city = "2040480"
		return city
	}
	return
}
