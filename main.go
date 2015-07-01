package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
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
	locale         string
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.StringVar(&path, "path", "", "path to a xml.gz or .xml file")
	flag.StringVar(&dir, "dir", "", "location of xml.gz or .xml files")

	flag.StringVar(&businessUnitId, "b", "", "Check for the existence of this business unit id")
	flag.StringVar(&merchantUrl, "url", "", "Check for the existense of this merchant url")
	flag.StringVar(&consumerId, "c", "", "Check for the existence of this consumer id")
	flag.StringVar(&reviewId, "r", "", "Check for the existence of review id")
	flag.StringVar(&locale, "l", "", "only look at the specified locale")

	flag.Parse()

	if path != "" {
		ProcessFeed(path)
		return
	}

	if dir != "" {
		ProcessDir(dir)
		return
	}

	ProcessDir(".")
}

func ProcessDir(dir string) error {
	potentials := []string{}

	finfos, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	if len(finfos) == 0 {
		return fmt.Errorf("No files in dir '%s'", dir)
	}

	locale = strings.ToLower(locale)

	for _, finfo := range finfos {
		if finfo.IsDir() || !strings.HasPrefix(finfo.Name(), "feed_") {
			continue
		}

		if strings.HasSuffix(finfo.Name(), "xml") || strings.HasSuffix(finfo.Name(), "xml.gz") {
			if locale != "" {
				if strings.Contains(strings.ToLower(finfo.Name()), locale) {
					potentials = append(potentials, finfo.Name())
				}
			} else {
				potentials = append(potentials, finfo.Name())
			}
		}
	}

	resp := make(chan int)
	for _, path := range potentials {
		go func(path string) {
			resp <- ProcessFeed(path)
		}(path)
	}
	matches := 0
	searching := businessUnitId == "" || consumerId == "" || reviewId == "" || merchantUrl == ""
	for _, _ = range potentials {
		matches += <-resp
	}

	if searching {
		log.Println("Matches: ", matches)
	} else {
		log.Println("Done")
	}

	return nil
}

func ProcessFeed(path string) int {
	log.Println("Processing: " + path)

	feed, err := ParseFeed(path)
	if err != nil {
		log.Println(err)
		fmt.Printf("%s failed to parse\n")
		return 0
	}

	if businessUnitId == "" && consumerId == "" && reviewId == "" && merchantUrl == "" {
		fmt.Printf("%s parsed successfully\n", path)
		return 0
	}

	matches := 0
	for _, m := range feed {
		if businessUnitId != "" && businessUnitId == m.BusinessUnitId {
			fmt.Printf("Merchant - %s\n\tName: %s \n\tBusinessUnitId: %s\n\tUrl: %s\n\tReviews: %d\n",
				path, m.Name, m.BusinessUnitId, m.Url, len(m.Reviews))
			matches++
		}

		if merchantUrl != "" && strings.HasPrefix(m.Url, merchantUrl) {
			fmt.Printf("Merchant - %s\n\tName: %s \n\tBusinessUnitId: %s\n\tUrl: %s\n\tReviews: %d\n",
				path, m.Name, m.BusinessUnitId, m.Url, len(m.Reviews))
			matches++
		}

		if consumerId != "" || reviewId != "" {
			for _, r := range m.Reviews {
				if r.ReviewerId == consumerId {
					fmt.Printf("Consumer - %s\n\tMerchant: %s\n\tConsumerId: %s\n\tReviewUrl: %s\n",
						path, m.Name, r.ReviewerId, r.ReviewUrl)
					matches++
				}

				if r.Id == reviewId {
					fmt.Printf("Review - %s\n\tMerchant: %s\n\tConsumerId: %s\n\tReviewUrl: %s\n",
						path, m.Name, r.ReviewerId, r.ReviewUrl)
					matches++
				}
			}
		}
	}

	return matches
}
