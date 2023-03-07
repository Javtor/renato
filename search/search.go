package search

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

func preprocessImages(images []string) []string {
	var filteredImages []string
	for _, image := range images {
		if strings.HasSuffix(image, ".jpg") || strings.HasSuffix(image, ".png") {
			filteredImages = append(filteredImages, image)
		}
	}

	return filteredImages
}

func GetRandomImage(query string) (string, error) {
	images := preprocessImages(getSearch(query))
	if len(images) == 0 {
		return "", errors.New("No images found")
	}

	return images[rand.Intn(len(images))], nil
}

func getSearch(searchQuery string) []string {
	searchString := strings.Replace(searchQuery, " ", "-", -1)
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X x.y; rv:42.0) Gecko/20100101 Firefox/42.0"
	c.AllowURLRevisit = true
	c.DisableCookies()
	array := []string{}

	url := fmt.Sprintf("https://google-scrapper.orewa.workers.dev/?img=%v", url.QueryEscape(searchQuery))

	c.OnHTML(".islrtb.isv-r", func(e *colly.HTMLElement) {
		src := e.Attr("data-ou")
		if src != "" {
			array = append(array, src)
		}
	})

	// Requesting a url for html
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	searchString = strings.Replace(searchString, " ", "+", -1)

	c.Visit(url)

	return array
}
