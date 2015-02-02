package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

var (
	// path and regex are used to locate files
	path  string
	dir   string
	regex string
	// ids
	businessUnitId string
	consumerId     string
	reviewId       string
	merchantUrl    string
)

func main() {
	flag.StringVar(&path, "path", "", "path to a xml.gz or .xml file")
	// flag.StringVar(&regex, "regex", "", "regex that matches xml.gz or zml files")
	// flag.StringVar(&dir, "dir", "", "location of xml.gz or .xml files")
	flag.StringVar(&businessUnitId, "b", "", "Check for the existence of this business unit id")
	flag.StringVar(&merchantUrl, "url", "", "Check for the existense of this merchant url")
	flag.StringVar(&consumerId, "c", "", "Check for the existence of this consumer id")
	flag.StringVar(&reviewId, "r", "", "Check for the existence of review id")

	flag.Parse()

	if path != "" {
		ProcessFeed(path)
		return
	}

	// if dir != "" {

	// }

	// if regex != "" {

	// }
}

func ProcessDir(dir string) {}

func ProcessRegex(regex string) {

}

func ProcessFeed(path string) {
	feed, err := ParseFeed(path)
	if err != nil {
		log.Fatal(err)
	}

	if businessUnitId == "" && consumerId == "" && reviewId == "" && merchantUrl == "" {
		fmt.Println("Parsed the GSR feed successfully")
		return
	}

	for _, m := range feed {
		if businessUnitId != "" && businessUnitId == m.BusinessUnitId {
			fmt.Printf("Merchant\n\tName: %s \n\tBusinessUnitId: %s\n\tUrl: %s\n\tReviews: %d\n",
				m.Name, m.BusinessUnitId, m.Url, len(m.Reviews))
		}

		if merchantUrl != "" && strings.HasPrefix(m.Url, merchantUrl) {
			fmt.Printf("Merchant\n\tName: %s \n\tBusinessUnitId: %s\n\tUrl: %s\n\tReviews: %d\n",
				m.Name, m.BusinessUnitId, m.Url, len(m.Reviews))
		}

		if consumerId != "" || reviewId != "" {
			for _, r := range m.Reviews {
				if r.ReviewerId == consumerId {
					fmt.Printf("Consumer\n\tMerchant: %s\n\tConsumerId: %s\n\tReviewUrl: %s\n",
						m.Name, r.ReviewerId, r.ReviewUrl)
				}

				if r.Id == reviewId {
					fmt.Printf("Review\n\tMerchant: %s\n\tConsumerId: %s\n\tReviewUrl: %s\n",
						m.Name, r.ReviewerId, r.ReviewUrl)
				}
			}
		}
	}

}
