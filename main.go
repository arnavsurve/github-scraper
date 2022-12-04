package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

var fullName string
var userName string
var pfpUrl string

var description string

var followCount [2]string

func scrape(username string) {
	var result []string
	result = make([]string, 3)

	url := fmt.Sprintf("https://github.com/%s", username)

	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)

	// name
	c.OnHTML("span.vcard-fullname", func(e *colly.HTMLElement) {
		fullName = strings.TrimSpace(e.Text)
		result = append(result, fullName)
	})

	// username
	c.OnHTML("span.p-nickname", func(e *colly.HTMLElement) {
		userName = strings.TrimSpace(e.Text)
		result = append(result, username)
	})

	// profile image url
	c.OnHTML("img.avatar-user.width-full", func(e *colly.HTMLElement) {
		pfpUrl = e.Attr("src")
		result = append(result, pfpUrl)
	})

	// bio / description
	c.OnHTML("div.user-profile-bio", func(e *colly.HTMLElement) {
		description = e.Text
		result = append(result, description)
	})

	// followers & following
	i := 0
	c.OnHTML("span.text-bold", func(e *colly.HTMLElement) {
		if i < 2 {
			followCount[i] = e.Text
			result = append(result, e.Text)
			i++
		}
	})

	c.Visit(url)
}

func main() {
	fmt.Print("Enter Github profile username: ")
	fmt.Scan(&userName)
	fmt.Println()

	scrape(userName)
	res := fmt.Sprintf("Name: %s\nUsername: %s\nProfile Image URL: %s\nBio: %s\nFollowers: %s\nFollowing: %s", fullName, userName, pfpUrl, description, followCount[0], followCount[1])
	fmt.Println(res)
}
