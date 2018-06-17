package stockapi

import (
	"fmt"
    "io/ioutil"
    "net/http"
	"strconv"
)


func BuildURL(stock string, interval int) string {

	return "https://www.google.com/finance/getprices?q=" + stock + "&i=" + strconv.Itoa(interval) + "&f=d,c,o,h,l"
}

func RequestStockData(stock string, interval int) string {


	url := BuildURL(stock, interval)
	resp, err := http.Get(url)

    if err != nil {
    	fmt.Println("HTTP REQUEST FAILED")
    }

    defer resp.Body.Close()

    if resp.StatusCode == http.StatusOK {
    	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
    	bodyString := string(bodyBytes)
    	fmt.Println(bodyString)

    	if err2 != nil {
    		fmt.Println("FAILED TO READ RESPONSE BODY")
    	}

    	return bodyString
	}

	fmt.Println("HTTP STATUS NOT OK")
	return "FAIL"
}
