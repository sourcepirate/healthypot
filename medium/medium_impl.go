package medium

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func New(tag string) *MediumTag {
	fmt.Println("Creating a new tag")
	return &MediumTag{tag}
}

func parseDoc(doc *goquery.Document) []MediumRecord {
	records := []MediumRecord{}
	doc.Find(".streamItem").Each(func(i int, s *goquery.Selection) {
		link := s.Find(".button--smaller").Get(0)
		for _, element := range link.Attr {
			if element.Key == "href" {
				records = append(records, MediumRecord{element.Val})
			}
		}
	})
	return records
}

func FetchRecordsForYear(tag string, year string) []MediumRecord {
	url := strings.Join([]string{MEDIUM_TAG, tag, "archive", year}, "/")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal("Non 200 code")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return parseDoc(doc)
}

func (med *MediumTag) fetchArchive(year string) []MediumRecord {
	return FetchRecordsForYear(med.Tag, year)
}

func (med *MediumTag) GetArchive() *MediumArchive {

	recordMaps := make(map[string][]MediumRecord)
	YEARS := []string{"2011", "2012", "2013", "2014", "2015", "2016", "2017", "2018"}
	for _, year := range YEARS {
		log.Println("Fetching archive for year %s for tag %s", year, med.Tag)
		recordMaps[year] = med.fetchArchive(year)
	}

	return &MediumArchive{recordMaps}
}
