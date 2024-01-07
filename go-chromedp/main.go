package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

// ElementData holds the scraped data
type ElementData struct {
	DateTime      string `json:"datetime"`
	LinkText      string `json:"linkText"`
	Source string `json:"source"`
}

func main() {
    // Initialize Chrome
    opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.NoSandbox,
		chromedp.DisableGPU,
		chromedp.Headless,
	)

	// Use the options when creating a context
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// Use the allocator context when creating a new chromedp context
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

    // Replace with your target URL
	// Define the ticker variable
	ticker := "MSFT"

	// Create the target URL using fmt.Sprintf
	targetURL := fmt.Sprintf("https://news.google.com/search?q=%s&hl=en-IN&gl=IN&ceid=IN%%3Aen", ticker)

	fmt.Println(targetURL)


    // Define the CSS selectors for the elements you want to scrape
    elementSelector := ".PO9Zff.Ccj79.kUVvS"
    timeSelector := "time.hvbAAd"
    linkSelector := "a.JtKRv"
    divSelector := "div.vr1PYe"

    // Scrape data
    elementsData, err := scrapeData(ctx, targetURL, elementSelector, timeSelector, linkSelector, divSelector)
    if err != nil {
        log.Fatal("Error while performing the automation logic:", err)
    }

    // Convert the entire slice to JSON
    jsonData, err := json.Marshal(elementsData)
    if err != nil {
        log.Fatal(err)
    }

    // Write the JSON data to a single file
    if err := os.WriteFile("all_data3.json", jsonData, 0644); err != nil {
        log.Fatal(err)
    }

    fmt.Println("All data saved successfully in all_data.json")
}

func scrapeData(ctx context.Context, targetURL, elementSelector, timeSelector, linkSelector, divSelector string) ([]ElementData, error) {
    var elementsData []ElementData // Slice to hold data of each element

    // Run chromedp tasks
    err := chromedp.Run(ctx,
        chromedp.Navigate(targetURL),
        chromedp.Sleep(2000*time.Millisecond), // Wait for page to load

        // Extract required data
        chromedp.Evaluate(fmt.Sprintf(`
            const elements = Array.from(document.querySelectorAll('%s'));
            const data = elements.map(el => {
                let timeElement = el.querySelector('%s');
                let linkElement = el.querySelector('%s');
                let divElement = el.querySelector('%s');
                return {
                    datetime: timeElement ? timeElement.getAttribute('datetime') : '',
                    linkText: linkElement ? linkElement.textContent.trim() : '',
                    source: divElement ? divElement.textContent.trim() : ''
                };
            });
            data;
        `, elementSelector, timeSelector, linkSelector, divSelector), &elementsData),
    )
    if err != nil {
        return nil, err
    }

    return elementsData, nil
}
