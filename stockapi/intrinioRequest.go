package stockapi

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"encoding/json"
)

type Stock struct {
	MarketCap float32 `json:"marketcap"`
	Ticker string `json:"ticker"`
	Group string `json:"industry_group"`
}

type IndustryGroup struct {
	StockList []Stock `json:"data"`
	ResultCount int `json:"result_count"`
}

func IntrinioQuery(industryGroup string, targetJSON *IndustryGroup) {
	var httpRequestHeaderKey = "Authorization"
	var authUsername = "794a1661b7de6375768a7c2f82fde1cb"
	var authPassword = "f42f60adb529b6a3b0d4052dd979a1ca"
	var httpRequestHeaderValue = "Basic " + base64.StdEncoding.EncodeToString([]byte(authUsername+":"+authPassword))

	industry := url.PathEscape(industryGroup)

	url := "https://api.intrinio.com/securities/search?conditions=industry_group~eq~" + industry
	url = url + "&order_column=marketcap&order_direction=desc"

	fmt.Println(url)

	client := &http.Client{}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating the HTTP request\n")
		return
	}

	request.Header.Add(httpRequestHeaderKey, httpRequestHeaderValue)

	response, err2 := client.Do(request)

	defer response.Body.Close()

	if err2 != nil {
		fmt.Printf("Error servicing the HTTP request\n")
	} else {
		json.NewDecoder(response.Body).Decode(targetJSON)
	}

}