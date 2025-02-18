package scraper

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func GetWebData(url string) string {

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

