package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main() {
	url := "https://example.com"

	document, err := fetchDocument(url)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	outputDir := "outputs"
	createOutputDir(outputDir)

	fileName := formatFileName(url)
	filePath := filepath.Join(outputDir, fileName)

	file := createOutputFile(filePath)
	defer file.Close()

	writeH1Tags(document, file)

	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Completed, check %s file\n", filePath)
}
