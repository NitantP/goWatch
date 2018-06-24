package stockapi

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"golang.org/x/net/html"
)

func GetIndustries() map[string]bool {
	var r io.Reader	

	response, err := http.Get("http://www.investorguide.com/industry-list.php")
	if err != nil {
		fmt.Printf("Error servicing the HTTP request: %s\n", err)
		return nil
	}
	
	data, _ := ioutil.ReadAll(response.Body)
	r = strings.NewReader(string(data))
	z := html.NewTokenizer(r)
	var count int = 0
	var industryNameMap map[string]bool = make(map[string]bool)

	for {
		token := z.Next()
		if token == html.ErrorToken {
			if z.Err() == io.EOF {
				break
			}
			fmt.Println(z.Err())
			return nil
		}
		if token == html.TextToken {
				industryName := strings.TrimSpace(string(z.Text()))
				if(industryName != "" && !strings.HasPrefix(industryName, "\n ") && len(industryName) < 50) {
					count++
					if count >= 53 && count <= 261 {
						industryNameMap[strings.ToLower(industryName)] = true
						// fmt.Printf("Added %s to list (total: %d)\n", industryName, len(industryNameMap))
						// fmt.Println("-----------------------------")
					}
				}
		}
	}

	// fmt.Printf("\n\nSize of map: %d\n\n", len(industryNameMap))
	// for key, value := range industryNameMap {
	// 	fmt.Printf("Key: %s, Value: %t\n", key, value)
	// }

	return industryNameMap

}
