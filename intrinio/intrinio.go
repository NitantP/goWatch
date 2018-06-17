package intrinio

import(
    "fmt"
    "io/ioutil"
    "net/http"
    "encoding/base64"
)

func IntrinioQuery() {
    var httpRequestHeaderKey = "Authorization"
    var authUsername = "794a1661b7de6375768a7c2f82fde1cb"
    var authPassword = "f42f60adb529b6a3b0d4052dd979a1ca"
    var httpRequestHeaderValue = "Basic " + base64.StdEncoding.EncodeToString([]byte(authUsername + ":" + authPassword))
    
    client := &http.Client{}
    response, err := client.Get("https://api.intrinio.com/indices?type=sic")
    if err != nil {
        fmt.Printf("Error creating the client\n")
        return
    }
    
    request, err := http.NewRequest("GET", "https://api.intrinio.com/indices?type=sic", nil)
    request.Header.Add(httpRequestHeaderKey, httpRequestHeaderValue)
    
    response, err = client.Do(request)
    if err != nil {
        fmt.Printf("Error servicing the HTTP request\n")
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Println(string(data))
    }
}

