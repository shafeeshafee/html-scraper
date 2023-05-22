package main

import (
	"fmt"
	"log"
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

func formatFileName(url string) string {
	date := time.Now().Format("01-02-06")
	websiteName := strings.TrimPrefix(strings.TrimPrefix(url, "https://"), "http://")
	websiteName = strings.ReplaceAll(websiteName, "/", "_")
	return fmt.Sprintf("%s-%s.txt", date, websiteName)
}

func createOutputFile(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func writeH1Tags(document *goquery.Document, file *os.File) {
	document.Find("h1").Each(func(index int, element *goquery.Selection) {
		_, err := file.WriteString(element.Text() + "\n")
		if err != nil {
			log.Fatal(err)
		}
	})
}
