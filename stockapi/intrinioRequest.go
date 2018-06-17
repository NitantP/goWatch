package stockapi

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func IntrinioQuery() {
	var httpRequestHeaderKey = "Authorization"
	var authUsername = "794a1661b7de6375768a7c2f82fde1cb"
	var authPassword = "f42f60adb529b6a3b0d4052dd979a1ca"
	var httpRequestHeaderValue = "Basic " + base64.StdEncoding.EncodeToString([]byte(authUsername+":"+authPassword))

	client := &http.Client{}

	request, err := http.NewRequest("GET", "https://api.intrinio.com/securities/search?conditions=sector~eq~Consumer%20Goods&order_column=marketcap&order_direction=desc", nil)
	if err != nil {
		fmt.Printf("Error creating the HTTP request\n")
		return
	}
	request.Header.Add(httpRequestHeaderKey, httpRequestHeaderValue)

	response, err2 := client.Do(request)
	if err2 != nil {
		fmt.Printf("Error servicing the HTTP request\n")
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
	}
}
