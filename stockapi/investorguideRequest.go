package stockapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

func GetIndustries() {
	var r io.Reader	

	response, err := http.Get("http://www.investorguide.com/industry-list.php")
	if err != nil {
		fmt.Println("Error servicing the HTTP request: %s", err)
		return
	}
	
	data, _ := ioutil.ReadAll(response.Body)
	r = strings.NewReader(string(data))
	z := html.NewTokenizer(r)
	var count int = 0
	var industryNameMap map[string]bool = make(map[string]bool)

	for {
		token := z.Next()
		switch token {
			case html.ErrorToken:
				fmt.Println(z.Err())
				return
			case html.TextToken: 
				industryName := strings.TrimSpace(string(z.Text()))
				if(industryName != "" && !strings.HasPrefix(industryName, "\n ") && len(industryName) < 50) {
					count++
					if count >= 53 && count <= 261 {
						industryNameMap[industryName] = true
						fmt.Printf("Added %s to list (total: %d)\n", industryName, len(industryNameMap))
						fmt.Println("-----------------------------")
					}
				}
		}
	}
	
}
