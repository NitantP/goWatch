package main

import (
    "goWatch/stockapi"
    "os"
    "bufio"
    "fmt"
    "strings"
)

func main() {

    inputReader := bufio.NewReader(os.Stdin)
    fmt.Printf("Enter industry name: ")
    inputText, readerErr := inputReader.ReadString('\n')
    if readerErr != nil {
        fmt.Printf("Error parsing user input: %s\n", readerErr)
        return
    }

    inputText = strings.TrimSuffix(inputText, "\n")
    inputText = strings.ToLower(inputText)
    industryNameMap := stockapi.GetIndustries()
    if !industryNameMap[inputText] {
        fmt.Printf("Industry name (%s) not found\n", inputText)
        return
    }
    fmt.Printf("Industry name (%s) found\n", inputText)

}
