/*
package main

import (
	"fmt"
	"hockey/features/scraper"
	"log"
	"strings"



	"github.com/PuerkitoBio/goquery"
)

type Player struct {
	Name                string
	Season              string
	Team                string
	Handedness          string // S/C for Shoots/Catches
	Position            string
	GamesPlayed         int
	Goals               int
	Assists             int
	Points              int
	PlusMinus           int
	PenaltyMinutes      int
	PointsPerGame       float64
	EvenStrengthGoals   int
	EvenStrengthPoints  int
	PowerPlayGoals      int
	PowerPlayPoints     int
	ShortHandedGoals    int
	ShortHandedPoints   int
	OvertimeGoals       int
	OvertimeWins        int
	Shots               int
	ShootingPercentage  float64
	TimeOnIcePerGame    float64
	FaceoffWinPercentage float64
}

func ScrapeDataWithHTML(htmlContent string) []string {
	// Load the HTML document from the content obtained via chromedp
	returnObj := []string{}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	// Use a CSS selector to target the tr with the specific class
	doc.Find("tr.sc-dSTIoc.lpdNXV.rt-tr.null").Each(func(i int, s *goquery.Selection) {
		// Get the HTML of the entire <tr> tag (discarding the error)
		trHTML, _ := s.Html() // We only care about the HTML, discard the error
		fmt.Println("Full <tr> HTML:", trHTML)
		returnObj = append(returnObj, trHTML)
	})
	return returnObj
}

func main() {
	url := "https://www.nhl.com/stats/skaters?reportType=season&seasonFrom=20242025&seasonTo=20242025&gameType=2&sort=points,goals,assists&page=0&pageSize=50"
	webData := scraper.GetWebData(url) // Assuming this function gets the web data successfully
	returnVal := ScrapeDataWithHTML(webData)
	fmt.Println(returnVal)
	fmt.Println(webData)
}
*/

package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"hockey/features/scraper"
)

type Player struct {
	Name                string
	Season              string
	Team                string
	Handedness          string // S/C for Shoots/Catches
	Position            string
	GamesPlayed         int
	Goals               int
	Assists             int
	Points              int
	PlusMinus           int
	PenaltyMinutes      int
	PointsPerGame       float64
	EvenStrengthGoals   int
	EvenStrengthPoints  int
	PowerPlayGoals      int
	PowerPlayPoints     int
	ShortHandedGoals    int
	ShortHandedPoints   int
	OvertimeGoals       int
	OvertimeWins        int
	Shots               int
	ShootingPercentage  float64
	TimeOnIcePerGame    string
	FaceoffWinPercentage float64
}

func ScrapeDataWithHTML(htmlContent string) []Player {
	var players []Player

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	// Use a CSS selector to target the tr with the specific class
	doc.Find("tr.sc-dSTIoc.lpdNXV.rt-tr.null").Each(func(i int, s *goquery.Selection) {
		var player Player

		// Extract data from each <td> within the <tr>
		s.Find("td").Each(func(j int, td *goquery.Selection) {
			switch j {
			case 1:
				player.Name = td.Find("a").Text()
			case 2:
				player.Season = td.Text()
			case 3:
				player.Team = td.Find("div").Text()
			case 4:
				player.Handedness = td.Text()
			case 5:
				player.Position = td.Text()
			case 6:
				player.GamesPlayed, _ = strconv.Atoi(td.Text())
			case 7:
				player.Goals, _ = strconv.Atoi(td.Text())
			case 8:
				player.Assists, _ = strconv.Atoi(td.Text())
			case 9:
				player.Points, _ = strconv.Atoi(td.Text())
			case 10:
				player.PlusMinus, _ = strconv.Atoi(strings.TrimSpace(td.Text()))
			case 11:
				player.PenaltyMinutes, _ = strconv.Atoi(td.Text())
			case 12:
				player.PointsPerGame, _ = strconv.ParseFloat(td.Text(), 64)
			case 13:
				player.EvenStrengthGoals, _ = strconv.Atoi(td.Text())
			case 14:
				player.EvenStrengthPoints, _ = strconv.Atoi(td.Text())
			case 15:
				player.PowerPlayGoals, _ = strconv.Atoi(td.Text())
			case 16:
				player.PowerPlayPoints, _ = strconv.Atoi(td.Text())
			case 17:
				player.ShortHandedGoals, _ = strconv.Atoi(td.Text())
			case 18:
				player.ShortHandedPoints, _ = strconv.Atoi(td.Text())
			case 19:
				player.OvertimeGoals, _ = strconv.Atoi(td.Text())
			case 20:
				player.OvertimeWins, _ = strconv.Atoi(td.Text())
			case 21:
				player.Shots, _ = strconv.Atoi(td.Text())
			case 22:
				player.ShootingPercentage, _ = strconv.ParseFloat(td.Text(), 64)
			case 23:
				player.TimeOnIcePerGame = td.Text()
			case 24:
				player.FaceoffWinPercentage, _ = strconv.ParseFloat(td.Text(), 64)
			}
		})

		players = append(players, player)
	})

	return players
}

func main() {
	// Assuming you have the HTML content in a variable called htmlContent
	url := "https://www.nhl.com/stats/skaters?reportType=season&seasonFrom=20242025&seasonTo=20242025&gameType=2&sort=points,goals,assists&page=0&pageSize=50"
	htmlContent := scraper.GetWebData(url) // Replace with actual HTML content

	players := ScrapeDataWithHTML(htmlContent)

	for _, player := range players {
		fmt.Printf("Player: %+v\n", player)
	}
}