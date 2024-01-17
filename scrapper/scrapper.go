package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gocolly/colly"
)

func main() {

	file, err := os.Open("countries.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close() // Close the file when done

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	c := colly.NewCollector()
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			// 각 tr 요소 내에서 class가 "pb-2 tbl-cell tbl-cell-name"인 td 요소의 텍스트 추출
			title := el.ChildText("td.pb-2.tbl-cell.tbl-cell-name")
			fmt.Println("Title found:", title)
		})
	})

	for scanner.Scan() {
		line := scanner.Text()
		site := "https://www.netflix.com/tudum/top10/" + line
		fmt.Println("Visiting", site)

		err := c.Visit(site)
		if err != nil {
			fmt.Println("Error visiting the site:", err)
			continue
		}
	}

	c.Wait() // Wait until all requests are completed

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
	}

}
