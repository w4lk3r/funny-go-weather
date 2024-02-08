package weather

type responseIPIP struct {
	Ret  string `json:"ret"`
	Data struct {
		IP       string   `json:"ip"`
		Location []string `json:"location"`
	} `json:"data"`
}

func GetMyIPLocation() (string, error) {
	response := new(responseIPIP)
	err := getJSON("https://myip.ipip.net/json", 3, response)
	if err != nil {
		return "", err
	}

	locationData := response.Data.Location

	province := locationData[1]
	city := locationData[2]
	return (province + "省" + city + "市"), nil
}
