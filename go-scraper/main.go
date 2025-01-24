package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Result struct {
	Model string
	Image string
}

func fetchHTML(url string) ([]Result, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching the HTML: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error loading HTML document: %w", err)
	}

	var results []Result
	doc.Find(".modellabel").Each(func(i int, s *goquery.Selection) {
		modelText := s.Text()
		img := s.Closest("tr").Find("td").Eq(1).Find("img")
		src, exists := img.Attr("src")
		if exists {
			results = append(results, Result{Model: modelText, Image: src})
		}
	})

	return results, nil
}

func scrape() ([]Result, error){
	urls := []string{
		"https://www.een.com/hardware/bridges",
		"https://www.een.com/hardware/cmvrs",
	}

	var ans []Result
	for _, url := range urls {
		results, err := fetchHTML(url)
		if err != nil {
			log.Printf("Error fetching URL %s: %v", url, err)
			continue
		}
		ans = append(ans, results...)
	}

	for _, result := range ans {
		fmt.Printf("Model: %s, Image: %s\n", result.Model, result.Image)
	}
	return ans, nil
}

func main() {
	results, err := scrape()
	if err != nil {
		log.Fatalf("Error scraping websites: %v", err)
	}

	for _, result := range results {
		fmt.Printf("Model: %s, Image: %s\n", result.Model, result.Image)
	}
}

