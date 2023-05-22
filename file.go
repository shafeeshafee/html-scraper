package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func createOutputDir(dirName string) {
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

// creates a unique filename based on the scrape
func formatFileName(url string) string {
	date := time.Now().Format("01-02-06")
	websiteName := strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://")
	websiteName = strings.ReplaceAll(websiteName, "/", "_")

	// Generate a random 4-digit number
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(9000) + 1000

	return fmt.Sprintf("%s-%s-%d.txt", date, websiteName, randomNumber)
}

func createOutputFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func writeTags(document *goquery.Document, tag string, file *os.File) {
	document.Find(tag).Each(func(index int, element *goquery.Selection) {
		_, err := file.WriteString(element.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	})
}
