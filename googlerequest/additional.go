package googlerequest
import "fmt"
import "net/http"
import "io/ioutil"

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