package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	// get url
	fmt.Print("Enter a URL to scrape: ")
	var url string
	fmt.Scanln(&url)

	// get html tag you want to gather from
	fmt.Print("Enter an HTML tag to scrape (e.g., h1, p, div): ")
	var tag string
	fmt.Scanln(&tag)

	document, err := fetchDocument(url)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// outputs in a file called dir
	outputDir := "outputs"
	createOutputDir(outputDir)

	fileName := formatFileName(url)
	filePath := filepath.Join(outputDir, fileName)

	file := createOutputFile(filePath)
	defer file.Close()

	writeTags(document, tag, file)

	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Scraping done, check %s file\n", filePath)
}
