package main

import (
	"context"
	"fmt"
	"os"

	"github.com/chromedp/chromedp"
)

//sc-dSTIoc lpdNXV rt-tr null tr class

func getWebData (url string)  string {

	ctx, cancel := chromedp.NewContext(context.Background())
    defer cancel()

    var res string
    err := chromedp.Run(ctx, 
		chromedp.Navigate(url),
		chromedp.WaitVisible("tbody.rt-tbody", chromedp.ByQuery),
		chromedp.OuterHTML("tbody.rt-tbody", &res, chromedp.ByQuery),
	)
    if err != nil {
          fmt.Println("Error navigating to the website:", err)
          return ""
    }
	return res
}

func writeToFile (fileName string, data string) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Ensure the file is closed when we're done
	defer file.Close()

	// Write the text to the file
	_, err = file.WriteString(data)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func main() {
	url := "https://www.nhl.com/stats/skaters?reportType=season&seasonFrom=20242025&seasonTo=20242025&gameType=2&sort=points,goals,assists&page=0&pageSize=50"
	webData := getWebData(url)
	writeToFile("test.txt", webData)

}
